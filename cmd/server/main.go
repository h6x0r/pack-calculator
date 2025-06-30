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

	"github.com/h6x0r/pack-calculator/config"
	"github.com/h6x0r/pack-calculator/internal/infrastructure/api"
	"github.com/h6x0r/pack-calculator/internal/infrastructure/persistence"
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
		log.Println("server started on port ", cfg.HTTPPort)
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
