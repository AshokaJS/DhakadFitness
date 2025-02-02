package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AshokaJS/DhakadFitness/config"
	"github.com/AshokaJS/DhakadFitness/pkg/db"
)

func main() {
	//load environment variables
	config.LoadEnv()
	// Connect to the database
	db.ConnectDB()

	// Define a simple health check route
	http.HandleFunc("/health", dbhandler)

	// Start the HTTP server on port 8080
	port := ":8080"
	fmt.Println("🚀 Server started on", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}

func dbhandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is running!"))
}
