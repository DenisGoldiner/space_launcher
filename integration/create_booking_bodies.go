package integration

import (
	"io"
	"strings"
	"testing"
)

func getCreateBookingBody(t *testing.T, caseName string) io.Reader {
	switch caseName {
	case "empty_first_name":
		return createBookingEmptyFirstNameBody()
	case "empty_last_name":
		return createBookingEmptyLastNameBody()
	case "empty_birthday":
		return createBookingEmptyBirthdayBody()
	case "invalid_birthday_format":
		return createBookingInvalidBirthdayFormatBody()
	case "empty_gender":
		return createBookingEmptyGenderBody()
	case "not_supported_gender":
		return createBookingNotSupporterGenderBody()
	case "empty_launchpad_id":
		return createBookingEmptyLaunchpadBody()
	case "empty_launch_date":
		return createBookingEmptyLaunchDateBody()
	case "empty_destination":
		return createBookingEmptyDestinationBody()
	case "not_supported_destination":
		return createBookingNotSupportedDestinationBody()
	default:
		t.Fail()
		return nil
	}
}

func createBookingEmptyFirstNameBody() io.Reader {
	return strings.NewReader(`
		{
			"last_name": "Bond",
			"gender": "male",
			"birthday": "2000-01-08",
			"launchpad_id": "5e9e4502f5090995de566f86",
			"destination": "Mars",
			"launch_date": "2008-08-03"
		}`)
}

func createBookingEmptyLastNameBody() io.Reader {
	return strings.NewReader(`
		{
			"first_name": "James",
			"gender": "male",
			"birthday": "2000-01-08",
			"launchpad_id": "5e9e4502f5090995de566f86",
			"destination": "Mars",
			"launch_date": "2008-08-03"
		}`)
}

func createBookingEmptyBirthdayBody() io.Reader {
	return strings.NewReader(`
		{
			"first_name": "James",
			"last_name": "Bond",
			"gender": "male",
			"launchpad_id": "5e9e4502f5090995de566f86",
			"destination": "Mars",
			"launch_date": "2008-08-03"
		}`)
}

func createBookingInvalidBirthdayFormatBody() io.Reader {
	return strings.NewReader(`
		{
			"first_name": "James",
			"last_name": "Bond",
			"birthday": "2000-01-08T15:04:05Z07:00",
			"gender": "male",
			"launchpad_id": "5e9e4502f5090995de566f86",
			"destination": "Mars",
			"launch_date": "2008-08-03"
		}`)
}

func createBookingEmptyGenderBody() io.Reader {
	return strings.NewReader(`
		{
			"first_name": "James",
			"last_name": "Bond",
			"birthday": "2000-01-08",
			"launchpad_id": "5e9e4502f5090995de566f86",
			"destination": "Mars",
			"launch_date": "2008-08-03"
		}`)
}

func createBookingNotSupporterGenderBody() io.Reader {
	return strings.NewReader(`
		{
			"first_name": "James",
			"last_name": "Bond",
			"gender": "qwerty",
			"birthday": "2000-01-08",
			"launchpad_id": "5e9e4502f5090995de566f86",
			"destination": "Mars",
			"launch_date": "2008-08-03"
		}`)
}

func createBookingEmptyLaunchpadBody() io.Reader {
	return strings.NewReader(`
		{
			"first_name": "James",
			"last_name": "Bond",
			"gender": "male",
			"birthday": "2000-01-08",
			"destination": "Mars",
			"launch_date": "2008-08-03"
		}`)
}

func createBookingEmptyLaunchDateBody() io.Reader {
	return strings.NewReader(`
		{
			"first_name": "James",
			"last_name": "Bond",
			"gender": "male",
			"birthday": "2000-01-08",
			"launchpad_id": "5e9e4502f5090995de566f86",
			"destination": "Mars"
		}`)
}

func createBookingEmptyDestinationBody() io.Reader {
	return strings.NewReader(`
		{
			"first_name": "James",
			"last_name": "Bond",
			"gender": "male",
			"birthday": "2000-01-08",
			"launchpad_id": "5e9e4502f5090995de566f86",
			"launch_date": "2008-08-03"
		}`)
}

func createBookingNotSupportedDestinationBody() io.Reader {
	return strings.NewReader(`
		{
			"first_name": "James",
			"last_name": "Bond",
			"gender": "male",
			"birthday": "2000-01-08",
			"launchpad_id": "5e9e4502f5090995de566f86",
			"destination": "Earth",
			"launch_date": "2008-08-03"
		}`)
}

//func createBookingOkBody() io.Reader {
//	return strings.NewReader(`
//		{
//			"first_name": "James",
//			"last_name": "Bond",
//			"gender": "male",
//			"birthday": "2000-01-08",
//			"launchpad_id": "5e9e4502f5090995de566f86",
//			"destination": "Mars",
//			"launch_date": "2008-08-03"
//		}`)
//}
