package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const portNumber = ":8080"

// Home is the home page handler

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "homepage.html")
}

// About is the home page handler

func About(w http.ResponseWriter, r *http.Request) {

}

func renderTemplate(w http.ResponseWriter, html string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + html)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting app on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
