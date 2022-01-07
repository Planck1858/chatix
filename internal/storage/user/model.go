package user

import (
	"database/sql"
	"githab.com/Planck1858/chatix/internal/user"
	"githab.com/Planck1858/chatix/pkg/utils"
	"time"
)

type User struct {
	Id        string       `db:"id"`
	Login     string       `db:"login"`
	Name      string       `db:"name"`
	Email     string       `db:"email"`
	Role      string       `db:"role"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

func (u *User) ConvToServ() *user.User {
	if u == nil {
		return nil
	}

	return &user.User{
		Id:        u.Id,
		Login:     u.Login,
		Name:      u.Name,
		Email:     u.Email,
		Role:      user.Role(u.Role),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: utils.ConvSqlNullTime(u.DeletedAt),
	}
}

func ConvRepUsersToServUsers(repU []User) []*user.User {
	servU := make([]*user.User, 0, len(repU))
	for i := range repU {
		u := &repU[i]
		servU = append(servU, u.ConvToServ())
	}

	return servU
}
