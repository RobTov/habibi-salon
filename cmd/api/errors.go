package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Internal Server error: %s path: %s. Error: %s\n", r.Method, r.URL.Path, err)

	writeJSONError(w, http.StatusInternalServerError, "The server encountered a problem.")
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Bad Request error: %s path: %s. Error: %s\n", r.Method, r.URL.Path, err)

	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Not Found error: %s path: %s. Error: %s\n", r.Method, r.URL.Path, err)

	writeJSONError(w, http.StatusNotFound, "not found")
}
