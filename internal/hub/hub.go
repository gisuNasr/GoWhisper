package hub

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type MessagePayload struct {
	RoomID uuid.UUID
	Data   []byte
}

type Hub struct {
	Users map[uuid.UUID]map[uuid.UUID]*Client

	Rooms map[uuid.UUID]map[*Client]bool

	Broadcast chan MessagePayload

	Register   chan *Client
	Unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Users:      make(map[uuid.UUID]map[uuid.UUID]*Client),
		Rooms:      make(map[uuid.UUID]map[*Client]bool),
		Broadcast:  make(chan MessagePayload),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}
