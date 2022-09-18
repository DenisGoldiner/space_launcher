package entities

import "time"

// User domain entity.
type User struct {
	ID        string
	FirstName string
	LastName  string
	Gender    Gender
	Birthday  time.Time
}
