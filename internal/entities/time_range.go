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

func ToMiddleWeekRange(t time.Time) TimeRange {
	return TimeRange{
		From: t.Add(-4 * dayDuration),
		To:   t.Add(4 * dayDuration),
	}
}
