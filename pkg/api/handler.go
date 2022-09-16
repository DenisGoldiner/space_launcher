package api

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type SpaceLauncherHandler struct {
}

func (slh SpaceLauncherHandler) Extend(r chi.Router) {
	r.Route(
		"/bookings", func(r chi.Router) {
			r.Get("/", slh.GetBookings)
			r.Post("/", slh.CreateBooking)
			r.Delete("/", slh.DeleteBooking)
		},
	)
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
}

func (slh SpaceLauncherHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateBooking")
}

func (slh SpaceLauncherHandler) DeleteBooking(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteBooking")
}
