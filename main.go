package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Infof("uri: %s", r.RequestURI)
	fmt.Fprint(w, "Hello Go")
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
