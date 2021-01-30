package dao

import (
	"coffeebeans-people-backend/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (service *Service) CreateProject(ctx context.Context, project models.Project) error {

	c := service.MongoConn.Collection("projects")
	model := mongo.IndexModel{
		Keys: bson.M{
			"project_id": project.ProjectId,
		},
		Options: options.Index().SetUnique(true),
	}

	c.Indexes().CreateOne(ctx, model)
	_, err := c.InsertOne(ctx, project)
	if err != nil {
		return err
	}

	return nil
}
