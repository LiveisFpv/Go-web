package main

import (
	"backend/internal/app"
	"backend/internal/ports/httpgin"
	"backend/internal/repository"
	"context"

	pgxLogrus "github.com/jackc/pgx-logrus"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	log "github.com/sirupsen/logrus"
)

func main() {
	logger := log.New()
	logger.SetLevel(log.InfoLevel)
	logger.SetFormatter(&log.TextFormatter{})

	config, err := pgxpool.ParseConfig("postgres://postgres:0000@localhost:5432/University_DB")
	if err != nil {
		logger.WithError(err).Fatalf("can't parse pgxpool config")
	}
	config.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger:   pgxLogrus.NewLogger(logger),
		LogLevel: tracelog.LogLevelDebug,
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		logger.WithError(err).Fatalf("can't create new pool")
	}
	defer pool.Close()

	repo := repository.NewRepository(pool, logger)
	usecase := app.NewApp(repo)
	server := httpgin.NewHTTPServer(":15432", usecase)
	err = server.Listen()
	if err != nil {
		panic(err)
	}
}
