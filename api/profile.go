package api

import (
	"coffeebeans-people-backend/models"
	"context"
)

func (apiSvc *ApiSvc) EditUser(ctx context.Context, user models.User) error {
	err := apiSvc.DbSvc.UpdateUserProfile(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
