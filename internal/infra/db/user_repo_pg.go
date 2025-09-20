package db

import "subscription_system_app/internal/domain"

// UserRepository persists users in PostgreSQL.
type UserRepository struct {
	pg *Postgres
}

// NewUserRepository wires the repository with a Postgres instance.
func NewUserRepository(pg *Postgres) *UserRepository {
	return &UserRepository{pg: pg}
}

// Create inserts a new user record. Implementation pending actual SQL.
func (r *UserRepository) Create(user domain.User) error {
	// TODO: implement INSERT statement once schema is defined.
	_ = user
	return nil
}
