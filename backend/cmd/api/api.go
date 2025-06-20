package main

import (
	"Arbitrax/cmd/api/handlers"
	"Arbitrax/cmd/api/routes"
	options_memory_cache "Arbitrax/pkg/cache/options_memory"
	user_memory_cache "Arbitrax/pkg/cache/user_memory"
	"Arbitrax/pkg/email"
	"Arbitrax/pkg/middleware"
	agent_repo "Arbitrax/pkg/repositories/agent"
	exchanges_repo "Arbitrax/pkg/repositories/exchanges"
	strategy_repo "Arbitrax/pkg/repositories/strategies"
	user_repo "Arbitrax/pkg/repositories/user"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewAPI(ctx context.Context, pool *pgxpool.Pool, client *http.Client) (*http.Server, error) {

	TWO_HOURS := 2 * time.Hour

	_, err := email.NewClient()

	if err != nil {
		log.Fatalf("Email client failed to init: %v", err)
	}

	//memory cache
	userCache := user_memory_cache.New(TWO_HOURS)
	optionsCache := options_memory_cache.New(TWO_HOURS)

	//repositories
	userRepo := user_repo.NewUserRepo(pool)
	strategyRepo := strategy_repo.NewStrategyRepo(pool)
	exchangeRepo := exchanges_repo.NewExchangeRepo(pool)
	agentsRepo := agent_repo.NewAgentRepo(pool)

	//handlers
	authHandler := handlers.NewAuthHandler(userRepo)
	newsHandler := handlers.NewNewsHandler()
	optionsHandler := handlers.NewOptionsHandler(
		strategyRepo,
		exchangeRepo,
		optionsCache,
	)
	agentsHandler := handlers.NewAgentHandler(agentsRepo, exchangeRepo, strategyRepo)

	authFresh := middleware.AuthAlwaysFreshMiddleware(userRepo, userCache)
	authCached := middleware.AuthCachedMiddleware(userRepo, userCache)

	//router
	r := mux.NewRouter()
	r.Use(middleware.Cors)
	api := r.PathPrefix("/api").Subrouter()

	//apply routes
	routes.Register(
		//main router
		api,
		//handlers
		authHandler,
		newsHandler,
		optionsHandler,
		agentsHandler,
		//middleware
		authFresh,
		authCached,
	)

	return &http.Server{
		Addr:    PORT,
		Handler: r,
	}, nil
}
