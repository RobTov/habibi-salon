package main

import "net/http"

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
