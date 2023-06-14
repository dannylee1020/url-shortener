package main

import (
	"net/http"
)

type envelop map[string]any

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelop{"error": message}

	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		w.WriteHeader(500)
	}
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	message := "The server encountered a problem and could not process your request"
	app.errorResponse(w, r, 500, message)
}

func (app *application) invalidRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}
