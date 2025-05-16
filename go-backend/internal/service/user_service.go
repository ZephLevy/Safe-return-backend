package service

import (
	"context"
	"fmt"

	"github.com/ZephLevy/Safe-return-backend/internal/db"
	"golang.org/x/crypto/bcrypt"
)

const (
	HashCost = 12
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), HashCost)
	isUnique, err := us.repo.IsEmailUnique(ctx, email)
	if err != nil {
		return fmt.Errorf("db error: %w", err)
	}
	if !isUnique {
		return fmt.Errorf("Email already in use")
	}

	token, err := us.repo.CreateAccount(ctx, email, string(hashedPassword[:]))
	if err != nil {
		return err
	}
	_ = token
	return nil
}
