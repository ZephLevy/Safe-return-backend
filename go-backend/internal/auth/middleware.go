package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/ZephLevy/Safe-return-backend/internal/endpoints/httputils"
)

type contextKey string

const userIDKey contextKey = "userID"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			httputils.WriteJSONError(w, http.StatusUnauthorized, "Missing or invalid token")
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := VerifyToken(tokenStr)
		if err != nil {
			httputils.WriteJSONError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		ctx := context.WithValue(r.Context(), userIDKey, claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
