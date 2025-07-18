package main

import (
	"net/http"

	"github.com/RobTov/habibi-salon/internal/store"
)

type clientKey string

const clientCtx clientKey = "client"

type CreateClientPayload struct {
	Name    string `json:"name" validate:"required,max=100"`
	Email   string `json:"email" validate:"required,mail"`
	Phone   string `json:"phone" validate:"required"`
	Address string `json:"address" validate:"required"`
}

func (app *application) getClientHandler(w http.ResponseWriter, r *http.Request) {
	clients, err := app.store.Clients.GetAll(r.Context())
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, clients); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) createClientHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateClientPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	client := &store.Clients{
		Name:    payload.Name,
		Email:   payload.Email,
		Phone:   payload.Phone,
		Address: payload.Address,
	}

	ctx := r.Context()
	if err := app.store.Clients.Create(ctx, client); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) updateClientHandler(w http.ResponseWriter, r *http.Request) {

}
