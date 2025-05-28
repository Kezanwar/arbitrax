package routes

import (
	"Arbitrax/cmd/api/handlers"
	"Arbitrax/pkg/middleware"
	"Arbitrax/pkg/output"

	"github.com/gorilla/mux"
)

func OptionsRoutes(r *mux.Router, h *handlers.OptionsHandler, authCached middleware.Middleware) {
	output.MakeRoute(r, "/exchanges", h.GetExchangesOptions, authCached).Methods("GET", "OPTIONS")
	output.MakeRoute(r, "/strategies", h.GetStrategiesOptions, authCached).Methods("GET", "OPTIONS")
}
