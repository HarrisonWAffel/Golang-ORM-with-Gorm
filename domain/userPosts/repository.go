package userPosts

import (
	"errors"
	"github.com/HarrisonWAffel/dbTrain/config"
	"github.com/HarrisonWAffel/dbTrain/domain"
	"github.com/HarrisonWAffel/dbTrain/domain/user"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewUserPostsRepository(db ...*gorm.DB) (*Repository, error) {
	if len(db) > 0 {
		return &Repository{
			db: db[0],
		}, nil
	} else {
		db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
		if err != nil {
			return &Repository{}, err
		}
		return &Repository{
			db: db.Model(&UserPost{}),
		}, nil
	}
}

func (r *Repository) GetById(id uuid.UUID) (domain.Entity, error) {
	var posts UserPost
	result := r.db.Find(&posts, "post_id = ?", id)
	if result.Error != nil {
		return posts, nil
	}
	if result.RowsAffected == 0 {
		return posts, errors.New("no rows returned")
	}
	return posts, nil
}

func (r *Repository) Create(entity domain.Entity) error {
	post := entity.(UserPost)
	return r.db.Create(&post).Error
}

func (r *Repository) Update(entity domain.Entity) error {
	post := entity.(UserPost)
	return r.db.Save(&post).Error
}

func (r *Repository) Delete(entity domain.Entity) error {
	post := entity.(UserPost)
	return r.db.Delete(&post).Error
}

func (r *Repository) GetUserPostByPostId(postId uuid.UUID) (UserPost, error) {
	var posts UserPost
	result := r.db.Find(&posts, "post_id = ?", postId)
	if result.Error != nil {
		return posts, nil
	}
	if result.RowsAffected == 0 {
		return posts, errors.New("no rows returned")
	}
	return posts, nil
}

func (r *Repository) GetUserPostsForUser(user user.User) ([]UserPost, error) {
	var posts []UserPost
	result := r.db.Find(&posts, "user_id = ?", user.ID)
	if result.Error != nil {
		return nil, nil
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("no rows returned")
	}
	return posts, nil
}
