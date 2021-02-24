package util

import (
	"github.com/HarrisonWAffel/dbTrain/config"
	"github.com/HarrisonWAffel/dbTrain/posts"
	"github.com/HarrisonWAffel/dbTrain/user"
	"github.com/HarrisonWAffel/dbTrain/userPosts"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AppCtx struct {
	UserService      *user.Service
	PostsService     *posts.Service
	UserPostsService *userPosts.Service
}

func NewServiceContext() (*AppCtx, error) {
	err := config.Read()
	if err != nil {
		return &AppCtx{}, err
	}
	
	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		return &AppCtx{}, err
	}

	userService, err := user.NewService(db)
	if err != nil {
		return &AppCtx{}, err
	}

	postsService, err := posts.NewService(db)
	if err != nil {
		return &AppCtx{}, err
	}

	userPostsService, err := userPosts.NewService(db, userService, postsService)
	if err != nil {
		return &AppCtx{}, err
	}


	return &AppCtx{
		UserService:      userService,
		PostsService:     postsService,
		UserPostsService: userPostsService,
	}, nil
}
