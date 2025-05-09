package endpoints

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
)

const (
	listenPort = "8080"
)

func OpenEndpoints(conn *pgx.Conn) {
	startLoginListen(conn)
	startCheckListen()
	fmt.Println("Started listening on port: " + listenPort)
	log.Fatal(http.ListenAndServe(":"+listenPort, nil))
}
