package routes

import (
	"user-profile-api/internal/handler"
	"user-profile-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
)

// Setup configures all application routes and middleware
func Setup(app *fiber.App, userHandler *handler.UserHandler, healthHandler *handler.HealthHandler, logger *zap.Logger) {
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(middleware.RequestID())
	app.Use(middleware.Logger(logger))
	app.Use(middleware.ErrorHandler(logger))

	// Health check endpoint
	app.Get("/", healthHandler.Default)
	app.Get("/health", healthHandler.Check)

	// API routes
	api := app.Group("/users")
	{
		api.Post("/", userHandler.CreateUser)
		api.Get("/", userHandler.ListUsers)
		api.Get("/:id", userHandler.GetUser)
		api.Put("/:id", userHandler.UpdateUser)
		api.Delete("/:id", userHandler.DeleteUser)
	}
}
