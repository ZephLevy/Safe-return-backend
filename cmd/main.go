package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
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
	log.Fatal(http.ListenAndServe(":8080", nil))
}
