package posts

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service interface {
	GetPostById(id uuid.UUID) (Post, error)
	CreateNewPost(post Post) error
	UpdatePost(post Post) error
	DeletePost(post Post) error
}

type service struct {
	repo *Repository
}

func NewService(dbConn *gorm.DB) (Service, error) {
	postRepo, err := NewPostsRepository(dbConn)
	if err != nil {
		return nil, err
	}
	return &service{repo: postRepo}, nil
}

func (s *service) GetPostById(id uuid.UUID) (Post, error) {
	e, err := s.repo.GetById(id)
	var p Post
	if err == nil {
		p = e.(Post)
	}
	return p, err
}

func (s *service) DeletePost(post Post) error {
	return s.repo.Delete(post)
}

func (s *service) CreateNewPost(post Post) error {
	if post.ID == uuid.Nil {
		post.ID = uuid.New()
	}
	return s.repo.Create(post)
}

func (s *service) UpdatePost(post Post) error {
	return s.repo.Update(post)
}
