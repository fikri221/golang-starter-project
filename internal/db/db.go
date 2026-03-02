package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func NewMySQLStorage(user, password, host, port, dbName string) (*sql.DB, error) {
	cfg := mysql.Config{
		User:                 user,
		Passwd:               password,
		Addr:                 host + ":" + port,
		DBName:               dbName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
