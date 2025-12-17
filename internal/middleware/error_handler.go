package middleware

import (
	"user-profile-api/internal/models"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// ErrorHandler creates a centralized error handling middleware
func ErrorHandler(logger *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		
		err := c.Next()

		if err == nil {
			return nil
		}

		requestID, _ := c.Locals("request_id").(string)

		logger.Error("request error",
			zap.String("request_id", requestID),
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Error(err),
		)

		// Handle Fiber errors
		if e, ok := err.(*fiber.Error); ok {
			return c.Status(e.Code).JSON(models.ErrorResponse{
				Error: e.Message,
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error: "internal server error",
		})
	}
}
