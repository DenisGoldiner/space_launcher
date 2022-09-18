package entities

const (
	LaunchpadStatusRetired = "retired"
)

type Launchpad struct {
	ID       string
	Name     string
	FullName string
	Region   string
	Status   string
}
