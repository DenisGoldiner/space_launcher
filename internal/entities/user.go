package entities

import "time"

type User struct {
	ID        string
	FirstName string
	LastName  string
	Gender    Gender
	Birthday  time.Time
}
