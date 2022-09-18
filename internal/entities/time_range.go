package entities

import "time"

const dayDuration = 24 * time.Hour

// TimeRange domain entity.
type TimeRange struct {
	From time.Time
	To   time.Time
}

// ToDayRange is used to represent all the day as a range.
func ToDayRange(t time.Time) TimeRange {
	return TimeRange{
		From: t,
		To:   t.Add(dayDuration),
	}
}

// ToMiddleWeekRange describes range of Destination uniqueness.
func ToMiddleWeekRange(t time.Time) TimeRange {
	return TimeRange{
		From: t.Add(-4 * dayDuration),
		To:   t.Add(4 * dayDuration),
	}
}
