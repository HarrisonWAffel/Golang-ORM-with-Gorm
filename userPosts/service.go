package userPosts

import (
	"github.com/HarrisonWAffel/dbTrain/models"
	"github.com/HarrisonWAffel/dbTrain/posts"
	"github.com/HarrisonWAffel/dbTrain/repositories"
	"github.com/HarrisonWAffel/dbTrain/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

type Service struct {
	repo *repositories.UserPostsRepository
	userService *user.Service
	postsService *posts.Service
}

func NewService(db *gorm.DB, userService *user.Service, postsService *posts.Service) (*Service, error) {
	repo, err := repositories.NewUserPostsRepository(db)
	if err != nil {
		return &Service{}, err
	}

	return &Service{
		repo: repo,
		userService: userService,
		postsService: postsService,
	}, nil
}

func (s *Service) CreateNewPost(post models.Post, email string) error {

	u, err := s.userService.GetUserByEmail(email)
	if  err != nil {
		return err
	}

	post, err = s.postsService.CreatePost(post)
	if err != nil {
		return err
	}

	userPost := models.UserPost {
		UserId:  u.ID,
		PostId:  post.ID,
		Private: post.Private,
	}

	err = s.repo.CreateUserPost(userPost)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func (s *Service) GetUserPostsByUserId(user models.User) ([]models.Post, error) {

	dbUser, err := s.userService.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	idList, err := s.repo.GetUserPostsForUser(dbUser)
	if err != nil {
		return nil, err
	}

	var userPosts []models.Post
	for _, e := range idList {
		p, err := s.postsService.GetPostById(e.PostId)
		if err != nil {
			return nil, err
		}
		userPosts = append(userPosts, p)
	}


	return userPosts, nil
}

func (s *Service) RemoveUserPostByPostId(postId uuid.UUID) error {
	p, err := s.repo.GetUserPostByPostId(postId)
	if err != nil {
		return err
	}

	return s.repo.DeleteUserPost(p)
}