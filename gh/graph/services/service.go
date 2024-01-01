package services

import (
	"context"

	"github.com/chiroruxxxx/graphql-study/gh/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
}

type RepositoryService interface {
	GetRepositoryByNameAndOwner(ctx context.Context, name, owner string) (*model.Repository, error)
}

type Services interface {
	UserService
	RepositoryService
}

type services struct {
	*userService
	*repositoryService
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		userService:       &userService{exec: exec},
		repositoryService: &repositoryService{exec: exec},
	}
}
