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

type SpaceLauncherHandler struct {
}

func (slh SpaceLauncherHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func (slh SpaceLauncherHandler) GetBookings(w http.ResponseWriter, r *http.Request) {
	log.Println("GetBookings")

	if err := WriteJSON(w, http.StatusOK, nil); err != nil {
		http.Error(w, fmt.Sprintf("output serialization failed, %v", err), http.StatusInternalServerError)
	}
}

func (slh SpaceLauncherHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateBooking")

	if err := WriteJSON(w, http.StatusNoContent, nil); err != nil {
		http.Error(w, fmt.Sprintf("output serialization failed, %v", err), http.StatusInternalServerError)
	}
}

func (slh SpaceLauncherHandler) DeleteBooking(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteBooking")

	if err := WriteJSON(w, http.StatusNoContent, nil); err != nil {
		http.Error(w, fmt.Sprintf("output serialization failed, %v", err), http.StatusInternalServerError)
	}
}

// WriteJSON with http.ResponseWriter.
func WriteJSON(w http.ResponseWriter, status int, response interface{}) error {
	w.Header().Set(ContentTypeKey, ContentTypeValueJSON)
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(response)
}
