package integration

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_integration_CreateBooking_ok(t *testing.T) {
	dbExec := setupDB(t)

	body := strings.NewReader("{}")
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, bookingsURL, body)
	require.NoError(t, err)
	recorder := httptest.NewRecorder()
	router := newTestRouter(dbExec)

	router.ServeHTTP(recorder, req)

	require.Equal(t, http.StatusNoContent, recorder.Code)
}

//func Test_integration_CreateBooking_emptyBody(t *testing.T) {
//	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, "/bookings", http.NoBody)
//	require.NoError(t, err)
//	recorder := httptest.NewRecorder()
//	router := newTestRouter(t)
//	expectedErr := fmt.Sprintln("there is no request body, but it is expected")
//
//	router.ServeHTTP(recorder, req)
//
//	require.Equal(t, http.StatusBadRequest, recorder.Code)
//	require.Equal(t, expectedErr, recorder.Body.String())
//}
