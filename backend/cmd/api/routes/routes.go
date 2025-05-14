package routes

import (
	"Arbitrax/cmd/api/handlers"
	"Arbitrax/pkg/middleware"
	"Arbitrax/pkg/output"

	"github.com/gorilla/mux"
)

func Register(r *mux.Router, authHandler *handlers.AuthHandler, authMiddleware middleware.Middleware) {
	output.MakeSubRouter(r, "/auth", func(sr *mux.Router) {
		AuthRoutes(sr, authHandler, authMiddleware)
	})
}
