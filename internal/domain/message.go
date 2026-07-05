package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type MessageStatus string

const (
	MessageStatusPending   MessageStatus = "PENDING"
	MessageStatusDelivered MessageStatus = "DELIVERED"
)

type Message struct {
	BaseModel
	RoomID           int64         `json:"room_id"`
	UserID           int64         `json:"user_id"`
	DeviceID         int64         `json:"device_id"`
	EncryptedPayload string        `json:"encrypted_payload"`
	Status           MessageStatus `json:"status"`
}

type MessageRepository interface {
	Create(ctx context.Context, message *Message) error
	GetChatHistory(ctx context.Context, roomID uuid.UUID, before time.Time) ([]*Message, error)
	GetPendingMessages(ctx context.Context, deviceID uuid.UUID) ([]*Message, error)
	UpdateStatus(ctx context.Context, messageID uuid.UUID, status MessageStatus) error
}
