package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/AshokaJS/DhakadFitness/config"
)

// Global variable to store the database connection
var DB *sql.DB

func ConnectDB() {

	// Fetching the database URL from environment variables
	dbURL := config.GetEnv("POSTGRESQL_URL")
	if dbURL == "" {
		log.Fatal("POSTGRESQL_URL not found in environment variables")
	}

	// Connecting to PostgreSQL
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	// Checking the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Database connection failed: %s", err)
	}

	fmt.Println("Connected to Database")
	DB = db
}
