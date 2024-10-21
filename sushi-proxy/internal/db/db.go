package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log/slog"
)

func ConnectDb() (*sql.DB, error) {

	// TODO: externalize configurations
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", "5432", "postgres", "mysecretpassword", "sushi",
	)

	// Connect to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	slog.Info("Successfully connected to the database!")
	return db, nil
}
