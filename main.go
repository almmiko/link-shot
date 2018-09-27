package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := "8080"

	if v := os.Getenv("PORT"); v != "" {
		port = v
	}

	r := NewRouter()

	log.Printf("Server is running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
