package repository

import (
	"context"
	"time"

	"github.com/gisuNasr/GoWhisper/internal/domain"
	"github.com/google/uuid"
)

type messageRepository struct {
	*BaseRepository[domain.Message]
}

func NewMessageRepository() domain.MessageRepository {
	return &messageRepository{
		NewBaseRepository[domain.Message](),
	}
}

func (r *messageRepository) Create(ctx context.Context, message *domain.Message) error {
	createdModel, err := r.BaseRepository.Create(ctx, *message)
	if err != nil {
		return err
	}
	*message = createdModel
	return nil
}

func (r *messageRepository) GetChatHistory(ctx context.Context, roomID uuid.UUID, before time.Time) ([]*domain.Message, error) {
	var messages []*domain.Message
	err := r.DB.WithContext(ctx).Where("room_id = ? AND created_at < ?", roomID, before).Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}
func (r *messageRepository) GetPendingMessages(ctx context.Context, deviceID uuid.UUID) ([]*domain.Message, error) {
	count, items, err := r.BaseRepository.GetByFilter(ctx, 15, 0, map[string]interface{}{"device_id": deviceID, "status": domain.MessageStatusPending})
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, nil
	}
	messages := make([]*domain.Message, len(items))

	for i := range items {
		messages[i] = &items[i]
	}

	return messages, nil
}

func (r *messageRepository) UpdateStatus(ctx context.Context, messageID uuid.UUID, status domain.MessageStatus) error {
	updateData := make(map[string]interface{})
	updateData["status"] = status
	_, err := r.BaseRepository.Update(ctx, messageID, updateData)
	return err
}
