package auth

import (
	"coffeebeans-people-backend/models"
	"context"
)

type Service struct {
	SecretKey string
}

type AuthSvc interface {
	GenerateToken(user *models.User) (string, error)
	AuthenticateToken(jwtTokenString string) (*models.User, string, error)
}

func NewService(ctx context.Context, secretKey string) (*Service, error) {
	return &Service{
		SecretKey: secretKey,
	}, nil
}
