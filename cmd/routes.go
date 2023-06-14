package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.HealthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/shorten", app.ShortenURL)
	router.Handle(http.MethodGet, "/v1/redirect/:url", app.RedirectHandler)

	return router
}
