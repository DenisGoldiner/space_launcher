package integration

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_integration_GetBookings_ok(t *testing.T) {
	dbExec := setupDB(t)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, bookingsURL, http.NoBody)
	require.NoError(t, err)
	recorder := httptest.NewRecorder()
	router := newTestRouter(dbExec)

	expectedBody := `
		[
		  {
			"first_name": "John",
			"last_name": "Smith",
			"gender": "male",
			"birthday": "1999-01-08",
			"Launches": [
			  {
				"launchpad_id": "5e9e4501f509094ba4566f84",
				"destination": "Mars",
				"launch_date": "2021-01-01"
			  },
			  {
				"launchpad_id": "5e9e4502f509092b78566f87",
				"destination": "Pluto",
				"launch_date": "2021-01-07"
			  }
			]
		  }
		]`

	router.ServeHTTP(recorder, req)

	// final test result comparison
	require.Equal(t, http.StatusOK, recorder.Code)
	require.JSONEq(t, expectedBody, recorder.Body.String())
}
