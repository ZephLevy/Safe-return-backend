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
	"github.com/ZephLevy/Safe-return-backend/internal/envloader"
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
	err := envloader.Load(".env")
	if err != nil {
		log.Fatalln("Had error reading environment variables:", err)
	}
	conn, err := db.Connect()
	if err != nil {
		// Right now, while the app is in development, I don't require a db connection most of the time
		// This should be replaced by log.fatal later though
		fmt.Println("Error conecting to database", err)
	}
	defer conn.Close(context.Background())

	endpoints.OpenEndpoints()
}
