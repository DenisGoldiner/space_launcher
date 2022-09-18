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
		body   io.Reader
		expErr string
	}{
		"no_body": {
			body:   http.NoBody,
			expErr: "there is no request body, but it is expected",
		},
		"empty_first_name": {
			expErr: "for the user First Name; field should not be empty; the request validation failed",
		},
		"empty_last_name": {
			expErr: "for the user Last Name; field should not be empty; the request validation failed",
		},
		"empty_birthday": {
			expErr: "for the user Birthday; field should not be empty; the request validation failed",
		},
		"invalid_birthday_format": {
			expErr: "parsing time \"2000-01-08T15:04:05Z07:00\": extra text: \"T15:04:05Z07:00\"; failed to decode the body",
		},
		"empty_gender": {
			expErr: "for the value \"\"; gender not supported; the request validation failed",
		},
		"not_supported_gender": {
			expErr: "for the value \"qwerty\"; gender not supported; the request validation failed",
		},
		"empty_launchpad_id": {
			expErr: "for the launchpad id; field should not be empty; the request validation failed",
		},
		"empty_launch_date": {
			expErr: "for the launch date; field should not be empty; the request validation failed",
		},
		"empty_destination": {
			expErr: "for the value \"\"; destination not supported; the request validation failed",
		},
		"not_supported_destination": {
			expErr: "for the value \"Earth\"; destination not supported; the request validation failed",
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
			require.Equal(t, fmt.Sprintln(tc.expErr), recorder.Body.String())
		})
	}
}

func Test_integration_CreateBooking_businessValidation(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		expErr string
	}{
		"not_existing_launchpad": {
			expErr: "the launchpad does not exist; the business validation was failed",
		},
		"retired_launchpad": {
			expErr: "can not plan booking for retired launchpad; the business validation was failed",
		},
		"planned_external": {
			expErr: "external booking; the launch date is planned; the business validation was failed",
		},
		"planned_internal": {
			expErr: "internal booking; the launch date is planned; the business validation was failed",
		},
		"destination_duplicate": {
			expErr: "exists for 2021-01-01T00:00:00Z; the launch destination is planned for close dates; the business validation was failed",
		},
	}

	// we will fail without persisting any data to DB.
	dbExec := setupDB(t)
	router := newTestRouter(dbExec)

	for tcName, tc := range testCases {
		tcName, tc := tcName, tc

		t.Run(tcName, func(t *testing.T) {
			body := getCreateBookingBody(t, tcName)

			req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, bookingsURL, body)
			require.NoError(t, err)
			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)

			// final test result comparison
			require.Equal(t, http.StatusBadRequest, recorder.Code)
			require.Equal(t, fmt.Sprintln(tc.expErr), recorder.Body.String())
		})
	}
}

func Test_integration_CreateBooking_ok(t *testing.T) {
	dbExec := setupDB(t)

	body := getCreateBookingBody(t, "ok")
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, bookingsURL, body)
	require.NoError(t, err)
	recorder := httptest.NewRecorder()
	router := newTestRouter(dbExec)

	router.ServeHTTP(recorder, req)

	// final test result comparison
	require.Equal(t, http.StatusNoContent, recorder.Code)
}
