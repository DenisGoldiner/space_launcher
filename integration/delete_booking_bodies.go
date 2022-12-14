package integration

import (
	"io"
	"strings"
	"testing"
)

func getDeleteBookingBody(t *testing.T, caseName string) io.Reader {
	switch caseName {
	case "empty_launchpad_id":
		return deleteBookingEmptyLaunchpadBody()
	case "empty_launch_date":
		return deleteBookingEmptyLaunchDateBody()
	case "empty_destination":
		return deleteBookingEmptyDestinationBody()
	case "not_supported_destination":
		return deleteBookingNotSupportedDestinationBody()
	case "not_existing_booking":
		return deleteBookingNotExistingBookingBody()
	case "ok":
		return deleteBookingOkBody()
	default:
		t.Fail()
		return nil
	}
}

func deleteBookingEmptyLaunchpadBody() io.Reader {
	return strings.NewReader(`
		{
			"destination": "Mars",
			"launch_date": "2008-08-03"
		}`)
}

func deleteBookingEmptyLaunchDateBody() io.Reader {
	return strings.NewReader(`
		{
			"launchpad_id": "5e9e4502f5090995de566f86",
			"destination": "Mars"
		}`)
}

func deleteBookingEmptyDestinationBody() io.Reader {
	return strings.NewReader(`
		{
			"launchpad_id": "5e9e4502f5090995de566f86",
			"launch_date": "2008-08-03"
		}`)
}

func deleteBookingNotSupportedDestinationBody() io.Reader {
	return strings.NewReader(`
		{
			"launchpad_id": "5e9e4502f5090995de566f86",
			"destination": "Earth",
			"launch_date": "2008-08-03"
		}`)
}

func deleteBookingNotExistingBookingBody() io.Reader {
	return strings.NewReader(`
		{
			"launchpad_id": "5e9e4502f5090995de566f86",
			"destination": "Mars",
			"launch_date": "2008-08-03"
		}`)
}

func deleteBookingOkBody() io.Reader {
	return strings.NewReader(`
		{
			"launchpad_id": "5e9e4501f509094ba4566f84",
			"destination": "Mars",
			"launch_date": "2021-01-01"
		}`)
}
