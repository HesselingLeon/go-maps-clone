package repository

import "database/sql"

type TileRepository interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

type tileRepository struct {
	db *sql.DB
}

func NewTileRepository(db *sql.DB) TileRepository {
	return &tileRepository{db: db}
}

func (r *tileRepository) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return r.db.Query(query, args...)
}
