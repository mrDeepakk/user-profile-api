package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

// PostgresRepository implements Repository interface using PostgreSQL
type PostgresRepository struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

// NewPostgresRepository creates a new PostgreSQL repository
func NewPostgresRepository(pool *pgxpool.Pool, logger *zap.Logger) *PostgresRepository {
	return &PostgresRepository{
		pool:   pool,
		logger: logger,
	}
}

// CreateUser creates a new user in the database
func (r *PostgresRepository) CreateUser(ctx context.Context, name string, dob time.Time) (*User, error) {
	query := `INSERT INTO users (name, dob) VALUES ($1, $2) RETURNING id, name, dob`
	
	var user User
	err := r.pool.QueryRow(ctx, query, name, dob).Scan(&user.ID, &user.Name, &user.DOB)
	if err != nil {
		r.logger.Error("failed to create user", zap.Error(err), zap.String("name", name))
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	r.logger.Info("user created", zap.Int32("id", user.ID), zap.String("name", user.Name))
	return &user, nil
}

// GetUserByID retrieves a user by ID
func (r *PostgresRepository) GetUserByID(ctx context.Context, id int32) (*User, error) {
	query := `SELECT id, name, dob FROM users WHERE id = $1`
	
	var user User
	err := r.pool.QueryRow(ctx, query, id).Scan(&user.ID, &user.Name, &user.DOB)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, fmt.Errorf("user not found")
		}
		r.logger.Error("failed to get user", zap.Error(err), zap.Int32("id", id))
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &user, nil
}

// UpdateUser updates an existing user
func (r *PostgresRepository) UpdateUser(ctx context.Context, id int32, name string, dob time.Time) (*User, error) {
	query := `UPDATE users SET name = $1, dob = $2 WHERE id = $3 RETURNING id, name, dob`
	
	var user User
	err := r.pool.QueryRow(ctx, query, name, dob, id).Scan(&user.ID, &user.Name, &user.DOB)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, fmt.Errorf("user not found")
		}
		r.logger.Error("failed to update user", zap.Error(err), zap.Int32("id", id))
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	r.logger.Info("user updated", zap.Int32("id", user.ID), zap.String("name", user.Name))
	return &user, nil
}

// DeleteUser deletes a user by ID
func (r *PostgresRepository) DeleteUser(ctx context.Context, id int32) error {
	query := `DELETE FROM users WHERE id = $1`
	
	result, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		r.logger.Error("failed to delete user", zap.Error(err), zap.Int32("id", id))
		return fmt.Errorf("failed to delete user: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("user not found")
	}

	r.logger.Info("user deleted", zap.Int32("id", id))
	return nil
}

// ListUsers retrieves a list of users with pagination
func (r *PostgresRepository) ListUsers(ctx context.Context, limit, offset int32) ([]User, error) {
	query := `SELECT id, name, dob FROM users ORDER BY id LIMIT $1 OFFSET $2`
	
	rows, err := r.pool.Query(ctx, query, limit, offset)
	if err != nil {
		r.logger.Error("failed to list users", zap.Error(err))
		return nil, fmt.Errorf("failed to list users: %w", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.DOB); err != nil {
			r.logger.Error("failed to scan user", zap.Error(err))
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("error iterating users", zap.Error(err))
		return nil, fmt.Errorf("error iterating users: %w", err)
	}

	// Return empty slice instead of nil if no users found
	if users == nil {
		users = []User{}
	}

	return users, nil
}

// CountUsers returns the total number of users
func (r *PostgresRepository) CountUsers(ctx context.Context) (int64, error) {
	query := `SELECT COUNT(*) FROM users`
	
	var count int64
	err := r.pool.QueryRow(ctx, query).Scan(&count)
	if err != nil {
		r.logger.Error("failed to count users", zap.Error(err))
		return 0, fmt.Errorf("failed to count users: %w", err)
	}

	return count, nil
}
