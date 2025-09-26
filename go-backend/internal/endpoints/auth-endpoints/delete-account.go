package authendpoints

import (
	"encoding/json"
	"net/http"

	"github.com/ZephLevy/Safe-return-backend/internal/auth"
	"github.com/ZephLevy/Safe-return-backend/internal/endpoints/httputils"
	"github.com/ZephLevy/Safe-return-backend/internal/service"
)

type AccountDeletionRequest struct {
	Password string `json:"password"`
}

type AccountDeletionResponse struct {
	Message string `json:"message"`
}

// registerAccountDeletionHandler sets up the handler for account deletion
//
// @Summary Delete an account
// @Description Use the Password to delete an account from the database
// @Tags Auth
// @Accept application/json
// @Produce application/json
// @Param request body AccountDeletionRequest true "Deletion request payload"
// @Success 200 {object} AccountDeletionResponse "Deletion successful"
// @Failure 400 {object} httputils.ErrorResponse "Bad Request"
// @Failure 405 {object} httputils.ErrorResponse "Method Not Allowed"
// @Security BearerAuth
// @Router /auth/delete-account [post]
func registerAccountDeletionHandler(us *service.UserService) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value(auth.UserIDKey).(string)

		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodPost {
			httputils.WriteJSONError(w, http.StatusBadRequest, "Method Not Allowed")
			return
		}

		var req AccountDeletionRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			httputils.WriteJSONError(w, http.StatusBadRequest, "Bad Request")
			return
		}

		us.DeleteAccount(r.Context(), userID, req.Password)

	})

	http.Handle("/auth/delete-account", auth.AuthMiddleware(handler))
}
