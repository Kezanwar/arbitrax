package routes

import (
	"Arbitrax/cmd/api/handlers"
	"Arbitrax/pkg/output"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	output.MakeRoute(r, "/", handlers.GetUsers).Methods("GET", "OPTIONS")
}
