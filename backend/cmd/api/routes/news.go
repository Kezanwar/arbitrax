package routes

import (
	"Arbitrax/cmd/api/handlers"
	"Arbitrax/pkg/middleware"
	"Arbitrax/pkg/output"

	"github.com/gorilla/mux"
)

func NewsRoutes(r *mux.Router, h *handlers.NewsHandler, authCached middleware.Middleware) {
	output.MakeRoute(r, "/", h.GetNews, authCached).Methods("GET", "OPTIONS")
}
