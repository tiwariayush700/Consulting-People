package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Service struct {
	MongoConn *mongo.Database
}

func NewService(ctx context.Context, mongoUri string, mongoDbName string) (*Service, error) {
	mongoConn, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	err = mongoConn.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &Service{
		MongoConn: mongoConn.Database(mongoDbName),
	}, nil
}
