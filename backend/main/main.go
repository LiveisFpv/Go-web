package main

import (
	"backend/internal/app"
	"backend/internal/ports/httpgin"
	"backend/internal/repository"
	"context"
	"fmt"
	"os"

	pgxLogrus "github.com/jackc/pgx-logrus"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	log "github.com/sirupsen/logrus"
)

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func main() {
	//Парсим для подключения к БД
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "0000")
	dbName := getEnv("DB_NAME", "University_DB")

	//Подключаемся к БД
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	logger := log.New()
	logger.SetLevel(log.InfoLevel)
	logger.SetFormatter(&log.TextFormatter{})

	config, err := pgxpool.ParseConfig(connStr)
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
