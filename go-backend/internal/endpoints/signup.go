package endpoints

import (
	"fmt"
	"net/http"

	"github.com/ZephLevy/Safe-return-backend/internal/service"
)

func startSignUpListen(userService *service.UserService) {
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

		err = userService.SignIn(r.Context(), firstName, lastName, email, password)
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
			default:
				errorCode = http.StatusInternalServerError
				errorMessage = "Internal Server Error"
				fmt.Println("Error signing in: " + err.Error())
			}
			http.Error(w, errorMessage, errorCode)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Login successful"))
	})
}
