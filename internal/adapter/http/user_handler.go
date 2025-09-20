package http

import (
	stdhttp "net/http"

	"subscription_system_app/internal/domain"
	"subscription_system_app/internal/usecase"
)

// UserHandler exposes HTTP endpoints for user operations.
type UserHandler struct {
	creator *usecase.UserCreator
}

// NewUserHandler constructs a handler with the provided use case.
func NewUserHandler(creator *usecase.UserCreator) *UserHandler {
	return &UserHandler{creator: creator}
}

// HandleCreate registers a user using the POST /users endpoint.
func (h *UserHandler) HandleCreate(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	if r.Method != stdhttp.MethodPost {
		w.WriteHeader(stdhttp.StatusMethodNotAllowed)
		return
	}

	if h.creator != nil {
		// Placeholder: integrate request parsing when implementing HTTP body handling.
		_ = h.creator.Create(domain.User{})
	}

	w.WriteHeader(stdhttp.StatusCreated)
}
