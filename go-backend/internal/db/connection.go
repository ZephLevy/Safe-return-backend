package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

func Connect() (*pgx.Conn, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	name := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	for i, val := range []any{host, port, user, name, password} {
		if val == nil || val == "" {
			return nil, fmt.Errorf("Null env variable: %d", i)
		}
	}
	// Docker compose starts both the db and the server at the same time,
	// but the server takes time to start up leading to a failed connection
	var conn *pgx.Conn
	var err error
	for range 10 {
		conn, err = pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, name))
		if err == nil {
			fmt.Println("Connection successful")
			break
		} else {
			time.Sleep(time.Second)
		}
	}

	if err != nil {
		return nil, err
	}
	return conn, nil
}
