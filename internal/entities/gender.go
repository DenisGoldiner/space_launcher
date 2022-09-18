package entities

import (
	"errors"
	"fmt"

	"github.com/DenisGoldiner/space_launcher/pkg"
)

const (
	// MaleGender `male` enum value
	MaleGender = "male"
	// FemaleGender `female` enum value
	FemaleGender = "female"
)

var (
	NotSupportedGenderErr = errors.New("gender not supported")
)

// Gender is an enum type needed to cover non-binary genders if needed.
type Gender string

func (g Gender) Validate() error {
	supportedGenders := [...]Gender{MaleGender, FemaleGender}

	for _, expGender := range supportedGenders {
		if g == expGender {
			return nil
		}
	}

	return pkg.WrapErr(fmt.Sprintf("for the value %q", g), NotSupportedGenderErr)
}
