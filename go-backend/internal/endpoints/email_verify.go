package endpoints

import (
	"net/http"

	"github.com/ZephLevy/Safe-return-backend/internal/service"
)

// registerEmailAuthHandler sets up an endpoint for verifying up an email
//
// @Summary Verify email for signup
// @Description Checks if an email is valid and not already in use. Responds with plain text.
// @Tags Auth
// @Accept application/x-www-form-urlencoded
// @Produce text/plain
// @Param email formData string true "Email"
// @Success 200 {string} string "Email valid"
// @Failure 400 {string} string "Invalid email"
// @Failure 401 {string} string "Missing required fields"
// @Failure 405 {string} string "Method not allowed"
// @Failure 409 {string} string "Email already in use"
// @Failure 500 {string} string "Internal server error"
// @Router /auth/verify-email [post]
func registerEmailAuthHandler(userService *service.UserService) {
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
