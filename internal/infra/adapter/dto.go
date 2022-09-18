package adapter

import "time"

// QueryOptions
// Body:
//
//	{
//	 	"query": {
//	 	  	"launchpad":"5e9e4502f5090995de566f86",
//	 	  	"date_utc": {
//	 	  	    "$gte": "2008-08-03T00:00:00.000Z",
//	 	  	    "$lt": "2008-08-04T00:00:00.000Z"
//	 	  	}
//	 	},
//	 	"options": {}
//	}
type QueryOptions struct {
	Query `json:"query"`
	//Options map[string]any `json:"options"`
}

type Query struct {
	LaunchpadID string    `json:"launchpad"`
	DateUTC     TimeRange `json:"date_utc"`
}

type TimeRange struct {
	From time.Time `json:"$gte"`
	To   time.Time `json:"$lt"`
}

type Launchpad struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Region   string `json:"region"`
	Status   string `json:"status"`
}

type LaunchQueryResponse struct {
	Docs []Launch `json:"docs"`
}

type Launch struct {
	ID          string    `json:"id"`
	LaunchpadID string    `json:"launchpad"`
	DateUTC     time.Time `json:"date_utc"`
}
