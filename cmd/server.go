package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

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
