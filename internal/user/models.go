package user

import "context"

type service struct{}

type User struct {
	Id       string `json:"id"`
	Login    string `json:"login"`
	FullName string `json:"fullName"`
}

func NewService() Service {
	return &service{}
}

func (s *service) GetUser(ctx context.Context, id string) (*User, error) {
	return &User{
		Id:       id,
		Login:    "reee",
		FullName: "reee",
	}, nil
}
