package db

import (
	"database/sql"
)

type GatewayRepository struct {
	db *sql.DB
}

func NewGatewayRepository(db *sql.DB) *GatewayRepository {
	return &GatewayRepository{db: db}
}

// TODO: add methods
