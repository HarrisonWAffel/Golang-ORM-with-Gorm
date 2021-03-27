package posts

import (
	"encoding/json"
	"github.com/HarrisonWAffel/dbTrain/domain/shared"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Post struct {
	shared.BaseEntity
	Content string `json:"content"`
	Private bool   `json:"is_private"`
}

func (p Post) ToJSON() []byte {
	j, _ := json.Marshal(p)
	return j
}

func (u *Post) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}
