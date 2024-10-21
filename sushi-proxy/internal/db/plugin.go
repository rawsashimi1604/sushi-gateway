package db

import (
	"database/sql"
)

type PluginRepository struct {
	db *sql.DB
}

func NewPluginRepository(db *sql.DB) *PluginRepository {
	return &PluginRepository{db: db}
}

// TODO: add methods
// scope -> global, service, route
func GetPlugins(scope string) {

}

func AddPlugin() {

}
