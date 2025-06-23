package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ZephLevy/Safe-return-backend/internal/db"
	"github.com/ZephLevy/Safe-return-backend/internal/endpoints"
	"github.com/ZephLevy/Safe-return-backend/internal/service"
)

func main() {

	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			char, _, err := reader.ReadRune()
			if err != nil {
				continue
			}
			if strings.ToLower(string(char)) == "c" {
				fmt.Print("\033[H\033[2J")
			}
		}
	}()
	conn, err := db.Connect()
	if err != nil {
		log.Fatalf("Error connecting to db: %v", err)
	}
	defer conn.Close(context.Background())

	UserRepository := db.NewUserRepo(conn)
	UserService := service.NewUserService(UserRepository)
	endpoints.OpenEndpoints(UserService)
}
