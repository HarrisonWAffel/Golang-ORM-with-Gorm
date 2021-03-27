package userPosts

import (
	"github.com/HarrisonWAffel/dbTrain/domain/posts"
	"github.com/HarrisonWAffel/dbTrain/domain/user"
	"gorm.io/gorm"
	"log"
)

type Service struct {
	repo *UserPostsRepository
}

func NewService(db *gorm.DB) (*Service, error) {
	repo, err := NewUserPostsRepository(db)
	if err != nil {
		return &Service{}, err
	}

	return &Service{
		repo: repo,
	}, nil
}

func (s *Service) CreateNewPost(post posts.Post, email string) error {

	u, err := s.userService.GetUserByEmail(email)
	if err != nil {
		return err
	}

	post, err = s.postsService.CreatePost(post)
	if err != nil {
		return err
	}

	userPost := UserPost{
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

func (s *Service) GetUserPostsByUserId(user user.User) ([]posts.Post, error) {

	dbUser, err := s.userService.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	idList, err := s.repo.GetUserPostsForUser(dbUser)
	if err != nil {
		return nil, err
	}

	var userPosts []posts.Post
	for _, e := range idList {
		p, err := s.postsService.GetPostById(e.ID)
		if err != nil {
			return nil, err
		}
		userPosts = append(userPosts, p)
	}

	return userPosts, nil
}

func (s *Service) RemoveUserPostByPostId(payload posts.Post) error {
	p, err := s.repo.GetUserPostByPostId(payload.ID)
	if err != nil {
		return err
	}

	err = s.repo.DeleteUserPost(p)
	if err != nil {
		return err
	}

	err = s.postsService.DeletePost(payload)
	if err != nil {
		return err
	}
	return nil
}
