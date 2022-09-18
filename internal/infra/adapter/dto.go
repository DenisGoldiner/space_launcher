package adapter

import "time"

// QueryOptions dto to do query requests.
type QueryOptions struct {
	Query `json:"query"`
	//Options map[string]any `json:"options"`
}

// Query dto describing specific query search.
type Query struct {
	LaunchpadID string    `json:"launchpad"`
	DateUTC     TimeRange `json:"date_utc"`
}

// TimeRange dto describing search date range.
type TimeRange struct {
	From time.Time `json:"$gte"`
	To   time.Time `json:"$lt"`
}

// Launchpad dto for parsing the response.
type Launchpad struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Region   string `json:"region"`
	Status   string `json:"status"`
}

// LaunchQueryResponse dto for parsing the response.
type LaunchQueryResponse struct {
	Docs []Launch `json:"docs"`
}

// Launch dto for parsing the response.
type Launch struct {
	ID          string    `json:"id"`
	LaunchpadID string    `json:"launchpad"`
	DateUTC     time.Time `json:"date_utc"`
}
