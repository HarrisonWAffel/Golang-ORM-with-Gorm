package user

import (
	"github.com/HarrisonWAffel/dbTrain/models"
	"github.com/HarrisonWAffel/dbTrain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	repo *repositories.UserRepository
}

func NewService(db *gorm.DB) (*Service, error) {
	repo, err := repositories.NewUserRepository(db)
	if err != nil {
		return &Service{}, err
	}

	return &Service{
		repo: repo,
	}, nil
}

func (s *Service) GetAllUsers() []models.User {
	users, _ := s.repo.FindAllUsers()
	return users
}

func (s *Service) GetUserById(userId uuid.UUID) (models.User, error) {
	user, err := s.repo.FindUserById(userId)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *Service) GetUserByEmail(email string) (models.User, error) {
	return s.repo.FindUserByEmail(email)
}

func (s *Service) SaveUser(user models.User) error {
	return s.repo.CreateUser(user)
}

func (s *Service) UpdateUser(user models.User) error {
	return s.repo.UpdateUser(user)
}

func (s *Service) DeleteUser(user models.User) error {

	if user.Verify() != nil { //if we only have an id get the full model
		var err error
		user, err = s.repo.FindUserById(user.ID)
		if err != nil {
			return err
		}
	}

	return s.repo.DeleteUser(user)
}
