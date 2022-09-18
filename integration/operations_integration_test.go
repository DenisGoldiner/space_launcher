package integration

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_integration_CreateAndGetBooking(t *testing.T) {
	dbExec := setupDB(t)
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
		  },
		  {
			"first_name": "James",
			"last_name": "Bond",
			"gender": "male",
			"birthday": "2000-01-08",
			"Launches": [
			  {
				"launchpad_id": "5e9e4501f509094ba4566f84",
				"destination": "Moon",
				"launch_date": "2021-01-02"
			  }
			]
		  }
		]`

	body := getCreateBookingBody(t, "ok")
	postReq, err := http.NewRequestWithContext(context.Background(), http.MethodPost, bookingsURL, body)
	require.NoError(t, err)
	postRecorder := httptest.NewRecorder()

	getReq, err := http.NewRequestWithContext(context.Background(), http.MethodGet, bookingsURL, http.NoBody)
	require.NoError(t, err)
	getRecorder := httptest.NewRecorder()

	router.ServeHTTP(postRecorder, postReq)
	router.ServeHTTP(getRecorder, getReq)

	require.Equal(t, http.StatusNoContent, postRecorder.Code)
	require.Equal(t, http.StatusOK, getRecorder.Code)
	require.JSONEq(t, expectedBody, getRecorder.Body.String())
}

func Test_integration_DeleteAndGetBooking(t *testing.T) {
	dbExec := setupDB(t)
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
				"launchpad_id": "5e9e4502f509092b78566f87",
				"destination": "Pluto",
				"launch_date": "2021-01-07"
			  }
			]
		  }
		]`

	body := getDeleteBookingBody(t, "ok")
	deleteReq, err := http.NewRequestWithContext(context.Background(), http.MethodDelete, bookingsURL, body)
	require.NoError(t, err)
	deleteRecorder := httptest.NewRecorder()

	getReq, err := http.NewRequestWithContext(context.Background(), http.MethodGet, bookingsURL, http.NoBody)
	require.NoError(t, err)
	getRecorder := httptest.NewRecorder()

	router.ServeHTTP(deleteRecorder, deleteReq)
	router.ServeHTTP(getRecorder, getReq)

	require.Equal(t, http.StatusNoContent, deleteRecorder.Code)
	require.Equal(t, http.StatusOK, getRecorder.Code)
	require.JSONEq(t, expectedBody, getRecorder.Body.String())
}
