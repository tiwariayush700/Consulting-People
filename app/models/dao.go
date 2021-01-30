package models

import (
	"context"
)

type Dao interface {
	CreateUser(ctx context.Context, user User) error
	GetUserByEmployeeId(ctx context.Context, employeeId int64) (User, error)
	GetUserByCredentials(ctx context.Context, email string, password string) (User, error)
	UpdateUserProfile(ctx context.Context, user User) error
	CreateProject(ctx context.Context, project Project) error
	GetAllUsers(ctx context.Context, params map[string]interface{}) ([]User, error)
}

type ApiSvc interface {
	RegisterUser(ctx context.Context, user User) error
	LoginUser(ctx context.Context, email string, password string) (User, bool, error)
	EditUser(ctx context.Context, user User) error
	CreateProjectByAdmin(ctx context.Context, project Project) error
	GetUsers(ctx context.Context, params map[string]interface{}) ([]User, error)
}
