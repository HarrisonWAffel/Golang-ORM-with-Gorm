package userPosts

import (
	"github.com/HarrisonWAffel/dbTrain/domain/posts"
	"github.com/HarrisonWAffel/dbTrain/domain/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

type Service interface {
	GetUserPostsByUserId(user user.User) ([]posts.Post, error)
	CreateNewPost(post posts.Post, email string) error
	UpdatePost(post posts.Post) error
	RemoveUserPostByPostId(post posts.Post) error
}

type service struct {
	userService   user.Service
	postsService  posts.Service
	userPostsRepo *Repository
}

func NewService(db *gorm.DB, userService user.Service, postsService posts.Service) (Service, error) {
	repo, err := NewUserPostsRepository(db)
	if err != nil {
		return &service{}, err
	}

	return &service{
		userPostsRepo: repo,
		userService:   userService,
		postsService:  postsService,
	}, nil
}

func (s *service) UpdatePost(post posts.Post) error {
	return s.postsService.UpdatePost(post)
}

func (s *service) CreateNewPost(post posts.Post, email string) error {

	u, err := s.userService.GetUserByEmail(email)
	if err != nil {
		return err
	}

	post.ID = uuid.New()

	err = s.postsService.CreateNewPost(post)
	if err != nil {
		return err
	}

	userPost := UserPost{
		UserId:  u.ID,
		PostId:  post.ID,
		Private: post.Private,
	}

	err = s.userPostsRepo.Create(userPost)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func (s *service) GetUserPostsByUserId(user user.User) ([]posts.Post, error) {

	dbUser, err := s.userService.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	idList, err := s.userPostsRepo.GetUserPostsForUser(dbUser)
	if err != nil {
		return nil, err
	}

	var userPosts []posts.Post
	for _, userPost := range idList {
		p, err := s.postsService.GetPostById(userPost.PostId)
		if err != nil {
			return nil, err
		}
		userPosts = append(userPosts, p)
	}

	return userPosts, nil
}

func (s *service) RemoveUserPostByPostId(payload posts.Post) error {
	p, err := s.userPostsRepo.GetUserPostByPostId(payload.ID)
	if err != nil {
		return err
	}

	err = s.userPostsRepo.Delete(p)
	if err != nil {
		return err
	}

	err = s.postsService.DeletePost(payload)
	if err != nil {
		return err
	}
	return nil
}
