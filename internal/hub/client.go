package hub

import "github.com/google/uuid"

type Client struct {
	Conn     *websocket.Conn
	Send     chan []byte
	UserID   uuid.UUID
	DeviceID uuid.UUID
}
