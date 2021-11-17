package rep

import "time"

type User struct {
	Id        string    `db:"id"`
	Login     string    `db:"login"`
	Name      string    `db:"name"`
	Phone     string    `db:"phone"`
	Role      string    `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt time.Time `db:"deleted_at"`
}
