package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"pc/config"
	"pc/internal/infrastructure/api"
	"pc/internal/infrastructure/persistence"
)

func main() {
	cfg := config.Load()

	db := persistence.DB(cfg)
	router := api.Router(db)

	srv := &http.Server{
		Addr:           cfg.HTTPPort,
		Handler:        router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		log.Printf("⇢ listening on http://0.0.0.0%s", cfg.HTTPPort)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGTSTP)

	sig := <-quit
	log.Printf("signal %s received → shutting down…", sig)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("graceful shutdown failed: %v", err)
	}

	log.Println("server exited cleanly")
}
