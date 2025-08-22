package authendpoints

import (
	"encoding/json"
	"net/http"

	"github.com/ZephLevy/Safe-return-backend/internal/endpoints/httputils"
	"github.com/ZephLevy/Safe-return-backend/internal/service"
)

type emailVerifyRequest struct {
	Email string `json:"email"`
}

type emailVerifyResponse struct {
	Response string `json:"response"`
}

// registerEmailAuthHandler sets up an endpoint for verifying up an email
//
// @Summary Verify email for signup
// @Description Checks if an email is valid and not already in use.
// @Tags Auth
// @Accept application/json
// @Produce application/json
// @Param request body emailVerifyRequest true "Email in JSON"
// @Success 200 {object} emailVerifyResponse "Email valid"
// @Failure 400 {object} httputils.ErrorResponse "Invalid email"
// @Failure 401 {object} httputils.ErrorResponse "Missing required fields"
// @Failure 405 {object} httputils.ErrorResponse "Method not allowed"
// @Failure 409 {object} httputils.ErrorResponse "Email already in use"
// @Failure 500 {object} httputils.ErrorResponse "Internal server error"
// @Router /auth/verify-email [post]
func registerEmailAuthHandler(userService *service.UserService) {
	http.HandleFunc("/auth/verify-email", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req emailVerifyRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			httputils.WriteJSONError(w, http.StatusBadRequest, "Bad request")
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
			httputils.WriteJSONError(w, statusCode, msg)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(emailVerifyResponse{
			Response: "Email verified.",
		})
	})
}
