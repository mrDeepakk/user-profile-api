package service

import (
	"context"
	"fmt"
	"time"

	"user-profile-api/internal/models"
	"user-profile-api/internal/repository"

	"go.uber.org/zap"
)

// UserService handles business logic for user operations
type UserService struct {
	repo   repository.Repository
	logger *zap.Logger
}

// NewUserService creates a new user service
func NewUserService(repo repository.Repository, logger *zap.Logger) *UserService {
	return &UserService{
		repo:   repo,
		logger: logger,
	}
}

// CreateUser creates a new user
func (s *UserService) CreateUser(ctx context.Context, req *models.CreateUserRequest) (*models.CreateUserResponse, error) {
	// Parse and validate DOB
	dob, err := models.ParseDate(req.DOB)
	if err != nil {
		return nil, fmt.Errorf("invalid date format: %w", err)
	}

	// Validate DOB is not in the future
	if dob.After(time.Now()) {
		return nil, fmt.Errorf("date of birth cannot be in the future")
	}

	// Create user in repository
	user, err := s.repo.CreateUser(ctx, req.Name, dob)
	if err != nil {
		return nil, err
	}

	// Convert to response without age
	return s.toCreateUserResponse(user), nil
}

// GetUserByID retrieves a user by ID
func (s *UserService) GetUserByID(ctx context.Context, id int32) (*models.UserResponse, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return s.toUserResponse(user), nil
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(ctx context.Context, id int32, req *models.UpdateUserRequest) (*models.CreateUserResponse, error) {
	// Parse and validate DOB
	dob, err := models.ParseDate(req.DOB)
	if err != nil {
		return nil, fmt.Errorf("invalid date format: %w", err)
	}

	// Validate DOB is not in the future
	if dob.After(time.Now()) {
		return nil, fmt.Errorf("date of birth cannot be in the future")
	}

	// Update user in repository
	user, err := s.repo.UpdateUser(ctx, id, req.Name, dob)
	if err != nil {
		return nil, err
	}

	return s.toCreateUserResponse(user), nil
}

// DeleteUser deletes a user by ID
func (s *UserService) DeleteUser(ctx context.Context, id int32) error {
	return s.repo.DeleteUser(ctx, id)
}

// ListUsers retrieves a paginated list of users
func (s *UserService) ListUsers(ctx context.Context, limit, offset int32) ([]models.UserResponse, error) {
	// Set default limit if not provided
	if limit <= 0 {
		limit = 10
	}
	
	// Cap maximum limit to prevent abuse
	if limit > 100 {
		limit = 100
	}

	users, err := s.repo.ListUsers(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	// Convert to response with calculated ages
	responses := make([]models.UserResponse, len(users))
	for i, user := range users {
		responses[i] = *s.toUserResponse(&user)
	}

	return responses, nil
}

// toCreateUserResponse converts a repository user to a create response DTO without age
func (s *UserService) toCreateUserResponse(user *repository.User) *models.CreateUserResponse {
	return &models.CreateUserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  models.FormatDate(user.DOB),
	}
}

// toUserResponse converts a repository user to a response DTO with calculated age
func (s *UserService) toUserResponse(user *repository.User) *models.UserResponse {
	return &models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  models.FormatDate(user.DOB),
		Age:  models.CalculateAge(user.DOB),
	}
}
