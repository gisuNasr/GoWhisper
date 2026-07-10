package service

import (
	"context"

	"github.com/gisuNasr/GoWhisper/internal/application/dto"
	"github.com/gisuNasr/GoWhisper/internal/domain"
	"github.com/google/uuid"
)

type RoomService struct {
	repo domain.RoomAggregatorRepository
}

func NewRoomService(repo domain.RoomAggregatorRepository) *RoomService {
	return &RoomService{repo}
}

func (s *RoomService) Create(ctx context.Context, req dto.CreateRoomRequest) (*dto.RoomResponse, error) {
	if req.Name == "" {
		return nil, domain.ErrInvalidInput
	}

	room := &domain.Room{
		Name: req.Name,
		Type: domain.RoomType(req.RoomType),
	}
	room.ID = uuid.New()
	err := s.repo.Create(ctx, room)
	if err != nil {
		return nil, err
	}
	return dto.ToRoomResponse(room), nil
}

func (s *RoomService) GetByID(ctx context.Context, id uuid.UUID) (*dto.RoomResponse, error) {
	room, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return dto.ToRoomResponse(room), nil
}

func (s *RoomService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteRoom(ctx, id)
}

//AddMemberToRoom(ctx context.Context, roomID, userId uuid.UUID) error
//RemoveMemberFromRoom(ctx context.Context, roomID, userId uuid.UUID) error
//GetRoomMembers(ctx context.Context, roomID uuid.UUID) ([]*User, error)
//GetUserRooms(ctx context.Context, userID uuid.UUID) ([]*Room, error)
