package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	name := "(no name provided)"
	if n := r.URL.Query().Get("name"); n != "" {
		name = n
	}
	fmt.Fprintf(w, "Hello, %s!\n", name)
}

func main() {
	http.HandleFunc("/hello", handler)

	log.Print("Starting HTTP server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
