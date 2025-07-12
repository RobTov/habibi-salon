package main

import (
	"net/http"
	"time"

	"github.com/RobTov/habibi-salon/internal/store"
)

type CreateAppointmentPayload struct {
	Date      time.Time `json:"date" validate:"required"`
	ClientID  int64     `json:"client_id" validate:"required"`
	ServiceID int64     `json:"service_id" validate:"required"`
	Status    string    `json:"status" validate:"required"`
}

func (app *application) getAppointmentHandler(w http.ResponseWriter, r *http.Request) {
	appointments, err := app.store.Appointments.GetAll(r.Context())
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, appointments); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) createAppointmentHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateAppointmentPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	appointment := &store.Appointments{
		Date:      payload.Date,
		ClientID:  payload.ClientID,
		ServiceID: payload.ServiceID,
		Status:    payload.Status,
	}

	ctx := r.Context()
	if err := app.store.Appointments.Create(ctx, appointment); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, appointment); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
