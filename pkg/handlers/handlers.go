package handlers

import (
	"net/http"

	"github.com/chunkERR/gowebapp/pkg/render"
)

// Home is the home page handler

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "homepage.html")
}

// About is the home page handler

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}

