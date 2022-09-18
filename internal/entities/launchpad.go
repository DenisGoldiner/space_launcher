package entities

// LaunchpadStatusRetired is for retired launchpads.
const LaunchpadStatusRetired = "retired"

// LaunchpadID is used for better code self documentation.
type LaunchpadID = string

// Launchpad domain entity.
type Launchpad struct {
	ID       string
	Name     string
	FullName string
	Region   string
	Status   string
}

// IsZero compares the Launchpad value to the default one.
func (lp Launchpad) IsZero() bool {
	return lp == Launchpad{}
}
