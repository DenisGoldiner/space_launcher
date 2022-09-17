package entities

import "time"

type Launch struct {
	ID          string
	LaunchpadID string
	Destination string
	LaunchDate  time.Time
	UserID      string
}
