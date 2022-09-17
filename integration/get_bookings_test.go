package integration

import (
	"context"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_integration_GetBookings_ok(t *testing.T) {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/bookings", nil)
	require.NoError(t, err)
	recorder := httptest.NewRecorder()
	router := newTestRouter(t)

	expectedBody := "null"

	router.ServeHTTP(recorder, req)

	require.Equal(t, http.StatusOK, recorder.Code)
	require.JSONEq(t, expectedBody, recorder.Body.String())
}