package repository

import (
	"cats-social/model/database"
	"context"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepositoryInterface {
	return &UserRepository{db}
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, nip string) (response database.User, err error) {
	err = r.db.QueryRowContext(ctx, "SELECT id, name, nip, password FROM users WHERE nip = $1", nip).Scan(&response.Id, &response.Name, &response.Nip, &response.Password)
	if err != nil {
		return
	}
	return
}

func (r *UserRepository) CreateUser(ctx context.Context, data database.User) (err error) {
	query := `
	INSERT INTO users (email, name, password, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id`

	_, err = r.db.ExecContext(
		ctx,
		query,
		data.Nip,
		data.Name,
		data.Password,
		data.CreatedAt,
		data.UpdatedAt,
	)

	return err
}
