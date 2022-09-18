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

func Test_integration_DeleteBooking_requestValidation(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		body   io.Reader
		expErr string
	}{
		"no_body": {
			body:   http.NoBody,
			expErr: "there is no request body, but it is expected",
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

//func Test_integration_DeleteBooking_ok(t *testing.T) {
//	dbExec := setupDB(t)
//
//	req, err := http.NewRequestWithContext(context.Background(), http.MethodDelete, bookingsURL, http.NoBody)
//	require.NoError(t, err)
//	recorder := httptest.NewRecorder()
//	router := newTestRouter(dbExec)
//
//	router.ServeHTTP(recorder, req)
//
//	require.Equal(t, http.StatusNoContent, recorder.Code)
//}
