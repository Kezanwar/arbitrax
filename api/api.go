package main

import (
	"Arbitrax/handlers"
	"Arbitrax/middleware"
	user_repo "Arbitrax/repositories/user"
	"Arbitrax/routes"
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewAPI(ctx context.Context, pool *pgxpool.Pool) (*http.Server, error) {
	//repositories
	userRepo := user_repo.NewPgxUserRepo(pool)

	//handlers
	authHandler := handlers.NewAuthHandler(userRepo)

	//custom middleware
	authMiddleware := middleware.AuthMiddleware(userRepo)

	//router
	r := mux.NewRouter()
	r.Use(middleware.Cors)
	api := r.PathPrefix("/api").Subrouter()

	//apply routes
	routes.Register(api, authHandler, authMiddleware)

	return &http.Server{
		Addr:    ":3000",
		Handler: r,
	}, nil
}
