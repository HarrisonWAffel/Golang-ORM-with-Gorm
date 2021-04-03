package domain

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type BaseEntity struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"created_at,omitempty"`
	DeletedAt *time.Time `sql:"index"`
}

func (e BaseEntity) ToJSON() []byte {
	j, _ := json.Marshal(e)
	return j
}

type Entity interface {
	ToJSON() []byte
}
