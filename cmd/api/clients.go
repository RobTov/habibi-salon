package main

import (
	"context"
	"net/http"
	"strconv"

	"github.com/RobTov/habibi-salon/internal/store"
	"github.com/go-chi/chi/v5"
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

func (app *application) getClientByIDHandler(w http.ResponseWriter, r *http.Request) {
	client := getClientFromCtx(r)

	if err := app.jsonResponse(w, http.StatusOK, client); err != nil {
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

func (app *application) deleteClientHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "clientID")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	err = app.store.Clients.Delete(r.Context(), id)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (app *application) clientsContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "clientID")
		id, err := strconv.ParseInt(idParam, 10, 64)

		if err != nil {
			app.internalServerError(w, r, err)
			return
		}

		ctx := r.Context()

		client, err := app.store.Clients.GetByID(ctx, id)
		if err != nil {
			app.internalServerError(w, r, err)
			return
		}

		ctx = context.WithValue(ctx, clientCtx, client)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getClientFromCtx(r *http.Request) *store.Clients {
	client, _ := r.Context().Value(clientCtx).(*store.Clients)
	return client
}
