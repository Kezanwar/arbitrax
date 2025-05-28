package routes

import (
	"Arbitrax/cmd/api/handlers"
	"Arbitrax/pkg/middleware"
	"Arbitrax/pkg/output"

	"github.com/gorilla/mux"
)

func Register(
	//router
	r *mux.Router,
	//handlers
	authHandler *handlers.AuthHandler,
	newsHandler *handlers.NewsHandler,
	optionsHandler *handlers.OptionsHandler,
	//middlewares
	authFresh middleware.Middleware,
	authCached middleware.Middleware) {

	output.MakeSubRouter(r, "/auth", func(sr *mux.Router) {
		AuthRoutes(sr, authHandler, authCached)
	})
	output.MakeSubRouter(r, "/news", func(sr *mux.Router) {
		NewsRoutes(sr, newsHandler, authCached)
	})
	output.MakeSubRouter(r, "/options", func(sr *mux.Router) {
		OptionsRoutes(sr, optionsHandler, authCached)
	})

}
