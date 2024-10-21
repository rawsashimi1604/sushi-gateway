package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"log/slog"
)

func ConnectDb() {
	// Connection details
	connStr := "host=localhost port=5432 user=postgres password=mysecretpassword dbname=sushi sslmode=disable"

	// Connect to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		slog.Error(err.Error())
	}

	slog.Info("Successfully connected to the database!")
}
