package rep

import "context"

type Repository interface {
	GetAllUsers(ctx context.Context) ([]User, error)
	GetUser(ctx context.Context, id string) (*User, error)
	CreateUser(ctx context.Context, user *User) (id string, err error)
	UpdateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, id string) error
}
