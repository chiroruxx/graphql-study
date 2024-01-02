package services

import (
	"context"

	"github.com/chiroruxxxx/graphql-study/gh/graph/db"
	"github.com/chiroruxxxx/graphql-study/gh/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func convertProject(project *db.Project) *model.ProjectV2 {
	return &model.ProjectV2{
		ID:    project.ID,
		Title: project.Title,
		URL:   project.URL,
		//Number: project.Number,
		//Items:  project.Items,
		//Owner:  project.Owner,
	}
}

type projectService struct {
	exec boil.ContextExecutor
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
