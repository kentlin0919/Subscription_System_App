package usecase

import "subscription_system_app/internal/domain"

// UserRepository defines persistence operations for users.
type UserRepository interface {
	Create(user domain.User) error
}

// UserCreator handles user onboarding logic.
type UserCreator struct {
	repo UserRepository
}

// NewUserCreator wires the user creation use case with a repository.
func NewUserCreator(repo UserRepository) *UserCreator {
	return &UserCreator{repo: repo}
}

// Create registers a new user in the underlying repository.
func (uc *UserCreator) Create(user domain.User) error {
	if uc.repo == nil {
		return nil
	}

	return uc.repo.Create(user)
}
