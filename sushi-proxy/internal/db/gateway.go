package db

import (
	"database/sql"
	"fmt"
)

type GatewayRepository struct {
	db *sql.DB
}

func NewGatewayRepository(db *sql.DB) *GatewayRepository {
	return &GatewayRepository{db: db}
}

// GetGatewayInfo retrieves the gateway information from the "gateway" table.
func (gatewayRepo *GatewayRepository) GetGatewayInfo() (string, error) {
	var gatewayName string

	query := `SELECT name FROM gateway LIMIT 1` // Assuming you are fetching the first or only row
	err := gatewayRepo.db.QueryRow(query).Scan(&gatewayName)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no gateway information found")
		}
		return "", fmt.Errorf("failed to fetch gateway information: %w", err)
	}

	return gatewayName, nil
}
