package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AshokaJS/DhakadFitness/config"
	"github.com/AshokaJS/DhakadFitness/internal/auth"
	"github.com/AshokaJS/DhakadFitness/internal/gym"
	"github.com/AshokaJS/DhakadFitness/internal/user"
	"github.com/AshokaJS/DhakadFitness/pkg/db"
	"github.com/AshokaJS/DhakadFitness/routes"
)

func main() {
	//load environment variables
	config.LoadEnv()
	// Connect to the database
	db.ConnectDB()

	authRepo := auth.NewAuthRepository(db.DB)
	authService := auth.NewAuthService(authRepo)

	userRepo := user.NewUserRepository(db.DB)
	userService := user.NewUserService(userRepo)

	gymRepo := gym.NewGymRepository(db.DB)
	gymService := gym.NewGymService(gymRepo)

	routes.SetupRoutes(authService, userService, gymService)
	// Start the HTTP server on port 8080
	port := ":8080"
	fmt.Println("Server started on", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
