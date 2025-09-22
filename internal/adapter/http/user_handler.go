package http

import (
	stdhttp "net/http"

	"github.com/gin-gonic/gin"

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
func (h *UserHandler) HandleCreate(c *gin.Context) {
	if h.creator != nil {
		// Placeholder: integrate request parsing when implementing HTTP body handling.
		_ = h.creator.Create(domain.User{})
	}

	c.Status(stdhttp.StatusCreated)
}
