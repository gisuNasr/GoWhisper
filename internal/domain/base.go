package domain

import (
	"time"

	"github.com/google/uuid"
)

type BaseModel struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BaseModelWithSoftDelete struct {
	BaseModel
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
