package repositories

import (
	"github.com/HarrisonWAffel/dbTrain/config"
	"github.com/HarrisonWAffel/dbTrain/models"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostsRepository struct {
	db *gorm.DB
}

func NewPostsRepository(db ...*gorm.DB) (*PostsRepository, error) {
	if len(db) > 0 {
		return &PostsRepository{
			db: db[0],
		}, nil
	} else {
		db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
		if err != nil {
			return &PostsRepository{}, err
		}
		return &PostsRepository{
			db: db.Model(&models.Post{}),
		}, nil
	}
}

// CRUD!
func (r *PostsRepository) CreatePost(post models.Post) (models.Post, error) {
	err := r.db.Create(&post).Error
	return post, err
}

func (r *PostsRepository) UpdatePost(post models.Post) error {
	return r.db.Save(&post).Error
}

func (r *PostsRepository) DeletePost(post models.Post) error {
	return r.db.Delete(&post).Error
}

func (r *PostsRepository) GetPostById(postId uuid.UUID) (models.Post, error) {
	var post models.Post
	result := r.db.First(&post, "id = ?", postId)
	if result.Error != nil {
		return post, result.Error
	}
	return post, nil
}

func (r *PostsRepository) GetPostsByUserId(userId uuid.UUID) ([]models.Post, error) {
	var foundPost []models.Post
	result := r.db.Find(&foundPost, "user_id = ?", userId)
	if result.Error != nil {
		return []models.Post{}, result.Error
	}
	return foundPost, nil
}
