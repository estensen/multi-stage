package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	helloHandler := http.HandlerFunc(hello)
	http.Handle("/", logging(helloHandler))
	log.Info("Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func( w http.ResponseWriter, r *http.Request) {
		log.Infof("uri: %s", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func hello(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello Go")
}
