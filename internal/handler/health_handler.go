package handler

import (
	"user-profile-api/internal/models"

	"github.com/gofiber/fiber/v2"
)

// HealthHandler handles health check requests
type HealthHandler struct{}

// NewHealthHandler creates a new health handler
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// Default handles GET /
func (h *HealthHandler) Default(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(models.HealthResponse{
		Status: "ok",
	})
}


// Check handles GET /health
func (h *HealthHandler) Check(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(models.HealthResponse{
		Status: "ok",
	})
}
