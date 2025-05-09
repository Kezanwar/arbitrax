package routes

import (
	"Arbitrax/handlers"
	"Arbitrax/middleware"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router, h *handlers.AuthHandler, auth middleware.Middleware) {
	MakeRoute(r, "/register", h.Register).Methods("POST", "OPTIONS")
	MakeRoute(r, "/sign-in", h.SignIn).Methods("POST", "OPTIONS")
	MakeRoute(r, "/initialize", h.Initialize, auth).Methods("GET", "OPTIONS")
}
