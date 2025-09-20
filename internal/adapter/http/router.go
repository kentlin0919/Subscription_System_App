package http

import stdhttp "net/http"

// NewRouter wires HTTP routes and handlers.
func NewRouter(userHandler *UserHandler) *stdhttp.ServeMux {
	mux := stdhttp.NewServeMux()

	if userHandler != nil {
		mux.HandleFunc("/users", userHandler.HandleCreate)
	}

	return mux
}
