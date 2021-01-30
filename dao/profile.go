package dao

import (
	"coffeebeans-people-backend/constants"
	"coffeebeans-people-backend/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func (service *Service) UpdateUserProfile(ctx context.Context, user models.User) error {
	c := service.MongoConn.Collection("users")
	userFromContext := ctx.Value(constants.USER_KEY).(models.User)

	filter := bson.M{}
	filter["employee_id"] = userFromContext.EmployeeId

	update := bson.M{"$set": user}

	_, err := c.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		return err
	}

	return nil
}
