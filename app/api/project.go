package api

import (
	"coffeebeans-people-backend/models"
	"context"
)

func (apiSvc *ApiSvc) CreateProjectByAdmin(ctx context.Context, project models.Project) error {
	err := apiSvc.DbSvc.CreateProject(ctx, project)
	if err != nil {
		return err
	}

	return nil
}
