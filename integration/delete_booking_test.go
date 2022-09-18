package integration

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_integration_DeleteBooking_ok(t *testing.T) {
	dbExec := setupDB(t)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodDelete, bookingsURL, http.NoBody)
	require.NoError(t, err)
	recorder := httptest.NewRecorder()
	router := newTestRouter(dbExec)

	router.ServeHTTP(recorder, req)

	require.Equal(t, http.StatusNoContent, recorder.Code)
}
