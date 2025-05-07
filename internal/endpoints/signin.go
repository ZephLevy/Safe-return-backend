package endpoints

import (
	"crypto/sha256"
	"fmt"
	"net/http"
)

func startLoginListen() {
	http.HandleFunc("/signIn", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		username := r.Form.Get("email")
		password := r.Form.Get("password")

		if username == "" || password == "" {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		// HACK: Don't use SHA256
		h := sha256.New()
		h.Write([]byte(password))
		passwordHash := h.Sum(nil)

		fmt.Println(username + password + "\n" + string(passwordHash))

		// TODO: Validate email, make sure user doesn't already exist, etc...
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Login Successful"))
	})
}
