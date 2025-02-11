package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/serj213/bookService/internal/config"
	"github.com/serj213/bookService/pkg/pg"
)

const (
	local = "local"
	dev = "dev"
)

func main(){
	cfg, err := config.GetConfig()

	if err != nil {
		panic(err)
	}

	log := setupLogger(cfg.Env)

	log = log.With(slog.String("service", "bookService"))

	log.Info("logger enabled")

	pgDb, err := pg.Deal(cfg.Dsn)
	if err != nil {
		log.Error(fmt.Sprintf("failed to connect to postgres: %v", err))
		panic(err)
	}

	log.Info("postgres connect succesfully")

	_ = pgDb

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch(env) {
	case local:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case dev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default: 
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}