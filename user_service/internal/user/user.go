package user

type service struct{}

type User struct {
	Id       string `json:"id"`
	Login    string `json:"login"`
	FullName string `json:"fullName"`
}

type Service interface {
	GetUser() *User
}

func New() *service {
	return &service{}
}

func (s *service) GetUser() *User {
	return &User{
		Id:       "1",
		Login:    "ilya",
		FullName: "Ilya Ilya Ilya",
	}
}
