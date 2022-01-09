package user

import (
	"database/sql"
	"githab.com/Planck1858/chatix/back-end/internal/storage/user"
	"githab.com/Planck1858/chatix/back-end/pkg/utils"
	"time"
)

/***** User *****/
type User struct {
	Id        string    `json:"id"`
	Login     string    `json:"login"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      Role      `json:"role"`
	Password  []byte    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type Role string

func (r Role) String() string {
	return string(r)
}

func (u *User) ConvToRep() *user.User {
	if u == nil {
		return nil
	}

	return &user.User{
		Id:        u.Id,
		Login:     u.Login,
		Name:      u.Name,
		Email:     u.Email,
		Role:      u.Role.String(),
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: utils.ConvTimeToSqlNullTime(u.DeletedAt),
	}
}

func ConvRepUserToServ(u *user.User) *User {
	if u == nil {
		return nil
	}

	return &User{
		Id:        u.Id,
		Login:     u.Login,
		Name:      u.Name,
		Email:     u.Email,
		Role:      Role(u.Role),
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: utils.ConvSqlNullTime(u.DeletedAt),
	}
}

func ConvRepUsersToServ(repU []user.User) []*User {
	servU := make([]*User, 0, len(repU))
	for i := range repU {
		u := &repU[i]
		servU = append(servU, ConvRepUserToServ(u))
	}

	return servU
}

/***** CreateUserDTO *****/
type CreateUserDTO struct {
	Login    string `validate:"required"`
	Name     string `validate:"required"`
	Email    string `validate:"required"`
	Role     Role   `validate:"required"`
	Password []byte `validate:"required"`
}

func (u *CreateUserDTO) ConvToRep() *user.User {
	return &user.User{
		Id:        "",
		Login:     u.Login,
		Name:      u.Name,
		Email:     u.Email,
		Role:      u.Role.String(),
		Password:  u.Password,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: sql.NullTime{},
	}
}

/***** UpdateUserDTO *****/
type UpdateUserDTO struct {
	Id    string `json:"id" validate:"required"`
	Login string `json:"login" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
	Role  Role   `json:"role" validate:"required"`
}

func (u *UpdateUserDTO) ConvToRep() *user.User {
	return &user.User{
		Id:        u.Id,
		Login:     u.Login,
		Name:      u.Name,
		Email:     u.Email,
		Role:      u.Role.String(),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: sql.NullTime{},
	}
}
