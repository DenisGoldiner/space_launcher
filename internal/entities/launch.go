package entities

import "time"

type Launch struct {
	LaunchpadID string
	Destination string
	LaunchDate  time.Time
}
