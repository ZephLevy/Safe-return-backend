package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/ZephLevy/Safe-return-backend/internal/service"
)

type EmailVerifyRequest struct {
	Email string `json:"email"`
}

type EmailVerifyResponse struct {
	Response string `json:"response"`
}

// registerEmailAuthHandler sets up an endpoint for verifying up an email
//
// @Summary Verify email for signup
// @Description Checks if an email is valid and not already in use. Responds with plain text.
// @Tags Auth
// @Accept application/json
// @Produce application/json
// @Param request body EmailVerifyRequest true "Email in JSON"
// @Success 200 {object} EmailVerifyResponse "Email valid"
// @Failure 400 {object} ErrorResponse "Invalid email"
// @Failure 401 {object} ErrorResponse "Missing required fields"
// @Failure 405 {object} ErrorResponse "Method not allowed"
// @Failure 409 {object} ErrorResponse "Email already in use"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /auth/verify-email [post]
func registerEmailAuthHandler(userService *service.UserService) {
	http.HandleFunc("/auth/verify-email", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req EmailVerifyRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSONError(w, http.StatusBadRequest, "Bad request")
			return
		}

		email := req.Email
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
			writeJSONError(w, statusCode, msg)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(EmailVerifyResponse{
			Response: "Email verified.",
		})
	})
}
