package endpoints

import (
	"net/http"

	"github.com/ZephLevy/Safe-return-backend/internal/service"
)

func startLoginListen(userService *service.UserService) {
	http.HandleFunc("/signIn", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		email := r.FormValue("email")
		password := r.FormValue("password")

		err = userService.SignIn(r.Context(), email, password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Login successful"))
	})
}
