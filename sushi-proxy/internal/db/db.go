package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log/slog"
)

func ConnectDb() {

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", "5432", "postgres", "mysecretpassword", "sushi",
	)

	// Connect to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		slog.Error(err.Error())
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		slog.Error(err.Error())
	}

	slog.Info("Successfully connected to the database!")
}
