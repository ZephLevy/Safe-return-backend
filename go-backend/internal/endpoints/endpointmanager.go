package endpoints

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ZephLevy/Safe-return-backend/internal/service"
)

const (
	listenPort = "8080"
)

func OpenEndpoints(userService *service.UserService) {
	startSignUpListen(userService)
	startEmailVerificationListen()
	startCheckListen()
	fmt.Println("Started listening on port: " + listenPort)
	log.Fatal(http.ListenAndServe(":"+listenPort, nil))
}
