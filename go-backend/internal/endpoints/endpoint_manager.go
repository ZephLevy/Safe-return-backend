package endpoints

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ZephLevy/Safe-return-backend/internal/service"

	_ "github.com/ZephLevy/Safe-return-backend/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

const (
	listenPort = "8080"
)

func OpenEndpoints(userService *service.UserService) {
	registerSignUpHandler(userService)
	registerEmailAuthHandler(userService)
	registerCheckHandler()
	fmt.Println("Started listening on port: " + listenPort)

	http.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	log.Fatal(http.ListenAndServe(":"+listenPort, nil))
}
