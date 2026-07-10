package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gisuNasr/GoWhisper/internal/application/dto"
	"github.com/gisuNasr/GoWhisper/internal/domain"
	"github.com/google/uuid"
)

type MessageService struct {
	repo domain.MessageRepository
	hub  domain.MessageDispatcher
}

func NewMessageService(repo domain.MessageRepository, h domain.MessageDispatcher) *MessageService {
	return &MessageService{
		repo: repo,
		hub:  h,
	}
}

func (s *MessageService) ProcessNewMessage(ctx context.Context, req dto.DispatchMessageRequest) (*dto.MessageResponse, error) {
	if req.EncryptedPayload == "" {
		return nil, domain.ErrInvalidInput
	}

	message := &domain.Message{
		RoomID:           req.RoomID,
		UserID:           req.UserID,
		DeviceID:         req.DeviceID,
		EncryptedPayload: req.EncryptedPayload,
		Status:           domain.MessageStatusPending,
	}

	message.ID = uuid.New()

	if err := s.repo.Create(ctx, message); err != nil {
		return nil, err
	}

	res := dto.ToMessageResponse(message)
	payload, err := json.Marshal(res)
	if err == nil {
		if delivered := s.hub.SendToDevice(req.UserID, req.DeviceID, payload); delivered {
			_ = s.repo.UpdateStatus(ctx, message.ID, domain.MessageStatusDelivered)
			res.Status = string(domain.MessageStatusDelivered)
		}
	}

	return res, nil
}

func (s *MessageService) GetHistory(ctx context.Context, roomID uuid.UUID, limit int, before time.Time) ([]*dto.MessageResponse, error) {
	messages, err := s.repo.GetChatHistory(ctx, roomID, before)
	if err != nil {
		return nil, err
	}

	results := make([]*dto.MessageResponse, len(messages))
	for i, message := range messages {
		results[i] = dto.ToMessageResponse(message)
	}

	return results, nil
}

func (s *MessageService) GetPending(ctx context.Context, deviceID uuid.UUID) ([]*dto.MessageResponse, error) {
	messages, err := s.repo.GetPendingMessages(ctx, deviceID)
	if err != nil {
		return nil, err
	}

	results := make([]*dto.MessageResponse, len(messages))
	for i, message := range messages {
		results[i] = dto.ToMessageResponse(message)
	}
	return results, nil
}

func (s *MessageService) MarkDelivered(ctx context.Context, messageID uuid.UUID) error {
	return s.repo.UpdateStatus(ctx, messageID, domain.MessageStatusDelivered)
}
