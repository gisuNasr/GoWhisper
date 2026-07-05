package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BaseModelWithSoftDelete struct {
	BaseModel
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
