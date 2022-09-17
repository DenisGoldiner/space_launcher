package api

import (
	"github.com/DenisGoldiner/space_launcher/internal/entities"
	"strings"
	"time"
)

const (
	dateLayout = "2006-01-02"
)

type LaunchesByUsers map[UserResource][]LaunchResource

func launchByUsersToResource(allBookings map[entities.User][]entities.Launch) LaunchesByUsers {
	allBookingsResource := make(LaunchesByUsers, len(allBookings))

	for user, launches := range allBookings {
		launchesResource := make([]LaunchResource, len(launches))
		for i := range launches {
			launchesResource[i] = LaunchResource{
				LaunchpadID: launches[i].LaunchpadID,
				Destination: launches[i].Destination,
				LaunchDate:  Date(launches[i].LaunchDate),
			}
		}

		userResource := UserResource{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Gender:    user.Gender,
			Birthday:  Date(user.Birthday),
		}

		allBookingsResource[userResource] = launchesResource
	}

	return allBookingsResource
}

// BookingResource represents the request body for create launch booking.
type BookingResource struct {
	UserResource
	LaunchResource
}

// UserResource represents the user information.
type UserResource struct {
	FirstName string          `json:"first_name"`
	LastName  string          `json:"last_name"`
	Gender    entities.Gender `json:"gender"`
	Birthday  Date            `json:"birthday"`
}

// TODO: make the Destination an entity in the DB with its ID

// LaunchResource represents the launch information.
type LaunchResource struct {
	LaunchpadID string `json:"launchpad_id"`
	Destination string `json:"destination"`
	LaunchDate  Date   `json:"launch_date"`
}

// TODO: add validation

type Date time.Time

func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	date, err := time.Parse(dateLayout, s)
	if err != nil {
		return err
	}

	*d = Date(date)

	return nil
}
