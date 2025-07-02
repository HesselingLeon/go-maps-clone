package database

import (
	"database/sql"
	"fmt"

	"github.com/hesselingleon/go-map-clone/internal/config"
	_ "github.com/lib/pq"
)

func NewDatabase(config *config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("dbname=%s user=%s password=%s host=%s sslmode=%s",
		config.DatabaseName,
		config.DatabaseUser,
		config.DatabasePassword,
		config.DatabaseURL,
		config.DatabaseSSLMODE)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
