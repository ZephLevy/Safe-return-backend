package service

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func (us *UserService) SignIn(ctx context.Context,
	firstName string,
	lastName string,
	email string,
	password string,
) error {
	if firstName == "" || email == "" || password == "" {
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

	token, err := us.repo.CreateAccount(ctx, firstName, lastName, email, string(hashedPassword[:]))
	if err != nil {
		return err
	}
	_ = token
	return nil
}
