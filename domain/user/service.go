package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service interface {
	GetUserByEmail(email string) (User, error)
	GetUserById(id uuid.UUID) (User, error)
	SaveUser(User) error
	DeleteUser(User) error
	UpdateUser(User) error
}

type service struct {
	repo *Repository
}

func NewService(dbconn *gorm.DB) (Service, error) {
	repo, err := NewUserRepository(dbconn)
	if err != nil {
		return nil, err
	}

	return &service{
		repo: repo,
	}, nil
}

func (s *service) GetUserByEmail(email string) (User, error) {
	return s.repo.FindUserByEmail(email)
}

func (s *service) GetUserById(id uuid.UUID) (User, error) {
	e, err := s.repo.GetById(id)
	var u User
	if err == nil {
		u = e.(User)
	}
	return u, err
}

func (s *service) SaveUser(u User) error {
	return s.repo.Create(u)
}

func (s *service) DeleteUser(u User) error {
	return s.repo.Delete(u)
}

func (s *service) UpdateUser(u User) error {
	return s.repo.Update(u)
}
