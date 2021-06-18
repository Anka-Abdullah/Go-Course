package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Anka-Abdullah/Go-Course/pkg/config"
	"github.com/Anka-Abdullah/Go-Course/pkg/handlers"
	"github.com/Anka-Abdullah/Go-Course/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("server running on port %s", portNumber)

	_ = http.ListenAndServe(portNumber, nil)
}

//command to run go run cmd/web/*.go
