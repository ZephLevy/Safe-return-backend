package userstatusendpoints

import (
	"encoding/json"
	"net/http"

	"github.com/ZephLevy/Safe-return-backend/internal/auth"
	"github.com/ZephLevy/Safe-return-backend/internal/endpoints/httputils"
)

type locationRequest struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Accuracy  string `json:"accuracy"`
	Speed     string `json:"speed"`
	Time      string `json:"time"`
}

// registerUserLocationHandler sets up listening for location
func registerUserLocationHandler() {
	locationHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		userID := r.Context().Value(auth.UserIDKey).(string)
		w.Header().Set("Content-Type", "application/json")

		// POST and not PUT here because we're storing a sequence of locations
		if r.Method != http.MethodPost {
			httputils.WriteJSONError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
			return
		}

		var req locationRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			httputils.WriteJSONError(w, http.StatusBadRequest, "Bad Request")
			return
		}

		// TODO: Actually respond in JSON
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(userID))
	})

	http.Handle("/user-status/update-location", auth.AuthMiddleware(locationHandler))
}
