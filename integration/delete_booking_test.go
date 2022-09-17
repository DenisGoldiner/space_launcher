package integration

import (
	"context"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_integration_DeleteBooking_ok(t *testing.T) {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodDelete, "/bookings", nil)
	require.NoError(t, err)
	recorder := httptest.NewRecorder()
	router := newTestRouter(t)

	expectedBody := "null"

	router.ServeHTTP(recorder, req)

	require.Equal(t, http.StatusNoContent, recorder.Code)
	require.JSONEq(t, expectedBody, recorder.Body.String())
}
