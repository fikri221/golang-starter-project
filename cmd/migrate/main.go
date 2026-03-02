package main

import (
	"jwt-auth/internal/config"
	"jwt-auth/internal/db"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file" // the _ is to call the init function without using it
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	database, err := db.NewMySQLStorage(cfg.DBUser, cfg.DBPasswd, cfg.DBHost, cfg.DBPort, cfg.DBName)
	if err != nil {
		log.Fatal(err)
	}

	driver, err := mysql.WithInstance(database, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://cmd/migrate/migrations", "mysql", driver)
	if err != nil {
		log.Fatal(err)
	}

	cmd := ""
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}
	switch cmd {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		log.Println("Migrations Up successfully executed")
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		log.Println("Migrations Down successfully executed")
	default:
		log.Fatal("Please specify command: 'up' or 'down'")
	}
}
