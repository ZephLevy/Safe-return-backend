package endpoints

import (
	"fmt"
	"net/http"

	"github.com/ZephLevy/Safe-return-backend/internal/service"
)

// registerSignUpHandler sets up the signup endpoint
//
// @Summary Sign up a new user
// @Description Creates a new user account using name, email, password, and a one-time email code.
// @Tags Auth
// @Accept application/x-www-form-urlencoded
// @Produce plain
// @Param firstName formData string true "First Name"
// @Param lastName formData string false "Last Name"
// @Param email formData string true "Email"
// @Param password formData string true "Password"
// @Param emailCode formData string true "Email OTP Code"
// @Success 200 {string} string "Signup successful"
// @Failure 400 {string} string "Incorrect one time code / Bad request"
// @Failure 401 {string} string "Missing required fields
// @Failure 403 {string} string "Email not verified / OTP expired"
// @Failure 405 {string} string "Method not allowed"
// @Failure 409 {string} string "Email already in use"
// @Failure 500 {string} string "Internal Server Error"
// @Router /auth/signup [post]
func registerSignUpHandler(userService *service.UserService) {
	http.HandleFunc("/auth/signup", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		firstName := r.FormValue("firstName")
		lastName := r.FormValue("lastName")
		email := r.FormValue("email")
		password := r.FormValue("password")
		emailOTP := r.FormValue("emailCode")

		err = userService.SignUp(r.Context(), firstName, lastName, email, password, emailOTP)
		if err != nil {
			var errorCode int
			var errorMessage string
			switch err.Error() {
			case "Missing fields":
				errorCode = http.StatusUnauthorized
				errorMessage = "Missing required fields"
			case "Email already in use":
				errorCode = http.StatusConflict
				errorMessage = "Email already in use"
			case "Incorrect code":
				errorCode = http.StatusBadRequest
				errorMessage = "Incorrect one time code"
			case "Email unknown":
				errorCode = http.StatusForbidden
				errorMessage = "Email not verified or the password has expired"
			default:
				errorCode = http.StatusInternalServerError
				errorMessage = "Internal Server Error"
				fmt.Println("Error signing up: " + err.Error())
			}
			http.Error(w, errorMessage, errorCode)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Signup successful"))
	})
}
