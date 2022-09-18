package api

import (
	"fmt"
	"strings"
	"time"

	"github.com/DenisGoldiner/space_launcher/internal/entities"
	"github.com/DenisGoldiner/space_launcher/pkg"
)

const dateLayout = "2006-01-02"

// LaunchesByUserResource represents response for GetBookings.
type LaunchesByUserResource struct {
	UserResource
	Launches []LaunchResource
}

func launchByUsersToResource(allBookings map[entities.User][]entities.Launch) []LaunchesByUserResource {
	allBookingsResource := make([]LaunchesByUserResource, 0, len(allBookings))

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

		allBookingsResource = append(allBookingsResource, LaunchesByUserResource{
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

// Validate checks if the all BookingResource components are valid.
func (br BookingResource) Validate() error {
	if err := br.UserResource.Validate(); err != nil {
		return pkg.WrapErr(err.Error(), ValidationFailedErr)
	}

	if err := br.LaunchResource.Validate(); err != nil {
		return pkg.WrapErr(err.Error(), ValidationFailedErr)
	}

	return nil
}

// UserResource represents the user information.
type UserResource struct {
	FirstName string          `json:"first_name"`
	LastName  string          `json:"last_name"`
	Gender    entities.Gender `json:"gender"`
	Birthday  DateResource    `json:"birthday"`
}

// Validate checks if the all UserResource components are valid.
func (ur UserResource) Validate() error {
	if ur.FirstName == "" {
		return pkg.WrapErr("for the user First Name", EmptyFieldErr)
	}

	if ur.LastName == "" {
		return pkg.WrapErr("for the user Last Name", EmptyFieldErr)
	}

	if err := ur.Birthday.Validate(); err != nil {
		return pkg.WrapErr("for the user Birthday", err)
	}

	if err := ur.Gender.Validate(); err != nil {
		return err
	}

	return nil
}

func (ur UserResource) toEntitiesUser() entities.User {
	return entities.User{
		FirstName: ur.FirstName,
		LastName:  ur.LastName,
		Gender:    ur.Gender,
		Birthday:  time.Time(ur.Birthday),
	}
}

// LaunchResource represents the launch information.
type LaunchResource struct {
	LaunchpadID string               `json:"launchpad_id"`
	Destination entities.Destination `json:"destination"`
	LaunchDate  DateResource         `json:"launch_date"`
}

// Validate checks if the all LaunchResource components are valid.
func (lr LaunchResource) Validate() error {
	if lr.LaunchpadID == "" {
		return pkg.WrapErr("for the launchpad id", EmptyFieldErr)
	}

	if err := lr.LaunchDate.Validate(); err != nil {
		return pkg.WrapErr("for the launch date", err)
	}

	if err := lr.Destination.Validate(); err != nil {
		return err
	}

	return nil
}

func (lr LaunchResource) toEntitiesLaunch() entities.Launch {
	return entities.Launch{
		LaunchpadID: lr.LaunchpadID,
		Destination: lr.Destination,
		LaunchDate:  time.Time(lr.LaunchDate),
	}
}

type DateResource time.Time

// UnmarshalJSON decodes the custom date format.
func (d *DateResource) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	date, err := time.Parse(dateLayout, s)
	if err != nil {
		return err
	}

	*d = DateResource(date)

	return nil
}

// MarshalJSON encodes the custom date format.
func (d *DateResource) MarshalJSON() ([]byte, error) {
	t := time.Time(*d)
	return []byte(fmt.Sprintf("%q", t.Format(dateLayout))), nil
}

// Validate checks that the date has not a default value.
func (d *DateResource) Validate() error {
	if time.Time(*d).IsZero() {
		return EmptyFieldErr
	}

	return nil
}
