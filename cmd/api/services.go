package main

import (
	"net/http"

	"github.com/RobTov/habibi-salon/internal/store"
)

// type serviceKey string

// const serviceCtx serviceKey = "service"

type CreateServicePayload struct {
	Name        string  `json:"name" validate:"required,max=100"`
	Description string  `json:"description" validate:"required"`
	Price       float32 `json:"price" validate:"required"`
	IsActive    bool    `json:"is_active"`
}

func (app *application) getServiceHandler(w http.ResponseWriter, r *http.Request) {
	services, err := app.store.Services.GetAll(r.Context())
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, services); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) createServiceHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateServicePayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	service := &store.Services{
		Name:        payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
		IsActive:    payload.IsActive,
	}

	ctx := r.Context()
	if err := app.store.Services.Create(ctx, service); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, service); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
