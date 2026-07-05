package websocket

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	Send     chan []byte
	UserID   uuid.UUID
	DeviceID uuid.UUID
}
