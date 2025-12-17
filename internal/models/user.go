package models

import "time"

// User represents a user in the system
type User struct {
	ID   int32     `json:"id"`
	Name string    `json:"name"`
	DOB  time.Time `json:"dob"`
	Age  int       `json:"age"` // Calculated field, not stored in DB
}

// CreateUserRequest represents the request body for creating a user
type CreateUserRequest struct {
	Name string `json:"name" validate:"required,min=1,max=255"`
	DOB  string `json:"dob" validate:"required,datetime=2006-01-02"`
}

// UpdateUserRequest represents the request body for updating a user
type UpdateUserRequest struct {
	Name string `json:"name" validate:"required,min=1,max=255"`
	DOB  string `json:"dob" validate:"required,datetime=2006-01-02"`
}

// CreateUserResponse represents the response for creating a user (without age)
type CreateUserResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
}

// UserResponse represents the response for user operations (with age)
type UserResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
	Age  int    `json:"age"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// HealthResponse represents a health check response
type HealthResponse struct {
	Status string `json:"status"`
}
