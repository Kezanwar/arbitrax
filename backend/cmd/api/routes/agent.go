package routes

import (
	"Arbitrax/cmd/api/handlers"
	"Arbitrax/pkg/middleware"
	"Arbitrax/pkg/output"

	"github.com/gorilla/mux"
)

func AgentRoutes(r *mux.Router, h *handlers.AgentHandler, authCached middleware.Middleware) {
	output.MakeRoute(r, "/", h.GetAgents, authCached).Methods("GET", "OPTIONS")
	output.MakeRoute(r, "/create", h.CreateAgent, authCached).Methods("POST", "OPTIONS")

}
