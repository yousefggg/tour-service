package main

import (
	"context"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	jwtlib "github.com/yousefggg/common-lib/pkg/jwt"
	"github.com/yousefggg/common-lib/pkg/logger"

	httpDelivery "github.com/yousefggg/tour-service/internal/delivery"
	"github.com/yousefggg/tour-service/internal/delivery/handler"
	"github.com/yousefggg/tour-service/internal/repository/postgres"
	"github.com/yousefggg/tour-service/internal/usecase"
)

type Config struct {
	DBUrl     string
	JWTSecret string
	JWTTTL    time.Duration
	Port      string
}
// @title Tour Service API
 // @version 1.0
 // @description Microservice for tours and bookings
 // @host localhost:8081
 // @BasePath /api/v1
func main() {


	cfg := Config{
		DBUrl:     "postgres://yousef:20062006@127.0.0.1:5432/tour_db?sslmode=disable",
		JWTSecret: "secret-key",
		JWTTTL:    24 * time.Hour,
		Port:      ":8081",
	}

	logger.Init("debug")

	logger.Info(
		"config loaded",
		"port", cfg.Port,
	)

	runMigrations(cfg.DBUrl)

	ctx := context.Background()

	dbpool, err := pgxpool.New(ctx, cfg.DBUrl)
	if err != nil {
		logger.Error(
			"failed to connect to database",
			"error", err,
		)
		return
	}

	defer dbpool.Close()

	logger.Info("postgres connected successfully")


	jwtManager, err := jwtlib.NewTokenManager(
		cfg.JWTSecret,
		cfg.JWTTTL,
	)
	if err != nil {
		logger.Error(
			"failed to initialize jwt manager",
			"error", err,
		)
		return
	}

	logger.Info("jwt manager initialized")

	
	tourRepo := postgres.NewTourRepository(dbpool)
	bookingRepo := postgres.NewBookingRepository(dbpool)

	logger.Info("repositories initialized")

	tourUC := usecase.NewTourUsecase(tourRepo)
	bookingUC := usecase.NewBookingUsecase(
		bookingRepo,
		tourRepo,
	)

	logger.Info("usecases initialized")

	tourHandler := handler.NewTourHandler(tourUC)
	bookingHandler := handler.NewBookingHandler(bookingUC)

	logger.Info("handlers initialized")

	
	router := httpDelivery.NewRouter(
		tourHandler,
		bookingHandler,
		jwtManager,
	)

	logger.Info("http router initialized")

	server := &http.Server{
		Addr: cfg.Port,
		Handler: router.Setup(),
	}

	logger.Info(
		"http server started",
		"host", "localhost",
		"port", cfg.Port,
	)

	if err := server.ListenAndServe(); err != nil {
		logger.Error(
			"http server stopped",
			"error", err,
		)
	}
}

func runMigrations(dbURL string) {

	logger.Info("starting database migrations")

	m, err := migrate.New(
		"file://migrations",
		dbURL,
	)
	if err != nil {
		logger.Error(
			"failed to initialize migrations",
			"error", err,
		)
		return
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {

		logger.Error(
			"failed to apply migrations",
			"error", err,
		)

		return
	}

	logger.Info("database migrations applied successfully")
}