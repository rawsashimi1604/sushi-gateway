package db

import (
	"context"
	"github.com/rawsashimi1604/sushi-gateway/admin-api/internal/config"
	"github.com/rawsashimi1604/sushi-gateway/admin-api/internal/error"
	"log/slog"
)
import "github.com/jackc/pgx/v5/pgxpool"

func CreatePostgresConnection() (*pgxpool.Pool, *error.GenericError) {
	slog.Info("Connecting to postgres...")
	conn, err := pgxpool.New(context.Background(), config.GlobalAppConfig.PostgresUrl)
	if err != nil {
		return nil, error.NewGenericError("POSTGRES_CONNECT_ERROR", "Unable to connect to postgres")
	}

	if err = conn.Ping(context.Background()); err != nil {
		return nil, error.NewGenericError("POSTGRES_PING_ERROR", "Unable to ping postgres")
	}
	slog.Info("Postgres connected successfully.")
	return conn, nil
}
