package routes

import (
	"Arbitrax/cmd/api/handlers"
	"Arbitrax/pkg/middleware"
	"Arbitrax/pkg/output"

	"github.com/gorilla/mux"
)

func AgentRoutes(r *mux.Router, h *handlers.AuthHandler, auth middleware.Middleware) {
	output.MakeRoute(r, "/create", h.Register).Methods("POST", "OPTIONS")
}
