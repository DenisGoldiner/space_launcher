package api

import (
	"fmt"
	"github.com/DenisGoldiner/space_launcher/internal/entities"
	"strings"
	"time"
)

const (
	dateLayout = "2006-01-02"
)

type LaunchesByUser struct {
	UserResource
	Launches []LaunchResource
}

func launchByUsersToResource(allBookings map[entities.User][]entities.Launch) []LaunchesByUser {
	allBookingsResource := make([]LaunchesByUser, 0, len(allBookings))

	for user, launches := range allBookings {
		launchesResource := make([]LaunchResource, len(launches))
		for i := range launches {
			launchesResource[i] = LaunchResource{
				LaunchpadID: launches[i].LaunchpadID,
				Destination: launches[i].Destination,
				LaunchDate:  DateResource(launches[i].LaunchDate),
			}
		}

		userResource := UserResource{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Gender:    user.Gender,
			Birthday:  DateResource(user.Birthday),
		}

		allBookingsResource = append(allBookingsResource, LaunchesByUser{
			UserResource: userResource,
			Launches:     launchesResource,
		})
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
	Birthday  DateResource    `json:"birthday"`
}

// TODO: make the Destination an entity in the DB with its ID

// LaunchResource represents the launch information.
type LaunchResource struct {
	LaunchpadID string       `json:"launchpad_id"`
	Destination string       `json:"destination"`
	LaunchDate  DateResource `json:"launch_date"`
}

// TODO: add validation

type DateResource time.Time

func (d *DateResource) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	date, err := time.Parse(dateLayout, s)
	if err != nil {
		return err
	}

	*d = DateResource(date)

	return nil
}

func (d *DateResource) MarshalJSON() ([]byte, error) {
	t := time.Time(*d)
	return []byte(fmt.Sprintf("%q", t.Format(dateLayout))), nil
}
