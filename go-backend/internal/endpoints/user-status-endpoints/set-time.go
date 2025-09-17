package userstatusendpoints

import (
	"encoding/json"
	"net/http"

	"github.com/ZephLevy/Safe-return-backend/internal/auth"
	"github.com/ZephLevy/Safe-return-backend/internal/endpoints/httputils"
)

type setTimeRequest struct {
	Time string `json:"time"`
}

type setTimeResponse struct {
	Message string `json:"message"`
}

// registerSetTimeHandler sets up listening to see when a user starts a trip
//
// @Summary Start a user's trip
// @Description Schedules a function to check on the user to see if they are ok
// @Tags user-status
// @Accept application/json
// @Produce application/json
// @Param request body setTimeRequest true "Time encoded in json"
// @Success 200 {object} setTimeResponse
// @Failure 400 {object} httputils.ErrorResponse "Bad Request"
// @Failure 401 {object} httputils.ErrorResponse "Unauthorized - see error message for more detail"
// @Failure 405 {object} httputils.ErrorResponse "Method not allowed"
// @Security BearerAuth
// @Router /user-status/set-time [POST]
func registerSetTimeHandler() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value(auth.UserIDKey)
		_ = userID
		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodPost {
			httputils.WriteJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}

		var req setTimeRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			httputils.WriteJSONError(w, http.StatusBadRequest, "Bad Request")
			return
		}

		// TODO: Schedule an event to check if user is OK
		// In the user-service probably

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(setTimeResponse{
			Message: "Time set sucessfully",
		})
	})

	http.Handle("/user-status/set-time", auth.AuthMiddleware(handler))
}
