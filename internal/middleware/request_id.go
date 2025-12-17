package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// RequestIDHeader is the header name for request ID
const RequestIDHeader = "X-Request-ID"

// RequestID generates a unique request ID for each request
func RequestID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestID := c.Get(RequestIDHeader)
		
		// Generate new UUID if not present
		if requestID == "" {
			requestID = uuid.New().String()
		}

		c.Locals("request_id", requestID)
		c.Set(RequestIDHeader, requestID)

		return c.Next()
	}
}
