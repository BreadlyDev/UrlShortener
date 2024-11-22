package main

import (
	"log/slog"
	"urlshortener/internal/config"
	"urlshortener/internal/utils/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env)

	log.Info("starting url-shortener", slog.String("env", cfg.Env))

	// TODO: init storage: sqlite

	// TODO: init router: chi, "chi render"

	// TODO: run server
}
