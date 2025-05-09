package endpoints

import (
	"context"
	"crypto/sha256"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func startLoginListen(conn *pgx.Conn) {
	http.HandleFunc("/signIn", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			fmt.Println("Error parsing form: ", err)
			return
		}

		email := r.Form.Get("email")
		password := r.Form.Get("password")

		if email == "" || password == "" {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		// HACK: Don't use SHA256
		h := sha256.New()
		h.Write([]byte(password))
		passwordHash := h.Sum(nil)

		fmt.Println(email + password + "\n" + string(passwordHash))

		// TODO: Validate email, make sure user doesn't already exist, etc...
		emailIsUnique, err := checkForUniqueEmail(email, conn)
		if err == nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		fmt.Println("Email is unique: ", emailIsUnique)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Login Successful"))
	})
}

func checkForUniqueEmail(email string, conn *pgx.Conn) (bool, error) {
	fmt.Println("Checking if email is unique...")
	row, err := conn.Query(context.Background(), "SELECT * FROM USERS WHERE EMAIL = $1;", email)
	if err != nil {
		fmt.Println("Error checking for server connection: ", err)
		return false, err
	}
	if row == nil {
		return true, nil
	}
	return false, nil
}
