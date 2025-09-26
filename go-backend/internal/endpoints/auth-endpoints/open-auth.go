package authendpoints

import "github.com/ZephLevy/Safe-return-backend/internal/service"

func OpenAuthEndpoints(userService *service.UserService) {
	registerEmailAuthHandler(userService)
	registerSignUpHandler(userService)
	registerAccountDeletionHandler(userService)
}
