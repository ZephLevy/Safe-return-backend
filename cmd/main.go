package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	timelisten "github.com/ZephLevy/Safe-return-backend/internal"
)

const port string = "8080"
const postgresPort string = "5432"

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

	timelisten.StartListen()
}
