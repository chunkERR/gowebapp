package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chunkERR/gowebapp/pkg/config"
	"github.com/chunkERR/gowebapp/pkg/handlers"
	"github.com/chunkERR/gowebapp/pkg/render"
)

const portNumber = ":8090"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.MyCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting app on port %s", portNumber))
	fmt.Println(http.ListenAndServe(portNumber, nil))
}
