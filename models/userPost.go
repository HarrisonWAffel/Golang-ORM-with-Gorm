package models

import (
	"encoding/json"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UserPost struct {
	Model
	UserId uuid.UUID `json:"user_id" gorm:"type:uuid"`
	PostId uuid.UUID `json:"post_id" gorm:"type:uuid"`
	Private bool `json:"private"`
}

func (u *UserPost) ToJSON() []byte {
	j, _ := json.Marshal(u)
	return j
}

func (u *UserPost) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}