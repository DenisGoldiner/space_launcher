package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	// ContentTypeKey is http header key for content type.
	ContentTypeKey = "Content-Type"

	// ContentTypeValueJSON is http header value for application/json.
	ContentTypeValueJSON = "application/json; charset=utf-8"
)

// SpaceLauncherHTTPHandler is handler for bookings endpoints
type SpaceLauncherHTTPHandler struct {
}

// ServeHTTP implements the http.Handler interface.
// it adds all supported endpoints to the router.
func (slh SpaceLauncherHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		slh.GetBookings(w, r)
	case http.MethodPost:
		slh.CreateBooking(w, r)
	case http.MethodDelete:
		slh.DeleteBooking(w, r)
	default:
		http.Error(w, fmt.Sprintf("tried to run %q method", r.Method), http.StatusMethodNotAllowed)
	}
}

func (slh SpaceLauncherHTTPHandler) GetBookings(w http.ResponseWriter, r *http.Request) {
	log.Println("GetBookings")

	if err := WriteJSON(w, http.StatusOK, nil); err != nil {
		http.Error(w, fmt.Sprintf("output serialization failed, %v", err), http.StatusInternalServerError)
	}
}

func (slh SpaceLauncherHTTPHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateBooking")

	w.WriteHeader(http.StatusNoContent)
}

func (slh SpaceLauncherHTTPHandler) DeleteBooking(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteBooking")

	w.WriteHeader(http.StatusNoContent)
}

// WriteJSON with http.ResponseWriter.
func WriteJSON(w http.ResponseWriter, status int, response interface{}) error {
	w.Header().Set(ContentTypeKey, ContentTypeValueJSON)
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(response)
}
