package routes

import (
	"Arbitrax/middleware"
	"Arbitrax/output"
	"net/http"

	"github.com/gorilla/mux"
)

func MakeSubRouter(r *mux.Router, path string, apply func(*mux.Router)) {
	sr := r.PathPrefix(path).Subrouter()
	apply(sr)
}

func MakeRoute(
	r *mux.Router,
	path string,
	handler output.JsonHandler,
	middlewares ...middleware.Middleware,
) *mux.Route {
	var h http.Handler = output.MakeJsonHandler(handler)

	// Apply middleware in order
	for _, mw := range middlewares {
		h = mw(h)
	}

	return r.Handle(path, h)
}
