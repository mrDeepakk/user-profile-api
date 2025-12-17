package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"user-profile-api/config"
	"user-profile-api/internal/handler"
	"user-profile-api/internal/logger"
	"user-profile-api/internal/repository"
	"user-profile-api/internal/routes"
	"user-profile-api/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("Failed to load configuration: %v\n", err)
		os.Exit(1)
	}

	// Initialize logger
	log, err := logger.New(cfg.LogLevel)
	if err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer log.Sync()

	log.Info("starting user profile API",
		zap.String("port", cfg.Port),
		zap.String("log_level", cfg.LogLevel),
	)

	// Connect to database
	poolConfig, err := pgxpool.ParseConfig(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("failed to parse database URL", zap.Error(err))
	}

	dbPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		log.Fatal("failed to connect to database", zap.Error(err))
	}
	defer dbPool.Close()

	// Verify database connection
	if err := dbPool.Ping(context.Background()); err != nil {
		log.Fatal("failed to ping database", zap.Error(err))
	}
	log.Info("successfully connected to database")

	// Initialize layers
	repo := repository.NewPostgresRepository(dbPool, log)
	userService := service.NewUserService(repo, log)
	userHandler := handler.NewUserHandler(userService, log)
	healthHandler := handler.NewHealthHandler()

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName:      "User Profile API",
		ErrorHandler: nil,
	})

	routes.Setup(app, userHandler, healthHandler, log)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		addr := fmt.Sprintf(":%s", cfg.Port)
		log.Info("server starting", zap.String("address", addr))
		if err := app.Listen(addr); err != nil {
			log.Fatal("failed to start server", zap.Error(err))
		}
	}()

	<-quit
	log.Info("shutting down server gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Error("server shutdown error", zap.Error(err))
	}

	log.Info("server stopped")
}
