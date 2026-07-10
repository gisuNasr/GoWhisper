package dto

import (
	"time"

	"github.com/gisuNasr/GoWhisper/internal/domain"
	"github.com/google/uuid"
)

type SendMessageRequest struct {
	RoomID           uuid.UUID `json:"room_id"`
	UserID           uuid.UUID `json:"user_id"`
	DeviceID         uuid.UUID `json:"device_id"`
	EncryptedPayload string    `json:"encrypted_payload"`
}

type DispatchMessageRequest struct {
	RoomID            uuid.UUID `json:"room_id"`
	UserID            uuid.UUID `json:"user_id"`
	DeviceID          uuid.UUID `json:"device_id"`
	RecipientUserID   uuid.UUID `json:"recipient_user_id"`
	RecipientDeviceID uuid.UUID `json:"recipient_device_id"`
	EncryptedPayload  string    `json:"encrypted_payload"`
}

type MessageResponse struct {
	ID               uuid.UUID `json:"id"`
	RoomID           uuid.UUID `json:"room_id"`
	UserID           uuid.UUID `json:"user_id"`
	DeviceID         uuid.UUID `json:"device_id"`
	EncryptedPayload string    `json:"encrypted_payload"`
	Status           string    `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
}

func ToMessageResponse(m *domain.Message) *MessageResponse {
	return &MessageResponse{
		ID:               m.ID,
		RoomID:           m.RoomID,
		UserID:           m.UserID,
		DeviceID:         m.DeviceID,
		EncryptedPayload: m.EncryptedPayload,
		Status:           string(m.Status),
		CreatedAt:        m.CreatedAt,
	}
}
