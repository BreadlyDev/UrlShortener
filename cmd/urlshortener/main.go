package main

import (
	"log/slog"
	"os"
	"urlshortener/internal/config"
	mwLogger "urlshortener/internal/http-server/middleware/logger"
	"urlshortener/internal/lib/logger/sl"
	"urlshortener/internal/storage/sqlite"
	"urlshortener/internal/utils/logger"

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
	router.Use(mwLogger.New(log))

	// middleware

	// TODO: init storage: sqlite

	// TODO: init router: chi, "chi render"

	// TODO: run server
}
