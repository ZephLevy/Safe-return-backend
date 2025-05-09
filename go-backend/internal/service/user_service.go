package service

import (
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/ZephLevy/Safe-return-backend/internal/db"
)

type UserService struct {
	repo *db.UserRepository
}

func NewUserService(ur *db.UserRepository) *UserService {
	return &UserService{repo: ur}
}

func (us *UserService) SignIn(ctx context.Context, email string, password string) error {
	if email == "" || password == "" {
		return fmt.Errorf("Missing fields")
	}

	// TODO: Use bcrypt for hashing
	hashedPassword := sha256.Sum256([]byte(password))
	_ = hashedPassword
	isUnique, err := us.repo.IsEmailUnique(ctx, email)
	if err != nil {
		return fmt.Errorf("db error: %w", err)
	}
	if !isUnique {
		return fmt.Errorf("Email already in use")
	}
	// TODO: create account
	return nil
}
