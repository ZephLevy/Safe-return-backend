package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepo(db *pgx.Conn) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) IsEmailUnique(ctx context.Context, email string) (bool, error) {
	row, err := ur.db.Query(ctx, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return false, err
	}
	defer row.Close()

	return !row.Next(), nil
}
