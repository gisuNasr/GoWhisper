package domain

import "encoding/json"

type Device struct {
	BaseModel
	UserID          int64           `json:"user_id"`
	IdentityKeyPub  string          `json:"identity_key_pub"`
	SignedPreKeyPub string          `json:"signed_pre_key_pub"`
	OneTimePreKeys  json.RawMessage `json:"one_time_pre_keys"`
}
