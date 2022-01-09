package user

import (
	"database/sql"
	"time"
)

type User struct {
	Id        string       `db:"id"`
	Login     string       `db:"login"`
	Name      string       `db:"name"`
	Email     string       `db:"email"`
	Role      string       `db:"role"`
	Password  []byte       `db:"password"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
