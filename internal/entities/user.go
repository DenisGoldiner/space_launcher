package entities

import "time"

type User struct {
	FirstName string
	LastName  string
	Gender    Gender
	Birthday  time.Time
}
