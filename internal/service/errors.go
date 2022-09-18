package service

import "errors"

var (
	BusinessValidationErr = errors.New("the business validation was failed")
	RetiredLaunchpadErr   = errors.New("can not plan booking for retired launchpad")
	TakenDateErr          = errors.New("the launch date is planned by an external vendor")
	ExternalVendorAPIErr  = errors.New("failed to call external vendor")
)
