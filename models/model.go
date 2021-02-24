package models

import (
	"github.com/google/uuid"
	"time"
)

type Model struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"created_at,omitempty"`
	DeletedAt *time.Time `sql:"index"`
}