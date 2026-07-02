package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	BaseModel
	RoomID           int64  `json:"room_id"`
	UserID           int64  `json:"user_id"`
	DeviceID         int64  `json:"device_id"`
	EncryptedPayload string `json:"encrypted_payload"`
}

type MessageRepository interface {
	GetChatHistory(ctx context.Context, roomID uuid.UUID, limit int, before time.Time) ([]*Message, error)
	GetPendingMessages(ctx context.Context, deviceID uuid.UUID) ([]*Message, error)
	UpdateStatus(ctx context.Context, messageID uuid.UUID, status string) error
}
