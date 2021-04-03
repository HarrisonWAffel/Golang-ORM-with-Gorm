package user

import (
	"errors"
	"github.com/HarrisonWAffel/dbTrain/config"
	"github.com/HarrisonWAffel/dbTrain/domain"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewUserRepository(db ...*gorm.DB) (*Repository, error) {
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
			db: db.Model(&User{}),
		}, nil
	}
}

func (r *Repository) GetById(id uuid.UUID) (domain.Entity, error) {
	var foundUser User
	result := r.db.Find(&foundUser, id)
	if result.Error != nil {
		return User{}, result.Error
	}
	return foundUser, nil
}

func (r *Repository) Create(entity domain.Entity) error {
	u := entity.(User)
	return r.db.Create(&u).Error
}

func (r *Repository) Update(entity domain.Entity) error {
	u := entity.(User)
	return r.db.Save(&u).Error
}

func (r *Repository) Delete(entity domain.Entity) error {
	u := entity.(User)
	return r.db.Delete(&u).Error
}

func (r *Repository) FindUserByEmail(email string) (User, error) {
	var foundUser User
	result := r.db.Find(&foundUser, "email = ?", email)
	if result.Error != nil {
		return User{}, result.Error
	}
	if result.RowsAffected == 0 {
		return User{}, errors.New("now rows returned")
	}
	return foundUser, nil
}

func (r *Repository) FindAllUsers() ([]User, error) {
	var allUsers []User
	result := r.db.Find(&allUsers)
	return allUsers, result.Error
}
