package services

import (
	"context"

	"github.com/chiroruxxxx/graphql-study/gh/graph/db"
	"github.com/chiroruxxxx/graphql-study/gh/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func convertRepository(repository *db.Repository) *model.Repository {
	return &model.Repository{
		ID: repository.ID,
		//Owner:        repository.Owner,
		Name:      repository.Name,
		CreatedAt: repository.CreatedAt,
		//Issue:        repository.Issues(),
		//Issues:       repository.Issues(),
		//PullRequest:  repository.Pullrequests(),
		//PullRequests: nil,
	}
}

type repositoryService struct {
	exec boil.ContextExecutor
}

func (r *repositoryService) GetRepositoryByNameAndOwner(ctx context.Context, name string, owner string) (*model.Repository, error) {
	repo, err := db.Repositories(
		r.selectColumns(),
		db.RepositoryWhere.Name.EQ(name),
		db.RepositoryWhere.Owner.EQ(owner),
	).One(ctx, r.exec)

	if err != nil {
		return nil, err
	}

	return convertRepository(repo), nil
}

func (r *repositoryService) getRepositoryByID(ctx context.Context, id string) (*model.Repository, error) {
	repo, err := db.Repositories(
		r.selectColumns(),
		db.RepositoryWhere.ID.EQ(id),
	).One(ctx, r.exec)

	if err != nil {
		return nil, err
	}

	return convertRepository(repo), nil
}

func (r *repositoryService) selectColumns() qm.QueryMod {
	return qm.Select(
		db.RepositoryColumns.ID,
		db.RepositoryColumns.Owner,
		db.RepositoryColumns.Name,
		db.RepositoryColumns.CreatedAt,
	)
}
