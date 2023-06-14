package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (app *application) serve() error {
	srv := &http.Server{
		Addr:        fmt.Sprintf(":%d", app.config.port),
		Handler:     app.routes(),
		IdleTimeout: time.Minute,
	}

	shutdownError := make(chan error)

	// run separate goroutine to intercept signal for shutting down server
	// and pass it to shutdownError channel
	go func() {
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		s := <-quit

		log.Printf("Shutting down server, signal: %v", s.String())

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		log.Printf("Completing background tasks, addr: %v", srv.Addr)

		// relay any error that might be caused during shutdown
		// if successful no error will be relayed to the channel
		shutdownError <- srv.Shutdown(ctx)
	}()

	log.Printf("Starting server at: %v", app.config.port)

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	// check if there was any error while shutting down the server
	err = <-shutdownError
	if err != nil {
		return err
	}

	log.Printf("Stopped server at port: %v", srv.Addr)

	return nil
}
