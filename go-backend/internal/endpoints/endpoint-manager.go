package endpoints

import (
	"fmt"
	"log"
	"net/http"

	authendpoints "github.com/ZephLevy/Safe-return-backend/internal/endpoints/auth-endpoints"
	userstatusendpoints "github.com/ZephLevy/Safe-return-backend/internal/endpoints/user-status-endpoints"
	"github.com/ZephLevy/Safe-return-backend/internal/service"

	_ "github.com/ZephLevy/Safe-return-backend/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

const (
	listenPort = "8080"
)

func OpenEndpoints(userService *service.UserService) {
	authendpoints.OpenAuthEndpoints(userService)
	userstatusendpoints.OpenUserStatusEndpoints()
	fmt.Println("Started listening on port: " + listenPort)

	http.Handle("/docs/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/docs/doc.json"),
	))

	log.Fatal(http.ListenAndServe(":"+listenPort, nil))
}
