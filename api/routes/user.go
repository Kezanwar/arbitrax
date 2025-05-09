package routes

import (
	"Arbitrax/handlers"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	MakeRoute(r, "/", handlers.GetUsers).Methods("GET", "OPTIONS")
}
