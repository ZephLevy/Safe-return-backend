package endpoints

import (
	"fmt"
	"log"
	"net/http"
)

func OpenEndpoints() {
	startLoginListen()
	startCheckListen()
	fmt.Println("Started listening on port: " + listenPort)
	log.Fatal(http.ListenAndServe(":"+listenPort, nil))
}
