package user

import "context"

type Service interface {
	GetAllUsers(ctx context.Context) ([]*User, error)
	GetUser(ctx context.Context, id string) (*User, error)
	CreateUser(ctx context.Context, user *CreateUserDTO) (id string, err error)
	UpdateUser(ctx context.Context, user *UpdateUserDTO) error
	DeleteUser(ctx context.Context, id string) error
}
