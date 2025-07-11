package main

import "net/http"

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
