package services

import (
	"context"
	"fmt"

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

	GetNodeByID(ctx context.Context, id string) (model.Node, error)
}

type services struct {
	*userService
	*repositoryService
	*issueService
	*nodeService
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		userService:       &userService{exec: exec},
		repositoryService: &repositoryService{exec: exec},
		issueService:      &issueService{exec: exec},
		nodeService:       &nodeService{},
	}
}

func (s *services) GetNodeByID(ctx context.Context, id string) (model.Node, error) {
	t, err := s.getTypeOfNode(id)
	if err != nil {
		return nil, err
	}

	switch t {
	case nodeTypeUser:
		return s.getUserByID(ctx, id)
	case nodeTypeRepository:
		return s.getRepositoryByID(ctx, id)
	case nodeTypeIssue:
		return s.getIssueByID(ctx, id)
	}

	return nil, fmt.Errorf("id type %d is not supported", t)
}
