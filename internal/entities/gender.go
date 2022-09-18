package entities

import (
	"errors"
	"fmt"

	"github.com/DenisGoldiner/space_launcher/pkg"
)

const (
	maleGender   = "male"
	femaleGender = "female"
)

// NotSupportedGenderErr validation error for the enum Gender.
var NotSupportedGenderErr = errors.New("gender not supported")

// Gender is an enum type needed to cover non-binary genders if needed.
type Gender string

// Validate checks if the value is a supported gender.
func (g Gender) Validate() error {
	supportedGenders := [...]Gender{maleGender, femaleGender}

	for _, expGender := range supportedGenders {
		if g == expGender {
			return nil
		}
	}

	return pkg.WrapErr(fmt.Sprintf("for the value %q", g), NotSupportedGenderErr)
}
