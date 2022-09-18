package entities

import "time"

type Launch struct {
	ID          string
	LaunchpadID string
	Destination Destination
	LaunchDate  time.Time
	UserID      string
}
