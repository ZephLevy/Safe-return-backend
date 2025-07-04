package endpoints

import (
	"fmt"
	"net/http"
)

func registerCheckHandler() {
	http.HandleFunc("/user-status/set-time", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		time := r.Form.Get("time")

		if time == "" {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		fmt.Println("Got time: " + time)
		w.Write([]byte("Received time: " + time))
	})
}
