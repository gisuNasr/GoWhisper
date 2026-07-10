package dto

import (
	"github.com/gisuNasr/GoWhisper/internal/domain"
	"github.com/google/uuid"
)

type RegisterDeviceRequest struct {
	UserID          uuid.UUID `json:"user_id"`
	IdentityKeyPub  string    `json:"identity_key_pub"`
	SignedPreKeyPub string    `json:"signed_pre_key_pub"`
	OneTimePreKeys  []string  `json:"one_time_pre_keys"`
}

type DeviceResponse struct {
	ID              uuid.UUID `json:"id"`
	UserID          uuid.UUID `json:"user_id"`
	IdentityKeyPub  string    `json:"identity_key_pub"`
	SignedPreKeyPub string    `json:"signed_pre_key_pub"`
}

type AddOneTimePreKeysRequest struct {
	Keys []string `json:"keys"`
}

type ClaimKeyResponse struct {
	Key string `json:"one_time_pre_key"`
}

func ToDeviceResponse(d *domain.Device) *DeviceResponse {
	return &DeviceResponse{
		ID:              d.ID,
		UserID:          d.UserID,
		IdentityKeyPub:  d.IdentityKeyPub,
		SignedPreKeyPub: d.SignedPreKeyPub,
	}
}
