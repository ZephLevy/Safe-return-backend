package service

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func (us *UserService) CheckPassword(ctx context.Context, userID string, password string) (bool, error) {
	passwordHash, err := us.repo.GetPasswordHash(ctx, userID)
	if err != nil {
		return false, err
	}

	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)) == nil, nil
}

func (us *UserService) DeleteAccount(ctx context.Context, userID string, password string) error {
	passwordCorrect, err := us.CheckPassword(ctx, userID, password)
	if err != nil {
		return err
	}
	if !passwordCorrect {
		return fmt.Errorf("Incorrect Password")
	}

	return us.repo.DeleteAccount(ctx, userID)
}
