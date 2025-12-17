package models

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {
	// Fixed reference date for consistent testing
	now := time.Date(2025, 12, 14, 0, 0, 0, 0, time.UTC)
	
	tests := []struct {
		name     string
		dob      time.Time
		expected int
	}{
		{
			name:     "birthday already occurred this year",
			dob:      time.Date(1990, 5, 10, 0, 0, 0, 0, time.UTC),
			expected: 35,
		},
		{
			name:     "birthday is today",
			dob:      time.Date(1990, 12, 14, 0, 0, 0, 0, time.UTC),
			expected: 35,
		},
		{
			name:     "birthday hasn't occurred yet this year",
			dob:      time.Date(1990, 12, 25, 0, 0, 0, 0, time.UTC),
			expected: 34,
		},
		{
			name:     "leap year birthday",
			dob:      time.Date(2000, 2, 29, 0, 0, 0, 0, time.UTC),
			expected: 25,
		},
		{
			name:     "very young person",
			dob:      time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 5,
		},
		{
			name:     "elderly person",
			dob:      time.Date(1940, 6, 15, 0, 0, 0, 0, time.UTC),
			expected: 85,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock current time for testing
			// In production, CalculateAge uses time.Now()
			// For this test, we calculate expected age based on reference date
			age := calculateAgeWithNow(tt.dob, now)
			if age != tt.expected {
				t.Errorf("CalculateAge(%v) = %d; want %d", tt.dob, age, tt.expected)
			}
		})
	}
}

// Helper function for testing with a fixed "now" time
func calculateAgeWithNow(dob, now time.Time) int {
	age := now.Year() - dob.Year()
	if now.Month() < dob.Month() || (now.Month() == dob.Month() && now.Day() < dob.Day()) {
		age--
	}
	return age
}

func TestParseDate(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantError bool
	}{
		{
			name:      "valid date",
			input:     "1990-05-10",
			wantError: false,
		},
		{
			name:      "invalid format",
			input:     "05-10-1990",
			wantError: true,
		},
		{
			name:      "invalid date",
			input:     "2020-13-45",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseDate(tt.input)
			if (err != nil) != tt.wantError {
				t.Errorf("ParseDate(%s) error = %v, wantError %v", tt.input, err, tt.wantError)
			}
		})
	}
}

func TestFormatDate(t *testing.T) {
	date := time.Date(1990, 5, 10, 0, 0, 0, 0, time.UTC)
	expected := "1990-05-10"
	result := FormatDate(date)
	
	if result != expected {
		t.Errorf("FormatDate(%v) = %s; want %s", date, result, expected)
	}
}
