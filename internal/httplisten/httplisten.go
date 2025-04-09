package httplisten

import (
	"fmt"
	"log"
	"net/http"
)

const listenPort string = "8080"

func StartTimeListen() {
	http.HandleFunc("/setTime", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		time := r.Form.Get("time")

		fmt.Println("Got time: " + time)
		w.Write([]byte("Received time: " + time))
	})
	fmt.Println("Started listening on port: " + listenPort)
	log.Fatal(http.ListenAndServe(":"+listenPort, nil))
}
