package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// Logger creates a logging middleware
func Logger(logger *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Start timer
		start := time.Now()

		// Get request ID from context
		requestID, _ := c.Locals("request_id").(string)

		// Process request
		err := c.Next()

		// Calculate duration
		duration := time.Since(start)

		// Log request
		logger.Info("request completed",
			zap.String("request_id", requestID),
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("duration", duration),
			zap.String("ip", c.IP()),
		)

		return err
	}
}
