package models

import "time"

// CalculateAge calculates age from date of birth
// Returns the age in years, accounting for whether the birthday has occurred this year
func CalculateAge(dob time.Time) int {
	now := time.Now()
	
	// Calculate years difference
	age := now.Year() - dob.Year()
	
	// Adjust if birthday hasn't occurred yet this year
	// Check if current month is before birth month, or
	// same month but current day is before birth day
	if now.Month() < dob.Month() || (now.Month() == dob.Month() && now.Day() < dob.Day()) {
		age--
	}
	
	return age
}

// ParseDate parses a date string in YYYY-MM-DD format
func ParseDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr)
}

// FormatDate formats a time.Time to YYYY-MM-DD string
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}
