package domain

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

type Device struct {
	BaseModelWithSoftDelete
	UserID          int64           `json:"user_id"`
	IdentityKeyPub  string          `json:"identity_key_pub"`
	SignedPreKeyPub string          `json:"signed_pre_key_pub"`
	OneTimePreKeys  json.RawMessage `json:"one_time_pre_keys"`
}

type DeviceRepository interface {
	Create(ctx context.Context, device *Device) error
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*Device, error)
	Delete(ctx context.Context, deviceID uuid.UUID) error
	UpdateSignedPreKey(ctx context.Context, deviceID uuid.UUID, newPubKey string) error
	ClaimOneTimePreKey(ctx context.Context, deviceID uuid.UUID) (string, error)
	AddOneTimePreKeys(ctx context.Context, deviceID uuid.UUID, newKeys []string) error
}
