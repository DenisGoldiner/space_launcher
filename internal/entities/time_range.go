package entities

import "time"

const (
	dayDuration = 24 * time.Hour
)

type TimeRange struct {
	From time.Time
	To   time.Time
}

func ToDayRange(t time.Time) TimeRange {
	return TimeRange{
		From: t,
		To:   t.Add(dayDuration),
	}
}
