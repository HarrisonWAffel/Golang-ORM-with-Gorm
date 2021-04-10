package user

import (
	"encoding/json"
	"github.com/HarrisonWAffel/dbTrain/domain"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type User struct {
	domain.BaseEntity
	UserName  string    `json:"user_name,omitempty"`
	Password  string    `json:"password,omitempty"`
	Email     string    `json:"email,omitempty"`
	LastLogin time.Time `json:"last_login,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}

func (u *User) Verify() error {
	retErr := errors.New("model verification error")

	if u.UserName == "" {
		return errors.Wrap(retErr, "empty username")
	}

	if u.Password == "" {
		return errors.Wrap(retErr, "empty password")
	}

	if u.Email == "" {
		return errors.Wrap(retErr, "empty email")
	}

	return nil
}

func (u User) ToJSON() []byte {
	j, _ := json.Marshal(u)
	return j
}
