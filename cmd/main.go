package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"

	"github.com/dannylee1020/url-shortener/internal/data"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn         string
		maxIdleTime string
	}
	limiter bool
}

type application struct {
	config config
	models data.Models
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "environment", "development", "dev|prod")
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("DB_DSN"), "db connection string")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "db max idle time")
	flag.BoolVar(&cfg.limiter, "rate-limiter", true, "turn on and off rate limiter")

	// connect to database
	db, err := openDB(cfg)
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	defer db.Close()

	app := &application{
		config: cfg,
		models: data.NewModels(db),
	}

	app.serve()
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
