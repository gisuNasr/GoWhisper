package domain

import (
	"context"

	"github.com/google/uuid"
)

type RoomType string

const (
	RoomTypePV    RoomType = "pv"
	RoomTypeGroup RoomType = "group"
)

type Room struct {
	BaseModel
	Name string   `json:"name"`
	Type RoomType `json:"type"`
}

type RoomMember struct {
	RoomID int64 `json:"room_id"`
	UserID int64 `json:"user_id"`
}

type RoomRepository interface {
	Create(ctx context.Context, room *Room) error
	GetByID(ctx context.Context, id uuid.UUID) (*Room, error)
	DeleteRoom(ctx context.Context, id uuid.UUID) error
}

type RoomMemberRepository interface {
	AddMemberToRoom(ctx context.Context, roomID, userId uuid.UUID) error
	RemoveMemberFromRoom(ctx context.Context, roomID, userId uuid.UUID) error
	GetRoomMembers(ctx context.Context, roomID uuid.UUID) ([]*User, error)
	GetUserRooms(ctx context.Context, user uuid.UUID) ([]*Room, error)
}
