package routes

import (
	"Arbitrax/handlers"
	"Arbitrax/middleware"

	"github.com/gorilla/mux"
)

func Register(r *mux.Router, authHandler *handlers.AuthHandler, authMiddleware middleware.Middleware) {
	MakeSubRouter(r, "/auth", func(sr *mux.Router) {
		AuthRoutes(sr, authHandler, authMiddleware)
	})
}
