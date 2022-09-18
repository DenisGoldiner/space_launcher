package entities

const (
	LaunchpadStatusRetired = "retired"
)

type LaunchpadID = string

type Launchpad struct {
	ID       string
	Name     string
	FullName string
	Region   string
	Status   string
}

func (lp Launchpad) IsZero() bool {
	return lp == Launchpad{}
}
