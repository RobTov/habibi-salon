package main

import (
	"net/http"
)

type CreateClientPayload struct {
	Name    string `json:"name" validate:"required,max=100"`
	Email   string `json:"email" validate:"required,mail"`
	Phone   string `json:"phone" validate:"required"`
	Address string `json:"address" validate:"required"`
}

func (app *application) getClientHandler(w http.ResponseWriter, r *http.Request) {
	// clients, err := app.store.Clients.Ge
}
