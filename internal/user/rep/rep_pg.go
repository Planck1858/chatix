package rep

import (
	"context"
	utils "githab.com/Planck1858/chatix/pkg/utils/repository"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"time"
)

type repository struct {
	conn *sqlx.DB
}

func NewRepository(conn *sqlx.DB) Repository {
	r := &repository{conn: conn}
	return RepositoryWithLogger(r)
}

func (r *repository) GetAllUsers(ctx context.Context) ([]User, error) {
	return getAllUsers(ctx, r.conn)
}

func getAllUsers(ctx context.Context, d utils.SqlDriver) ([]User, error) {
	query := `
	SELECT
		id,
		login,
		name,
		phone,
		role,
		created_at,
		updated_at,
		deleted_at
	FROM chatix.user;`

	u := make([]User, 0)
	err := d.SelectContext(ctx, &u, query)

	return u, err
}

func (r *repository) GetUser(ctx context.Context, id string) (*User, error) {
	return getUser(ctx, r.conn, id)
}

func getUser(ctx context.Context, d utils.SqlDriver, id string) (*User, error) {
	query := `
	SELECT
		id,
		login,
		name,
		phone,
		role,
		created_at,
		updated_at,
		deleted_at
	FROM chatix.user WHERE id = $1 AND deleted_at is null;`

	var u User
	err := d.GetContext(ctx, &u, query, id)

	return &u, err
}

func (r *repository) CreateUser(ctx context.Context, user *User) (string, error) {
	return createUser(ctx, r.conn, user)
}

func createUser(ctx context.Context, d utils.SqlDriver, u *User) (id string, err error) {
	t := time.Now()
	u.Id = uuid.New().String()
	u.CreatedAt = t
	u.UpdatedAt = t

	query := `
	INSERT INTO chatix.user (
		id, 
		login,
		name,
		phone,
		role,
		created_at,
		updated_at
	)
	VALUES (
		:id, 
		:login,
		:name,
		:phone,
		:role,
		:created_at,
		:updated_at
	);`

	_, err = d.NamedExecContext(ctx, query, u)
	if err != nil {
		return "", err
	}

	return u.Id, nil
}

func (r *repository) UpdateUser(ctx context.Context, user *User) error {
	return updateUser(ctx, r.conn, user)
}

func updateUser(ctx context.Context, d utils.SqlDriver, user *User) error {
	user.UpdatedAt = time.Now()

	query := `
	UPDATE chatix.user SET
		login = :login,
		name = :name,
		phone = :phone,
		role = :role,
		updated_at = :updated_at
	WHERE id = :id;`

	_, err := d.NamedExecContext(ctx, query, user)

	return err
}

func (r *repository) DeleteUser(ctx context.Context, id string) error {
	return deleteUser(ctx, r.conn, id)
}

func deleteUser(ctx context.Context, d utils.SqlDriver, id string) error {
	deletedAt := time.Now()

	query := `
	UPDATE chatix.user SET
		deleted_at = $1
	WHERE id = $2;`

	_, err := d.ExecContext(ctx, query, deletedAt, id)

	return err
}
