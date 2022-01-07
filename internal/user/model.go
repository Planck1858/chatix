package user

import (
	"database/sql"
	"githab.com/Planck1858/chatix/internal/storage/user"
	"githab.com/Planck1858/chatix/pkg/utils"
	"time"
)

/***** User *****/
type User struct {
	Id        string    `json:"id"`
	Login     string    `json:"login"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      Role      `json:"role"`
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
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: utils.ConvTimeToSqlNullTime(u.DeletedAt),
	}
}

/***** CreateUserDTO *****/
type CreateUserDTO struct {
	Login string `json:"login"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  Role   `json:"role"`
}

func (u *CreateUserDTO) ConvToRep() *user.User {
	return &user.User{
		Id:        "",
		Login:     u.Login,
		Name:      u.Name,
		Email:     u.Email,
		Role:      u.Role.String(),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: sql.NullTime{},
	}
}

/***** UpdateUserDTO *****/
type UpdateUserDTO struct {
	Id    string `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  Role   `json:"role"`
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