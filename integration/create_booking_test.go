package integration

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_integration_CreateBooking_requestValidation(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		body           io.Reader
		expResponseErr string
	}{
		"no_body": {
			body:           http.NoBody,
			expResponseErr: "there is no request body, but it is expected",
		},
		"empty_first_name": {
			expResponseErr: "for the user First Name; field should not be empty; the request validation failed",
		},
		"empty_last_name": {
			expResponseErr: "for the user Last Name; field should not be empty; the request validation failed",
		},
		"empty_birthday": {
			expResponseErr: "for the user Birthday; field should not be empty; the request validation failed",
		},
		"invalid_birthday_format": {
			expResponseErr: "parsing time \"2000-01-08T15:04:05Z07:00\": extra text: \"T15:04:05Z07:00\"; failed to decode the body",
		},
		"empty_gender": {
			expResponseErr: "for the value \"\"; gender not supported; the request validation failed",
		},
		"not_supported_gender": {
			expResponseErr: "for the value \"qwerty\"; gender not supported; the request validation failed",
		},
		"empty_launchpad_id": {
			expResponseErr: "for the launchpad id; field should not be empty; the request validation failed",
		},
		"empty_launch_date": {
			expResponseErr: "for the launch date; field should not be empty; the request validation failed",
		},
		"empty_destination": {
			expResponseErr: "for the value \"\"; destination not supported; the request validation failed",
		},
		"not_supported_destination": {
			expResponseErr: "for the value \"Earth\"; destination not supported; the request validation failed",
		},
	}

	// we won't reach the DB as all cases will fail the validation.
	dbExec := setupDB(t)
	router := newTestRouter(dbExec)

	for tcName, tc := range testCases {
		tcName, tc := tcName, tc

		t.Run(tcName, func(t *testing.T) {
			if tc.body == nil {
				tc.body = getCreateBookingBody(t, tcName)
			}

			req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, bookingsURL, tc.body)
			require.NoError(t, err)
			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)

			// final test result comparison
			require.Equal(t, http.StatusBadRequest, recorder.Code)
			require.Equal(t, fmt.Sprintln(tc.expResponseErr), recorder.Body.String())
		})
	}
}

//func Test_integration_CreateBooking_ok(t *testing.T) {
//	dbExec := setupDB(t)
//
//	body := strings.NewReader("{}")
//	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, bookingsURL, body)
//	require.NoError(t, err)
//	recorder := httptest.NewRecorder()
//	router := newTestRouter(dbExec)
//
//	router.ServeHTTP(recorder, req)
//
//	require.Equal(t, http.StatusNoContent, recorder.Code)
//}
