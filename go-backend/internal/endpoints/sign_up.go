package endpoints

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ZephLevy/Safe-return-backend/internal/service"
)

type SignUpRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	EmailOTP  string `json:"emailCode"`
}

type SignUpResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// registerSignUpHandler sets up the signup endpoint
//
// @Summary Sign up a new user
// @Description Creates a new user account using name, email, password, and a one-time email code.
// @Tags Auth
// @Accept application/json
// @Produce application/json
// @Param request body SignUpRequest true "SignUp request payload"
// @Success 200 {object} SignUpResponse "Signup successful, returns access and refresh tokens"
// @Failure 400 {object} ErrorResponse "Incorrect one-time code or bad request"
// @Failure 401 {object} ErrorResponse "Missing required fields"
// @Failure 403 {object} ErrorResponse "Email not verified / OTP expired"
// @Failure 405 {object} ErrorResponse "Method not allowed"
// @Failure 409 {object} ErrorResponse "Email already in use"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /auth/signup [post]
func registerSignUpHandler(userService *service.UserService) {
	http.HandleFunc("/auth/signup", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodPost {
			writeJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}

		var req SignUpRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSONError(w, http.StatusBadRequest, "Bad request")
			return
		}

		accessToken, refreshToken, err := userService.SignUp(
			r.Context(),
			req.FirstName,
			req.LastName,
			req.Email,
			req.Password,
			req.EmailOTP,
		)

		if err != nil {
			var code int
			var message string

			switch err.Error() {
			case "Missing fields":
				code = http.StatusUnauthorized
				message = "Missing required fields"
			case "Email already in use":
				code = http.StatusConflict
				message = "Email already in use"
			case "Incorrect code":
				code = http.StatusBadRequest
				message = "Incorrect one-time code"
			case "Email unknown":
				code = http.StatusForbidden
				message = "Email not verified or OTP expired"
			default:
				code = http.StatusInternalServerError
				message = "Internal Server Error"
				log.Println("Unexpected signup error:", err)
			}

			writeJSONError(w, code, message)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(SignUpResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		})
	})
}
