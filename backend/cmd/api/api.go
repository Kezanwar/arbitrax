package main

import (
	"Arbitrax/cmd/api/handlers"
	"Arbitrax/cmd/api/routes"
	"Arbitrax/pkg/middleware"
	user_repo "Arbitrax/pkg/repositories/user"
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewAPI(ctx context.Context, pool *pgxpool.Pool) (*http.Server, error) {
	//repositories
	userRepo := user_repo.NewUserRepo(pool)

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
		Addr:    PORT,
		Handler: r,
	}, nil
}
