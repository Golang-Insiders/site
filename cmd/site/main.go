package main

import (
	"log"
	"net/http"

	"github.com/golang-insiders/site/internal/data"
)

type application struct {
	cfg      config
	tmpl     Templates
	services data.Services
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatal("Couldn't load config")
	}

	db, err := data.OpenDB(cfg.db)
	if err != nil {
		log.Fatal("Couldn't open db", err)
	}
	defer db.Close()

	services, err := data.NewServices(db)
	if err != nil {
		log.Fatal(err)
	}

	app := application{
		tmpl:     newTemplate(),
		cfg:      cfg,
		services: services,
	}

	http.HandleFunc("/", app.handleHome)
	http.HandleFunc("/talk", app.handleGetTalkByID)
	http.HandleFunc("/submit-form", app.handleTalkPost)
	http.HandleFunc("/submit-form/new", app.handleTalkForm)

	log.Printf("Starting server on %s", cfg.port)
	if err := http.ListenAndServe(cfg.port, nil); err != nil {
		log.Fatal("failed to serve home page.")
	}
}
