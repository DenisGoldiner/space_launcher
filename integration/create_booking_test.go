package integration

import (
	"context"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_integration_CreateBooking_ok(t *testing.T) {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, "/bookings", nil)
	require.NoError(t, err)
	recorder := httptest.NewRecorder()
	router := newTestRouter(t)

	router.ServeHTTP(recorder, req)

	require.Equal(t, http.StatusNoContent, recorder.Code)
}
