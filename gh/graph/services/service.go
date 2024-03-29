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

type ProjectService interface {
	AddProjectItem(ctx context.Context, input *model.AddProjectV2ItemByIDInput) (*model.AddProjectV2ItemByIDPayload, error)
}

type Services interface {
	UserService
	RepositoryService
	ProjectService

	GetNodeByID(ctx context.Context, id string) (model.Node, error)
}

type services struct {
	*userService
	*repositoryService
	*issueService
	*pullRequestService
	*projectService
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		userService:        &userService{exec: exec},
		repositoryService:  &repositoryService{exec: exec},
		issueService:       &issueService{exec: exec},
		pullRequestService: &pullRequestService{exec: exec},
		projectService:     &projectService{exec: exec},
	}
}

func (s *services) GetNodeByID(ctx context.Context, id string) (model.Node, error) {
	t, err := getTypeOfNode(id)
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
	case nodeTypePullRequest:
		return s.getPullRequestByID(ctx, id)
	case nodeTypeProject:
		return s.getProjectByID(ctx, id)
	default:
		return nil, fmt.Errorf("id type %d is not supported", t)
	}
}
