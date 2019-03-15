package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Go")
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
