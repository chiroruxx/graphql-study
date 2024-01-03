package services

import (
	"context"

	"github.com/chiroruxxxx/graphql-study/gh/graph/db"
	"github.com/chiroruxxxx/graphql-study/gh/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func convertIssue(issue *db.Issue) *model.Issue {
	return &model.Issue{
		ID:    issue.ID,
		URL:   mustParseURL(issue.URL),
		Title: issue.Title,
		//Closed: issue.Closed,
		//Number: issue.Number,
		//Author: issue.Author,
		//Repository: issue.Repository,
		//ProjectItems: issue.ProjectItems,
	}
}

type issueService struct {
	exec boil.ContextExecutor
}

func (i *issueService) getIssueByID(ctx context.Context, id string) (*model.Issue, error) {
	issue, err := db.Issues(
		i.selectColumns(),
		db.IssueWhere.ID.EQ(id),
	).One(ctx, i.exec)

	if err != nil {
		return nil, err
	}

	return convertIssue(issue), nil
}

func (i *issueService) selectColumns() qm.QueryMod {
	return qm.Select(
		db.IssueColumns.ID,
		db.IssueColumns.Title,
		db.IssueColumns.URL,
		db.IssueColumns.Number,
		db.IssueColumns.Closed,
		db.IssueColumns.Repository,
	)
}
