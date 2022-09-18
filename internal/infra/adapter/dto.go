package adapter

type Launchpad struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Region   string `json:"region"`
	Status   string `json:"status"`
}
