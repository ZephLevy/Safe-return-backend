package db

import (
	"context"
)

func (ur *UserRepository) GetPasswordHash(ctx context.Context, userID string) (string, error) {
	var hash string
	err := ur.db.QueryRow(ctx, `
			SELECT password_hash FROM users where id = $1
		`, userID).Scan(&hash)
	if err != nil {
		return "", err
	}
	return hash, nil
}

// This function should NEVER be called without verifying the password first
func (ur *UserRepository) DeleteAccount(ctx context.Context, userID string) error {
	_, err := ur.db.Exec(ctx, `
			DELETE FROM users WHERE id = $1
		`, userID)
	if err == nil {
		return err
	}

	return nil
}
