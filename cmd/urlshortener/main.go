package main

import (
	"log/slog"
	"os"
	"urlshortener/internal/config"
	"urlshortener/internal/lib/logger/sl"
	"urlshortener/internal/storage/sqlite"
	"urlshortener/internal/utils/logger"

	"github.com/docker/docker/api/server/middleware"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env)

	log.Info("starting url-shortener", slog.String("env", cfg.Env))

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	_ = storage

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	_ = router

	// middleware

	// TODO: init storage: sqlite

	// TODO: init router: chi, "chi render"

	// TODO: run server
}
