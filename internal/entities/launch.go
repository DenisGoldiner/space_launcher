package entities

import "time"

// Launch domain entity.
type Launch struct {
	ID          string
	LaunchpadID string
	Destination Destination
	LaunchDate  time.Time
	UserID      string
}
