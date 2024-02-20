package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	webPort := os.Getenv("PORT")
	cfg.port, _ = strconv.Atoi(webPort)
	if cfg.port == 0 {
		cfg.port = 4000
	}

	env := os.Getenv("ENV")
	if env == "" {
		cfg.env = "development"
	}

	app := application{
		config: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)

}
