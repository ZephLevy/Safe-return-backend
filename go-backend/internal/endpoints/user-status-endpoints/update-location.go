package userstatusendpoints

import (
	"encoding/json"
	"net/http"

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
	http.HandleFunc("/user-status/update-location", func(w http.ResponseWriter, r *http.Request) {
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

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("All ok!"))
	})
}
