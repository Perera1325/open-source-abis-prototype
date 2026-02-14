package main

import (
	"log"
	"net/http"

	"github.com/Perera1325/open-source-abis-prototype/internal/handlers"
)

func main() {
	http.HandleFunc("/health", handlers.Health)
	http.HandleFunc("/enroll", handlers.Enroll)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
