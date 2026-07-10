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

func (s *RoomService) AddMember(ctx context.Context, roomID, userId uuid.UUID) error {
	return s.repo.AddMemberToRoom(ctx, roomID, userId)
}

func (s *RoomService) RemoveMember(ctx context.Context, roomID, userId uuid.UUID) error {
	return s.repo.RemoveMemberFromRoom(ctx, roomID, userId)
}

func (s *RoomService) GetRoomMembers(ctx context.Context, roomID uuid.UUID) ([]*dto.UserResponse, error) {
	users, err := s.repo.GetRoomMembers(ctx, roomID)
	if err != nil {
		return nil, err
	}

	result := make([]*dto.UserResponse, len(users))
	for i, user := range users {
		result[i] = dto.ToUserResponse(user)
	}
	return result, err
}

func (s *RoomService) GetUserRooms(ctx context.Context, userID uuid.UUID) ([]*dto.RoomResponse, error) {
	rooms, err := s.repo.GetUserRooms(ctx, userID)
	if err != nil {
		return nil, err
	}

	result := make([]*dto.RoomResponse, len(rooms))
	for i, room := range rooms {
		result[i] = dto.ToRoomResponse(room)
	}
	return result, err
}
