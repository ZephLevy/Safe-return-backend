package endpoints

import "net/http"

func startEmailVerificationListen() {
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
		_ = email
		// TODO: Verify email

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Email valid"))
	})
}
