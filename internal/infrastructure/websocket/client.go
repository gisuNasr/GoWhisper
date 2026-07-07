package websocket

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

const pongWait = 60 * time.Second

const pingPeriod = (pongWait * 9) / 10
const writeWait = 10 * time.Second

type Client struct {
	Conn     *websocket.Conn
	Send     chan []byte
	UserID   uuid.UUID
	DeviceID uuid.UUID
}

func NewClient(conn *websocket.Conn, userID, deviceId uuid.UUID) *Client {
	return &Client{
		Conn:     conn,
		Send:     make(chan []byte, 256),
		UserID:   userID,
		DeviceID: deviceId,
	}
}

func (c *Client) WritePump(ctx context.Context) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				_ = c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
				_ = c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			_ = c.Conn.SetWriteDeadline(time.Now().Add(writeWait))

			if err := c.Conn.WriteMessage(websocket.BinaryMessage, message); err != nil {
				log.Printf("ws write error user=%s device=%s: %v", c.UserID, c.DeviceID, err)
				return
			}

		case <-ticker.C:
			_ = c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Printf("ws ping error user=%s device=%s: %v", c.UserID, c.DeviceID, err)
				return
			}

		case <-ctx.Done():
			_ = c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			cleanCloseMessage := websocket.FormatCloseMessage(websocket.CloseNormalClosure, "server shutting down")
			_ = c.Conn.WriteMessage(websocket.CloseMessage, cleanCloseMessage)
			return
		}
	}
}
func (c *Client) ReadPump(ctx context.Context, cancel context.CancelFunc, onMessage func([]byte)) {
	defer cancel()

	c.Conn.SetReadLimit(512 * 1024)
	_ = c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error {
		return c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	})

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err,
				websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure,
				websocket.CloseNormalClosure,
			) {
				log.Printf("ws unexpected close user=%s device=%s: %v", c.UserID, c.DeviceID, err)
			}
			return
		}
		onMessage(msg)
	}
}
