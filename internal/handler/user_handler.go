package handler

import (
	"fmt"
	"strconv"

	"user-profile-api/internal/models"
	"user-profile-api/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// UserHandler handles HTTP requests for user operations
type UserHandler struct {
	service  *service.UserService
	logger   *zap.Logger
	validate *validator.Validate
}

// NewUserHandler creates a new user handler
func NewUserHandler(service *service.UserService, logger *zap.Logger) *UserHandler {
	return &UserHandler{
		service:  service,
		logger:   logger,
		validate: validator.New(),
	}
}

// CreateUser handles POST /users
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req models.CreateUserRequest
	
	if err := c.BodyParser(&req); err != nil {
		h.logger.Warn("invalid request body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "invalid request body",
		})
	}

	// Validate request
	if err := h.validate.Struct(&req); err != nil {
		h.logger.Warn("validation failed", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: fmt.Sprintf("validation failed: %v", err),
		})
	}

	// Create user
	user, err := h.service.CreateUser(c.Context(), &req)
	if err != nil {
		h.logger.Error("failed to create user", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

// GetUser handles GET /users/:id
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	
	id, err := strconv.ParseInt(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "invalid user ID",
		})
	}

	// Get user
	user, err := h.service.GetUserByID(c.Context(), int32(id))
	if err != nil {
		if err.Error() == "user not found" {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
				Error: "user not found",
			})
		}
		h.logger.Error("failed to get user", zap.Error(err), zap.Int64("id", id))
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error: "internal server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

// UpdateUser handles PUT /users/:id
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {

	id, err := strconv.ParseInt(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "invalid user ID",
		})
	}

	var req models.UpdateUserRequest
	
	if err := c.BodyParser(&req); err != nil {
		h.logger.Warn("invalid request body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "invalid request body",
		})
	}

	// Validate request
	if err := h.validate.Struct(&req); err != nil {
		h.logger.Warn("validation failed", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: fmt.Sprintf("validation failed: %v", err),
		})
	}

	// Update user
	user, err := h.service.UpdateUser(c.Context(), int32(id), &req)
	if err != nil {
		if err.Error() == "user not found" {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
				Error: "user not found",
			})
		}
		h.logger.Error("failed to update user", zap.Error(err), zap.Int64("id", id))
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

// DeleteUser handles DELETE /users/:id
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {

	id, err := strconv.ParseInt(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "invalid user ID",
		})
	}

	err = h.service.DeleteUser(c.Context(), int32(id))
	if err != nil {
		if err.Error() == "user not found" {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
				Error: "user not found",
			})
		}
		h.logger.Error("failed to delete user", zap.Error(err), zap.Int64("id", id))
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error: "internal server error",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// ListUsers handles GET /users
func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 10)
	offset := c.QueryInt("offset", 0)

	users, err := h.service.ListUsers(c.Context(), int32(limit), int32(offset))
	if err != nil {
		h.logger.Error("failed to list users", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error: "internal server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}
