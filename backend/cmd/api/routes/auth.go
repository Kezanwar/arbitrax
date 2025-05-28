package routes

import (
	"Arbitrax/cmd/api/handlers"
	"Arbitrax/pkg/middleware"
	"Arbitrax/pkg/output"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router, h *handlers.AuthHandler, authCached middleware.Middleware) {
	output.MakeRoute(r, "/register", h.Register).Methods("POST", "OPTIONS")
	output.MakeRoute(r, "/sign-in", h.SignIn).Methods("POST", "OPTIONS")
	output.MakeRoute(r, "/initialize", h.Initialize, authCached).Methods("GET", "OPTIONS")
}
