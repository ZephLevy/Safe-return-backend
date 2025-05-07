package endpoints

import (
	"fmt"
	"log"
	"net/http"
)

const (
	listenPort = "8080"
)

func OpenEndpoints() {
	startLoginListen()
	startCheckListen()
	fmt.Println("Started listening on port: " + listenPort)
	log.Fatal(http.ListenAndServe(":"+listenPort, nil))
}
