package db

import "database/sql"

// Postgres wraps the sql.DB connection for PostgreSQL.
type Postgres struct {
	db *sql.DB
}

// NewPostgres creates a Postgres helper around an existing sql.DB instance.
func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{db: db}
}

// DB exposes the underlying sql.DB instance.
func (p *Postgres) DB() *sql.DB {
	if p == nil {
		return nil
	}

	return p.db
}
