package main

import (
	"flag"
	"log"
	"os"
	"text/template"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	api  string
	db   struct {
		dsn string
	}
	stripe struct {
		secret string
		key    string
	}
}

type application struct {
	config        config
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
	version       string
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment {development|production}")
	flag.StringVar(&cfg.env, "api", "http://localhost:4001", "URL to API")

	flag.Parse()

	cfg.stripe.key = os.Getenv("STRIPE_KEY")
	cfg.stripe.secret = os.Getenv("STRIPE_SECRET")
}
