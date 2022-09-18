package entities

import (
	"errors"
	"fmt"
	"github.com/DenisGoldiner/space_launcher/pkg"
)

const (
	MarsDest         = "Mars"
	MoonDest         = "Moon"
	PlutoDest        = "Pluto"
	AsteroidBeltDest = "Asteroid Belt"
	EuropaDest       = "Europa"
	TitanDest        = "Titan"
	GanymedeDest     = "Ganymede"
)

var (
	NotSupportedDestinationErr = errors.New("destination not supported")
)

type Destination string

func (d Destination) Validate() error {
	supportedDestinations := [...]Destination{
		MarsDest,
		MoonDest,
		PlutoDest,
		AsteroidBeltDest,
		EuropaDest,
		TitanDest,
		GanymedeDest,
	}

	for _, expDest := range supportedDestinations {
		if d == expDest {
			return nil
		}
	}

	return pkg.WrapErr(fmt.Sprintf("for the value %q", d), NotSupportedDestinationErr)
}
