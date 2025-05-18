package endpoints

import (
	"net/http"

	"github.com/ZephLevy/Safe-return-backend/internal/service"
)

func startEmailVerificationListen(userService *service.UserService) {
	http.HandleFunc("/auth/verify-email", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		email := r.Form.Get("email")
		if email == "" {
			http.Error(w, "Missing required fields", http.StatusUnauthorized)
			return
		}
		isValid, msg := userService.VerifyEmail(r.Context(), email)
		if !isValid {
			var statusCode int
			switch msg {
			case "Email already in use.":
				statusCode = http.StatusConflict
			case "Invalid email.":
				statusCode = http.StatusBadRequest
			default:
				statusCode = http.StatusInternalServerError
			}
			http.Error(w, msg, statusCode)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Email valid"))
	})
}
