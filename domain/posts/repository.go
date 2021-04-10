package posts

import (
	"fmt"
	"github.com/HarrisonWAffel/dbTrain/config"
	"github.com/HarrisonWAffel/dbTrain/domain"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewPostsRepository(db ...*gorm.DB) (*Repository, error) {
	if len(db) > 0 {
		return &Repository{
			db: db[0],
		}, nil
	} else {
		config.Read()
		fmt.Println(config.Dsn)
		db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
		if err != nil {
			return &Repository{}, err
		}
		return &Repository{
			db: db.Model(&Post{}),
		}, nil
	}
}

func (r *Repository) GetById(id uuid.UUID) (domain.Entity, error) {
	var post Post
	result := r.db.First(&post, "id = ?", id)
	if result.Error != nil {
		return post, result.Error
	}
	return post, nil
}

func (r *Repository) Create(entity domain.Entity) error {
	post := entity.(Post)
	err := r.db.Create(&post).Error
	return err
}

func (r *Repository) Update(entity domain.Entity) error {
	post := entity.(Post)
	return r.db.Save(&post).Error
}

func (r *Repository) Delete(entity domain.Entity) error {
	post := entity.(Post)
	return r.db.Delete(&post).Error
}
