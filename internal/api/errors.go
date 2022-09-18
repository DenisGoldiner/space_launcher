package api

import (
	"errors"
	"net/http"

	"github.com/DenisGoldiner/space_launcher/internal/service"
)

var (
	EmptyFieldErr       = errors.New("field should not be empty")
	ValidationFailedErr = errors.New("the request validation failed")
)

func handleCreateBookingError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, service.BusinessValidationErr):
		http.Error(w, err.Error(), http.StatusBadRequest)
	case errors.Is(err, service.ExternalVendorAPIErr):
		http.Error(w, err.Error(), http.StatusInternalServerError)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleDeleteBookingError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, service.BookingNotFoundErr):
		http.Error(w, err.Error(), http.StatusNotFound)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
