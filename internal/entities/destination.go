package entities

import (
	"errors"
	"fmt"

	"github.com/DenisGoldiner/space_launcher/pkg"
)

const (
	marsDest         = "Mars"
	moonDest         = "Moon"
	plutoDest        = "Pluto"
	asteroidBeltDest = "Asteroid Belt"
	europaDest       = "Europa"
	titanDest        = "Titan"
	ganymedeDest     = "Ganymede"
)

// NotSupportedDestinationErr validation error for the enum Destination.
var NotSupportedDestinationErr = errors.New("destination not supported")

// Destination enum for launch destinations.
type Destination string

// Validate checks if the value is a supported launch destination.
func (d Destination) Validate() error {
	supportedDestinations := [...]Destination{
		marsDest,
		moonDest,
		plutoDest,
		asteroidBeltDest,
		europaDest,
		titanDest,
		ganymedeDest,
	}

	for _, expDest := range supportedDestinations {
		if d == expDest {
			return nil
		}
	}

	return pkg.WrapErr(fmt.Sprintf("for the value %q", d), NotSupportedDestinationErr)
}
