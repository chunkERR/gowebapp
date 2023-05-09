package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chunkERR/gowebapp/pkg/config"
	"github.com/chunkERR/gowebapp/pkg/handlers"
	"github.com/chunkERR/gowebapp/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.MyCache = templateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting app on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
