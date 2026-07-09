package websocket

import (
	"sync"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

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
