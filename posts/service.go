package posts

import (
	"github.com/HarrisonWAffel/dbTrain/models"
	"github.com/HarrisonWAffel/dbTrain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	repo *repositories.PostsRepository
}

func NewService(db *gorm.DB) (*Service, error) {
	repo, err := repositories.NewPostsRepository(db)
	if err != nil {
		return &Service{}, err
	}
	return &Service{
		repo: repo,
	}, nil
}

func (s *Service) CreatePost(post models.Post) (models.Post, error) {
	p, err := s.repo.CreatePost(post)
	if err != nil {
		return models.Post{}, err
	}

	return p, nil
}

func (s *Service) GetPostsByUserId(userId uuid.UUID) ([]models.Post, error) {
	return s.repo.GetPostsByUserId(userId)
}

func (s *Service) GetPostById(postId uuid.UUID) (models.Post, error) {
	return s.repo.GetPostById(postId)
}

func (s *Service) UpdatePost(post models.Post) error {
	return s.repo.UpdatePost(post)
}

func (s *Service) DeletePost(post models.Post) error {
	return s.repo.DeletePost(post)
}