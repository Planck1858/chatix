package user

import "context"

type Service interface {
	GetUser(ctx context.Context, id string) (*User, error)
}
