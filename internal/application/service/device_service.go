package service

import (
	"context"
	"encoding/json"

	"github.com/gisuNasr/GoWhisper/internal/application/dto"
	"github.com/gisuNasr/GoWhisper/internal/domain"
	"github.com/google/uuid"
)

type DeviceService struct {
	repo domain.DeviceRepository
}

func NewDeviceService(repo domain.DeviceRepository) *DeviceService {
	return &DeviceService{repo: repo}
}

func (s *DeviceService) RegisterDevice(ctx context.Context, req dto.RegisterDeviceRequest) (*dto.DeviceResponse, error) {
	if req.IdentityKeyPub == "" || req.SignedPreKeyPub == "" {
		return nil, domain.ErrInvalidInput
	}

	keysJson, err := json.Marshal(req.OneTimePreKeys)
	if err != nil {
		return nil, err
	}

	device := &domain.Device{
		UserID:          req.UserID,
		IdentityKeyPub:  req.IdentityKeyPub,
		SignedPreKeyPub: req.SignedPreKeyPub,
		OneTimePreKeys:  keysJson,
	}
	device.ID = uuid.New()

	if err := s.repo.Create(ctx, device); err != nil {
		return nil, err
	}
	return dto.ToDeviceResponse(device), nil
}
func (s *DeviceService) GetUserDevices(ctx context.Context, userID uuid.UUID) ([]*dto.DeviceResponse, error) {
	devices, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	res := make([]*dto.DeviceResponse, len(devices))
	for i, device := range devices {
		res[i] = dto.ToDeviceResponse(device)
	}
	return res, nil
}
func (s *DeviceService) DeleteDevice(ctx context.Context, deviceID uuid.UUID) error {
	return s.repo.Delete(ctx, deviceID)
}
func (s *DeviceService) ClaimOneTimePreKey(ctx context.Context, deviceID uuid.UUID) (*dto.ClaimKeyResponse, error) {
	key, err := s.repo.ClaimOneTimePreKey(ctx, deviceID)
	if err != nil {
		return nil, err
	}
	return &dto.ClaimKeyResponse{Key: key}, nil
}
func (s *DeviceService) AddOneTimePreKeys(ctx context.Context, deviceID uuid.UUID, req dto.AddOneTimePreKeysRequest) error {
	if len(req.Keys) == 0 {
		return domain.ErrInvalidInput
	}
	return s.repo.AddOneTimePreKeys(ctx, deviceID, req.Keys)
}
