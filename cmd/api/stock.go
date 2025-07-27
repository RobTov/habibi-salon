package main

import (
	"net/http"

	"github.com/RobTov/habibi-salon/internal/store"
)

type CreateStockPayload struct {
	ProductID int64 `json:"product_id" validate:"required"`
	ServiceID int64 `json:"service_id" validate:"required"`
	Quantity  int64 `json:"quantity" validate:"required"`
}

func (app *application) getStockHandler(w http.ResponseWriter, r *http.Request) {
	stock, err := app.store.Stock.GetAll(r.Context())
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, stock); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) createStockHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateStockPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	stock := &store.Stock{
		ProductID: payload.ProductID,
		ServiceID: payload.ServiceID,
		Quantity:  payload.Quantity,
	}

	ctx := r.Context()
	if err := app.store.Stock.Create(ctx, stock); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, stock); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
