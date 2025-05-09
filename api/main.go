package main

import (
	"Arbitrax/db"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	ctx := context.Background()

	// Connect to the DB pool
	pool, err := db.Connect(ctx)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}

	db.MigrateUp()

	api, err := NewAPI(ctx, pool)

	if err != nil {
		log.Fatalf("Server init failed: %v", err)
	}

	go func() {
		log.Println("🚀 Server running on :3000")
		if err := api.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("🧼 Shutting down...")
	ctxShutdown, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	api.Shutdown(ctxShutdown)

	db.Close(pool)
}
