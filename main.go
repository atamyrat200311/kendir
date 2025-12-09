package main

import (
	"log"
	"net/http"

	"kendir-mini/controller"
	"kendir-mini/db"
)

func main() {
	// Connect database
	db.Connect()

	// Routes
	http.HandleFunc("/api/user/create", controller.UserCreate)
	http.HandleFunc("/api/user/get", controller.UserGet)

	// Start server
	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
