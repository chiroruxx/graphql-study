package services

import (
	"context"
	"crypto/sha1"
	"fmt"

	"github.com/chiroruxxxx/graphql-study/gh/graph/db"
	"github.com/chiroruxxxx/graphql-study/gh/graph/model"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func convertProject(project *db.Project) *model.ProjectV2 {
	return &model.ProjectV2{
		ID:    project.ID,
		Title: project.Title,
		URL:   mustParseURL(project.URL),
		//Number: project.Number,
		//Items:  project.Items,
		//Owner:  project.Owner,
	}
}

func convertProjectItem(item *db.Projectcard) *model.ProjectV2Item {
	return &model.ProjectV2Item{
		ID: item.ID,
		//Project: item.Project,
		//Content: item.Issue,
	}
}

type projectService struct {
	exec boil.ContextExecutor
}

func (p *projectService) AddProjectItem(ctx context.Context, input *model.AddProjectV2ItemByIDInput) (*model.AddProjectV2ItemByIDPayload, error) {
	var item db.Projectcard
	item.ID = p.calcID(input.ProjectID, input.ContentID)
	item.Project = input.ProjectID

	contentType, err := getTypeOfNode(input.ContentID)
	if err != nil {
		return nil, err
	}

	switch contentType {
	case nodeTypeIssue:
		item.Issue = null.StringFrom(input.ContentID)
	case nodeTypePullRequest:
		item.Pullrequest = null.StringFrom(input.ContentID)
	default:
		return nil, fmt.Errorf("id type %d is not supported", contentType)
	}

	if err := item.Insert(ctx, p.exec, boil.Infer()); err != nil {
		return nil, err
	}

	return &model.AddProjectV2ItemByIDPayload{Item: convertProjectItem(&item)}, nil
}

func (p *projectService) calcID(projectID, contentID string) string {
	id := fmt.Sprintf("%s_%s", projectID, contentID)
	res := sha1.Sum([]byte(id))
	return fmt.Sprintf("PVTI_%x", res)
}

func (p *projectService) getProjectByID(ctx context.Context, id string) (*model.ProjectV2, error) {
	project, err := db.Projects(
		p.selectColumns(),
		db.ProjectWhere.ID.EQ(id),
	).One(ctx, p.exec)

	if err != nil {
		return nil, err
	}

	return convertProject(project), nil
}

func (p *projectService) selectColumns() qm.QueryMod {
	return qm.Select(
		db.ProjectColumns.ID,
		db.ProjectColumns.Title,
		db.ProjectColumns.URL,
		db.ProjectColumns.Owner,
	)
}
