package dto

import (
	"github.com/gisuNasr/GoWhisper/internal/domain"
	"github.com/google/uuid"
)

type CreateRoomRequest struct {
	Name     string `json:"name"`
	RoomType string `json:"type"`
}

type RoomResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Type string    `json:"type"`
}

func ToRoomResponse(r *domain.Room) *RoomResponse {
	return &RoomResponse{
		ID:   r.ID,
		Name: r.Name,
		Type: string(r.Type),
	}
}
