package userstatusendpoints

import (
	"encoding/json"
	"fmt"
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

type locationResponse struct {
	Response string `json:"response"`
}

// registerUserLocationHandler sets up listening for location
//
// @Summary Updates the last known location of a user
// @Description Inserts a locationRequest into the db describing location, accuracy, speed, and time
// @Tags user-status
// @Accept application/json
// @Produce application/json
// @Param request body locationRequest true "Location updater request payload"
// @Success 201 {object} locationResponse "Location updated successfully"
// @Failure 405 {object} httputils.ErrorResponse "Method not allowed"
// @Failure 401 {object} httputils.ErrorResponse "Unauthorized - see error message for more detail"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request"
// @Security BearerAuth
// @Router /user-status/update-location [post]
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

		fmt.Println(userID)

		// TODO: Save the location in the db
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(locationResponse{
			Response: "Location updated successfully",
		})
	})

	http.Handle("/user-status/update-location", auth.AuthMiddleware(locationHandler))
}
