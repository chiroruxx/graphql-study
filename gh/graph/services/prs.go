package services

import (
	"context"

	"github.com/chiroruxxxx/graphql-study/gh/graph/db"
	"github.com/chiroruxxxx/graphql-study/gh/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func convertPullRequest(pr *db.Pullrequest) *model.PullRequest {
	return &model.PullRequest{
		ID:          pr.ID,
		BaseRefName: pr.BaseRefName,
		//Closed:      pr.Closed,
		HeadRefName: pr.HeadRefName,
		URL:         pr.URL,
		//Number: pr.Number,
		//Repository: pr.Repository,
		//ProjectItems: pr.ProjectItems
	}
}

type pullRequestService struct {
	exec boil.ContextExecutor
}

func (p *pullRequestService) getPullRequestByID(ctx context.Context, id string) (*model.PullRequest, error) {
	pr, err := db.Pullrequests(
		p.selectColumns(),
		db.PullrequestWhere.ID.EQ(id),
	).One(ctx, p.exec)

	if err != nil {
		return nil, err
	}

	return convertPullRequest(pr), nil
}

func (p *pullRequestService) selectColumns() qm.QueryMod {
	return qm.Select(
		db.PullrequestColumns.ID,
		db.PullrequestColumns.BaseRefName,
		db.PullrequestColumns.URL,
		db.PullrequestColumns.HeadRefName,
		db.PullrequestColumns.Number,
		db.PullrequestColumns.Closed,
		db.PullrequestColumns.Repository,
	)
}
