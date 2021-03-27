package userPosts

import (
	"errors"
	"github.com/HarrisonWAffel/dbTrain/config"
	"github.com/HarrisonWAffel/dbTrain/user"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserPostsRepository struct {
	db *gorm.DB
}

func NewUserPostsRepository(db ...*gorm.DB) (*UserPostsRepository, error) {
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
			db: db.Model(&UserPost{}),
		}, nil
	}
}

func (r *UserPostsRepository) GetUserPostByPostId(postId uuid.UUID) (UserPost, error) {
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

func (r *UserPostsRepository) CreateUserPost(post UserPost) error {
	return r.db.Create(&post).Error
}

func (r *UserPostsRepository) UpdateUserPost(post UserPost) error {
	return r.db.Save(&post).Error
}

func (r *UserPostsRepository) DeleteUserPost(post UserPost) error {
	return r.db.Delete(&post).Error
}

func (r *UserPostsRepository) GetUserPostsForUser(user user.User) ([]UserPost, error) {
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
