package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DenisGoldiner/space_launcher/internal/entities"
	"log"
	"net/http"
)

const (
	// ContentTypeKey is http header key for content type.
	ContentTypeKey = "Content-Type"
	// ContentTypeValueJSON is http header value for application/json.
	ContentTypeValueJSON = "application/json; charset=utf-8"
)

type SpaceLauncherInteractor interface {
	CreateBooking(u entities.User, l entities.Launch) error
}

// SpaceLauncherHTTPHandler is handler for bookings endpoints
type SpaceLauncherHTTPHandler struct {
	Service SpaceLauncherInteractor
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

	payload, err := parseCreateBookingRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logError(err)
		return
	}

	log.Printf("%#v", payload)

	usr := entities.User(payload.UserResource)
	launch := entities.Launch(payload.LaunchResource)

	if err := slh.Service.CreateBooking(usr, launch); err != nil {
		handleCreateBookingError(w, err)
		logError(err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func handleCreateBookingError(w http.ResponseWriter, err error) {
	switch {
	// TODO: define cases
	case errors.Is(err, nil):
		http.Error(w, err.Error(), http.StatusBadRequest)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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

func logError(err error) {
	log.Printf("error: %v", err.Error())
}
