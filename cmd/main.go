package main

import (
	"flag"
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
