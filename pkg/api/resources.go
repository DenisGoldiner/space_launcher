package api

import (
	"github.com/DenisGoldiner/space_launcher/pkg/entities"
	"time"
)

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
	Birthday  time.Time       `json:"birthday"`
}

// TODO: make the Destination an entity in the DB with its ID

// LaunchResource represents the launch information.
type LaunchResource struct {
	LaunchpadID string    `json:"launchpad_id"`
	Destination string    `json:"destination"`
	LaunchDate  time.Time `json:"launch_date"`
}
