package domain

import "errors"

var (
	ErrNotFound         = errors.New("record not found")
	ErrAlreadyExists    = errors.New("record already exists")
	ErrUnauthorized     = errors.New("unauthorized")
	ErrInvalidInput     = errors.New("invalid input")
	ErrNoOneTimePreKeys = errors.New("no one-time pre-keys available for this device")
)
