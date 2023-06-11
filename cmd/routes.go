package main

import (
	"github.com/dannylee1020/url-shortener/cmd/api"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", api.HealthcheckHandler)

	return router
}
