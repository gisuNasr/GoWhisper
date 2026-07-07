package websocket

import (
	"log"
	"sync"

	"github.com/google/uuid"
)

type Hub struct {
	mu    sync.RWMutex
	users map[uuid.UUID]map[uuid.UUID]*Client
}

func NewHub() *Hub {
	return &Hub{
		users: make(map[uuid.UUID]map[uuid.UUID]*Client),
	}
}

func (h *Hub) Register(c *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if _, exists := h.users[c.UserID]; !exists {
		h.users[c.UserID] = make(map[uuid.UUID]*Client)
	}
	h.users[c.UserID][c.DeviceID] = c
}

func (h *Hub) Unregister(userID, deviceID uuid.UUID) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if devices, exists := h.users[userID]; exists {
		delete(devices, deviceID)
		if len(devices) == 0 {
			delete(h.users, userID)
		}
	}
}

func (h *Hub) SendToDevice(userID, deviceID uuid.UUID, payload []byte) bool {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if devices, exists := h.users[userID]; exists {
		if client, connected := devices[deviceID]; connected {
			select {
			case client.Send <- payload:
				return true
			default:
				log.Printf("hub: send buffer full for device %s — message will be stored as pending", deviceID)
				return false
			}
		}
	}
	return false
}

func (h *Hub) IsOnline(userID uuid.UUID) bool {
	h.mu.RLock()
	defer h.mu.RUnlock()

	devices, exists := h.users[userID]
	return exists && len(devices) > 0
}
