package service

import "errors"

var (
	BusinessValidationErr   = errors.New("the business validation was failed")
	RetiredLaunchpadErr     = errors.New("can not plan booking for retired launchpad")
	NotExistingLaunchpadErr = errors.New("the launchpad does not exist")
	TakenDateErr            = errors.New("the launch date is planned")
	TakenDestinationErr     = errors.New("the launch destination is planned for close dates")
	ExternalVendorAPIErr    = errors.New("failed to call external vendor")
)
