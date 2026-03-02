package main

import (
	"database/sql"
	"jwt-auth/internal/api"
	"jwt-auth/internal/config"
	"jwt-auth/internal/db"
	"log"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	database, err := db.NewMySQLStorage(cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	if err != nil {
		log.Fatal(err)
	}

	// Check database connection
	initStorage(database)

	// Initialize API server
	apiServer := api.NewAPIServer(database, cfg)
	if err := apiServer.Run(); err != nil {
		log.Fatal(err)
	}
}

// initStorage checks the database connection
func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database successfully connected")
}
