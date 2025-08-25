package db

import (
	"context"
	"strconv"
	"time"

	"github.com/ZephLevy/Safe-return-backend/internal/auth"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
)

type UserRepository struct {
	db    *pgx.Conn
	cache *redis.Client
}

func NewUserRepo(db *pgx.Conn) *UserRepository {
	cache := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	return &UserRepository{db: db, cache: cache}
}

func (ur *UserRepository) IsEmailUnique(ctx context.Context, email string) (bool, error) {
	row, err := ur.db.Query(ctx, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return false, err
	}
	defer row.Close()

	return !row.Next(), nil
}

func (ur *UserRepository) CreateAccount(ctx context.Context,
	firstName string,
	lastName string,
	email string,
	passwordHash string,
) (string, string, error) {

	tx, err := ur.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return "", "", err
	}
	defer tx.Rollback(ctx)

	var userID int
	err = tx.QueryRow(ctx, `
		INSERT INTO users (first_name, last_name, email, password_hash)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`, firstName, lastName, email, passwordHash).Scan(&userID)
	if err != nil {
		return "", "", err
	}

	accessToken, refreshToken, err := auth.GenerateTokens(strconv.Itoa(userID))
	if err != nil {
		return "", "", err
	}

	// TODO: Hash the refresh token
	_, err = tx.Exec(ctx, `
			UPDATE users SET refresh_token=$1 WHERE id=$2
		`, refreshToken, userID)
	if err != nil {
		return "", "", err
	}

	if err := tx.Commit(ctx); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (ur *UserRepository) SetEmailOTP(ctx context.Context, email string, passcode string) {
	ur.cache.Set(ctx, email, passcode, time.Minute*5)
}

func (ur *UserRepository) VerifyEmailOTP(
	ctx context.Context,
	email string,
	potentialPass string,
) (bool, error) {
	pass, err := ur.cache.Get(ctx, email).Result()
	if err != nil {
		return false, err
	}

	return pass == potentialPass, nil
}
