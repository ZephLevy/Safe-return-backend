package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ZephLevy/Safe-return-backend/internal/httplisten"
)

const port string = "8080"

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
	httplisten.StartTimeListen()
}
