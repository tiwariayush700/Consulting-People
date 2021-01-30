package dao

import (
	"coffeebeans-people-backend/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type KafkaService struct {

}

func (k KafkaService)  CreateUser(ctx context.Context, user models.User) error {
	return nil
}

func (service *Service) CreateUser(ctx context.Context, user models.User) error {

	c := service.MongoConn.Collection("users")
	model := mongo.IndexModel{
		Keys: bson.M{
			"employee_id": user.EmployeeId,
		},
		Options: options.Index().SetUnique(true),
	}
	c.Indexes().CreateOne(ctx, model)

	_, err := c.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (service *Service) GetUserByEmployeeId(ctx context.Context, employeeId int64) (models.User, error) {
	user := models.User{}

	collection := service.MongoConn.Collection("users")

	doc := collection.FindOne(ctx, bson.M{"employee_id": employeeId})

	err := doc.Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (service *Service) GetUserByCredentials(ctx context.Context, email string, password string) (models.User, error) {
	user := models.User{}

	collection := service.MongoConn.Collection("users")

	doc := collection.FindOne(ctx, bson.M{"email": email, "password": password})

	err := doc.Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (service *Service) GetAllUsers(ctx context.Context, params map[string]interface{}) ([]models.User, error) {
	users := make([]models.User, 0)
	collection := service.MongoConn.Collection("users")

	filter := bson.M{}

	if params != nil {
		if param, ok := params["skill"]; ok {
			filter["skill"] = param
		} else if param, ok := params["id"]; ok {
			filter["employee_id"] = param
		}
	}

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return users, err
	}
	defer cur.Close(ctx)

	for cur.Next(context.Background()) {
		user := models.User{}

		err := cur.Decode(&user)
		if err != nil {
			return users, err
		}

		users = append(users, user)
	}
	if err := cur.Err(); err != nil {
		return users, err
	}

	return users, nil
}
