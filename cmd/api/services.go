package main

import (
	"context"
	"net/http"
	"strconv"

	"github.com/RobTov/habibi-salon/internal/store"
	"github.com/go-chi/chi/v5"
)

type serviceKey string

const serviceCtx serviceKey = "service"

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

func (app *application) getServiceByIDHandler(w http.ResponseWriter, r *http.Request) {
	service := getServiceFromCtx(r)

	if err := app.jsonResponse(w, http.StatusOK, service); err != nil {
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

func (app *application) servicesContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "serviceID")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			app.internalServerError(w, r, err)
			return
		}

		ctx := r.Context()

		service, err := app.store.Services.GetByID(ctx, id)
		if err != nil {
			app.internalServerError(w, r, err)
			return
		}

		ctx = context.WithValue(ctx, serviceCtx, service)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getServiceFromCtx(r *http.Request) *store.Services {
	service, _ := r.Context().Value(serviceCtx).(*store.Services)
	return service
}
