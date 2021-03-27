package user

import (
	"errors"
	"github.com/HarrisonWAffel/dbTrain/config"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db ...*gorm.DB) (*UserRepository, error) {
	if len(db) > 0 {
		return &UserRepository{
			db: db[0],
		}, nil
	} else {
		db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
		if err != nil {
			return &UserRepository{}, err
		}
		return &UserRepository{
			db: db.Model(&User{}),
		}, nil
	}
}

// CRUD!

func (r *UserRepository) CreateUser(user User) error {
	return r.db.Create(&user).Error
}

func (r *UserRepository) UpdateUser(user User) error {
	return r.db.Save(&user).Error
}

func (r *UserRepository) DeleteUser(user User) error {
	return r.db.Delete(&user).Error
}

func (r *UserRepository) FindUserById(userId uuid.UUID) (User, error) {
	var foundUser User
	result := r.db.Find(&foundUser, userId)
	if result.Error != nil {
		return User{}, result.Error
	}
	return foundUser, nil
}

func (r *UserRepository) FindUserByEmail(email string) (User, error) {
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

func (r *UserRepository) FindAllUsers() ([]User, error) {
	var allUsers []User
	result := r.db.Find(&allUsers)
	return allUsers, result.Error
}
