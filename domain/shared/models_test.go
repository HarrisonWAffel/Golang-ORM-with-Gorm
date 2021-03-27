package shared

import (
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestPost_ToJSON(t *testing.T) {
	p := Post{
		BaseEntity: BaseEntity{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
		},
		Content:   "",
		Private:   false,
		CreatedOn: time.Now(),
	}
	t.Log(string(p.ToJSON()))
	if len(p.ToJSON()) == 0 {
		t.FailNow()
	}
}

func TestUserPost_ToJSON(t *testing.T) {
	p := UserPost{
		BaseEntity: BaseEntity{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
		},
		UserId:  uuid.New(),
		PostId:  uuid.New(),
		Private: false,
	}
	t.Log(string(p.ToJSON()))
	if len(p.ToJSON()) == 0 {
		t.FailNow()
	}
}
