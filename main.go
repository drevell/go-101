package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abcxyz/pkg/healthcheck"
)

func handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "nobody"
	}
	fmt.Fprintf(w, "👋, %s!\n", name)
}

func main() {
	http.HandleFunc("/hello", handler)
	http.Handle("/health", healthcheck.HandleHTTPHealthCheck())

	log.Print("Starting HTTP server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
