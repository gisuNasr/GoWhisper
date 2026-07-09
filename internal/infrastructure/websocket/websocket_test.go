package websocket

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func TestHub_ConcurrencyAndRaceCondition(t *testing.T) {
	hub := NewHub()
	var wg sync.WaitGroup

	userID := uuid.New()
	deviceID := uuid.New()

	clientMock := &Client{
		UserID:   userID,
		DeviceID: deviceID,
		Send:     make(chan []byte, 10),
	}

	hub.Register(clientMock)
	assert.True(t, hub.IsOnline(userID))

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			hub.SendToDevice(userID, deviceID, []byte("test_message"))
		}()
	}
	wg.Wait()

	assert.Equal(t, 10, len(clientMock.Send))

	hub.Unregister(userID, deviceID)
	assert.False(t, hub.IsOnline(userID))
}

func TestIntegration_HubToClientFlow(t *testing.T) {
	hub := NewHub()
	deviceID := uuid.New()
	userID := uuid.New()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		connection, err := Upgrader.Upgrade(writer, request, nil)
		assert.NoError(t, err)

		client := NewClient(connection, userID, deviceID)
		hub.Register(client)

		go client.WritePump(ctx)
		<-ctx.Done()
	}))

	defer server.Close()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http")
	ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	assert.NoError(t, err)
	defer ws.Close()

	payload := []byte("secret_message")

	succes := hub.SendToDevice(userID, deviceID, payload)
	assert.True(t, succes)

	_, received, err := ws.ReadMessage()
	assert.NoError(t, err)
	assert.Equal(t, string(payload), string(received))
}
