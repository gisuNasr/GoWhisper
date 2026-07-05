package repository

import (
	"context"

	"github.com/gisuNasr/GoWhisper/internal/domain"
	"github.com/google/uuid"
)

type roomRepository struct {
	*BaseRepository[domain.Room]
}

func NewRoomRepository() domain.RoomAggregatorRepository {
	return &roomRepository{
		BaseRepository: NewBaseRepository[domain.Room](),
	}
}

func (r *roomRepository) Create(ctx context.Context, room *domain.Room) error {
	createdRoom, err := r.BaseRepository.Create(ctx, *room)
	if err != nil {
		return err
	}
	*room = createdRoom
	return nil
}

func (r *roomRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Room, error) {
	room, err := r.BaseRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &room, nil
}
func (r *roomRepository) DeleteRoom(ctx context.Context, id uuid.UUID) error {
	err := r.BaseRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *roomRepository) AddMemberToRoom(ctx context.Context, roomID, userId uuid.UUID) error {
	member := &domain.RoomMember{
		roomID,
		userId,
	}

	err := r.DB.WithContext(ctx).Create(member).Error

	return err
}

func (r *roomRepository) RemoveMemberFromRoom(ctx context.Context, roomID, userId uuid.UUID) error {
	err := r.DB.WithContext(ctx).Where("room_id = ? AND user_id = ?", roomID, userId).Delete(&domain.RoomMember{}).Error
	return err
}

func (r *roomRepository) GetRoomMembers(ctx context.Context, roomID uuid.UUID) ([]*domain.User, error) {
	var users []*domain.User
	err := r.DB.WithContext(ctx).
		Joins("JOIN room_members rm ON rm.user_id = users.id").
		Where("rm.room_id = ?", roomID).
		Find(&users).Error

	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *roomRepository) GetUserRooms(ctx context.Context, userID uuid.UUID) ([]*domain.Room, error) {
	var rooms []*domain.Room

	err := r.DB.WithContext(ctx).Joins("JOIN room_members rm ON rm.room_id = rooms.id").Where("rm.user_id = ?", userID).Find(&rooms).Error
	if err != nil {
		return nil, err
	}
	return rooms, nil
}
