package posts

import (
	"github.com/HarrisonWAffel/dbTrain/config"
	"github.com/HarrisonWAffel/dbTrain/domain/shared"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func (r *Repository) GetById(id uuid.UUID) (shared.Entity, error) {
	var post Post
	result := r.db.First(&post, "id = ?", id)
	if result.Error != nil {
		return post, result.Error
	}
	return post, nil
}

func (r *Repository) Create(entity shared.BaseEntity) (shared.Entity, error) {
	panic("implement me")
}

func (r *Repository) Update(entity shared.BaseEntity) (shared.Entity, error) {
	panic("implement me")
}

func (r *Repository) Delete(entity shared.BaseEntity) error {
	panic("implement me")
}

func NewPostsRepository(db ...*gorm.DB) (*Repository, error) {
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
			db: db.Model(&Post{}),
		}, nil
	}
}

// CRUD!
func (r *Repository) CreatePost(post Post) (Post, error) {
	err := r.db.Create(&post).Error
	return post, err
}

func (r *Repository) UpdatePost(post Post) error {
	return r.db.Save(&post).Error
}

func (r *Repository) DeletePost(post Post) error {
	return r.db.Delete(&post).Error
}

func (r *Repository) GetPostById(postId uuid.UUID) (Post, error) {
	var post Post
	result := r.db.First(&post, "id = ?", postId)
	if result.Error != nil {
		return post, result.Error
	}
	return post, nil
}

func (r *Repository) GetPostsByUserId(userId uuid.UUID) ([]Post, error) {
	var foundPost []Post
	result := r.db.Find(&foundPost, "user_id = ?", userId)
	if result.Error != nil {
		return []Post{}, result.Error
	}
	return foundPost, nil
}
