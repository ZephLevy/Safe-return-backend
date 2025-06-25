package service

import (
	"context"
	"fmt"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

func (us *UserService) SignUp(ctx context.Context,
	firstName string,
	lastName string,
	email string,
	password string,
	emailOTP string,
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

	correctCode, err := us.repo.VerifyEmailOTP(ctx, email, emailOTP)
	if err != nil {
		if err.Error() == "redis: nil" {
			return fmt.Errorf("Email unknown")
		}
		return err
	}

	if !correctCode {
		return fmt.Errorf("Incorrect code")
	}

	token, err := us.repo.CreateAccount(ctx, firstName, lastName, email, string(hashedPassword[:]))
	if err != nil {
		return err
	}
	_ = token
	return nil
}

func (us *UserService) VerifyEmail(ctx context.Context, email string) (bool, string) {
	isUnique, err := us.repo.IsEmailUnique(ctx, email)
	if err != nil {
		fmt.Println("Error checking for unique email: ", err)
		return false, "Internal server error."
	}
	if !isUnique {
		return false, "Email already in use."
	}

	_, err = mail.ParseAddress(email)
	if err != nil {
		return false, "Invalid email."
	}
	// TODO: Check for email ownership with a 6 digit code
	// (stored in redis)
	// I'll just make it 111111 for now
	// until I set up sending via SMTP
	us.repo.SetEmailOTP(ctx, email, "111111")
	return true, ""
}
