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
