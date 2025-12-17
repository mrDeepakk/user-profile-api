package repository

import (
	"context"
	"time"
)

// Repository defines the interface for user data access
type Repository interface {
	CreateUser(ctx context.Context, name string, dob time.Time) (*User, error)
	GetUserByID(ctx context.Context, id int32) (*User, error)
	UpdateUser(ctx context.Context, id int32, name string, dob time.Time) (*User, error)
	DeleteUser(ctx context.Context, id int32) error
	ListUsers(ctx context.Context, limit, offset int32) ([]User, error)
	CountUsers(ctx context.Context) (int64, error)
}

// User represents a user from the database
type User struct {
	ID   int32
	Name string
	DOB  time.Time
}
