package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
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

	http.HandleFunc("/setTime", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		time := r.Form.Get("time")

		fmt.Println("Got time: " + time)
		w.Write([]byte("Received time: " + time))
	})
	fmt.Println("Started listening on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
