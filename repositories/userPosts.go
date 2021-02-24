package repositories

import (
	"errors"
	"github.com/HarrisonWAffel/dbTrain/config"
	"github.com/HarrisonWAffel/dbTrain/models"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserPostsRepository struct {
	db *gorm.DB
}

func NewUserPostsRepository(db... *gorm.DB) (*UserPostsRepository, error) {
	if len(db) > 0 {
		return &UserPostsRepository{
			db: db[0],
		}, nil
	} else {
		db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
		if err != nil {
			return &UserPostsRepository{}, err
		}
		return &UserPostsRepository{
			db: db,
		}, nil
	}
}

func (r *UserPostsRepository) GetUserPostByPostId(postId uuid.UUID) (models.UserPost, error) {
	var posts models.UserPost
	result := r.db.Find(&posts, "post_id = ?", postId)
	if result.Error != nil {
		return posts, nil
	}
	if result.RowsAffected == 0 {
		return posts, errors.New("no rows returned")
	}
	return posts, nil
}

func (r *UserPostsRepository) CreateUserPost(post models.UserPost) error {
	return r.db.Model(&post).Create(&post).Error
}

func (r *UserPostsRepository) UpdateUserPost(post models.UserPost) error {
	return r.db.Model(&post).Save(&post).Error
}

func (r *UserPostsRepository) DeleteUserPost(post models.UserPost) error {
	return r.db.Delete(&post).Error
}

func (r *UserPostsRepository) GetUserPostsForUser(user models.User) ([]models.UserPost, error) {
	var posts []models.UserPost
	result := r.db.Find(&posts, "user_id = ?", user.ID)
	if result.Error != nil {
		return nil, nil
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("no rows returned")
	}
	return posts, nil
}
