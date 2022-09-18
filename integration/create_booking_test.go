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
			expResponseErr: fmt.Sprintln("there is no request body, but it is expected"),
		},
		//"empty_first_name":     {},
		//"empty_last_name":      {},
		//"empty_birthday":       {},
		//"empty_gender":         {},
		//"not_supported_gender": {},
		//"empty_launchpad_id":   {},
		//"empty_launch_date":    {},
		//"empty_destination":    {},
	}

	// we won't reach the DB as all cases will fail the validation.
	dbExec := setupDB(t)
	router := newTestRouter(dbExec)

	for tcName, tc := range testCases {
		tc := tc

		t.Run(tcName, func(t *testing.T) {
			req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, bookingsURL, tc.body)
			require.NoError(t, err)
			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)

			require.Equal(t, http.StatusBadRequest, recorder.Code)
			require.Equal(t, tc.expResponseErr, recorder.Body.String())
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
