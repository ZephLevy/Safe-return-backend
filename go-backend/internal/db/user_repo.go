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

func (ur *UserRepository) CreateAccount(ctx context.Context, email string, passwordHash string) (string, error) {
	query := "INSERT INTO users (email, password_hash) VALUES ($1, $2)"
	_, err := ur.db.Exec(ctx, query, email, passwordHash)
	if err != nil {
		return "", err
	}
	// TODO: Generate bearer and refresh tokens
	return "randomstringblablabla", nil
}
