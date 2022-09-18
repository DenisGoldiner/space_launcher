package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/DenisGoldiner/space_launcher/internal/entities"
	"github.com/DenisGoldiner/space_launcher/pkg"
)

const (
	// ContentTypeKey is http header key for content type.
	ContentTypeKey = "Content-Type"
	// ContentTypeValueJSON is http header value for application/json.
	ContentTypeValueJSON = "application/json; charset=utf-8"
)

// SpaceLauncherInteractor is a service layer abstraction
type SpaceLauncherInteractor interface {
	CreateBooking(context.Context, entities.User, entities.Launch) error
	GetAllBookings(context.Context) (map[entities.User][]entities.Launch, error)
	DeleteBooking(context.Context, entities.Launch) error
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

// GetBookings
// GET: /space_launcher/bookings
func (slh SpaceLauncherHTTPHandler) GetBookings(w http.ResponseWriter, r *http.Request) {
	log.Println("GetBookings")

	ctx := r.Context()

	allBookings, err := slh.Service.GetAllBookings(ctx)
	if err != nil {
		// all the current errors for get bookings are related to DB interactions
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logError(err)
		return
	}

	if err := WriteJSON(w, http.StatusOK, launchByUsersToResource(allBookings)); err != nil {
		http.Error(w, fmt.Sprintf("output serialization failed, %v", err), http.StatusInternalServerError)
	}
}

// CreateBooking
// POST: /space_launcher/bookings
func (slh SpaceLauncherHTTPHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateBooking")

	ctx := r.Context()

	var payload BookingResource

	if err := pkg.ParseRequestPayload(r, &payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logError(err)
		return
	}

	if err := payload.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logError(err)
		return
	}

	usr := payload.UserResource.toEntitiesUser()
	launch := payload.LaunchResource.toEntitiesLaunch()

	if err := slh.Service.CreateBooking(ctx, usr, launch); err != nil {
		handleCreateBookingError(w, err)
		logError(err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteBooking
// DELETE: /space_launcher/bookings
func (slh SpaceLauncherHTTPHandler) DeleteBooking(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteBooking")

	ctx := r.Context()

	var payload LaunchResource

	if err := pkg.ParseRequestPayload(r, &payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logError(err)
		return
	}

	if err := payload.Validate(); err != nil {
		err := pkg.WrapErr(err.Error(), ValidationFailedErr)
		http.Error(w, err.Error(), http.StatusBadRequest)
		logError(err)
		return
	}

	launch := payload.toEntitiesLaunch()

	if err := slh.Service.DeleteBooking(ctx, launch); err != nil {
		handleDeleteBookingError(w, err)
		logError(err)
		return
	}

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
