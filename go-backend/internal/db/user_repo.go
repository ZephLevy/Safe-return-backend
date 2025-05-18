package db

import (
	"context"
	"time"

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
) (string, error) {
	query := `
		INSERT INTO users (first_name, last_name, email, password_hash)
	 	VALUES ($1, $2, $3, $4)
	 `
	_, err := ur.db.Exec(ctx, query, firstName, lastName, email, passwordHash)
	if err != nil {
		return "", err
	}
	// TODO: Generate bearer and refresh tokens
	return "randomstringblablabla", nil
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
