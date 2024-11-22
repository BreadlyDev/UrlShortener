package main

import (
	"log/slog"
	"os"
	"urlshortener/internal/config"
	"urlshortener/internal/lib/logger/sl"
	"urlshortener/internal/storage/sqlite"
	"urlshortener/internal/utils/logger"
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

	// TODO: init storage: sqlite

	// TODO: init router: chi, "chi render"

	// TODO: run server
}
