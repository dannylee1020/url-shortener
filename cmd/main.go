package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

type config struct {
	port int
	env  string
}

type application struct {
	config config
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "environment", "development", "dev|prod")

	app := &application{
		config: cfg,
	}

	app.serve()
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:        fmt.Sprintf(":%d", app.config.port),
		Handler:     app.routes(),
		IdleTimeout: time.Minute,
	}

	err := srv.ListenAndServe()

	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", healthcheckHandler)

	return router
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "ok"}
	jsonResponse, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
