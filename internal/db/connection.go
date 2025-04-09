package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func Connect() (*pgx.Conn, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	name := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	for _, val := range []interface{}{host, port, user, name, password} {
		if val == nil {
			return nil, fmt.Errorf("One or more of the environment variables was nil")
		}
	}
	conn, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, name))
	if err != nil {
		return nil, err
	}
	return conn, nil
}
