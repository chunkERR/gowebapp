package main

import (
	"fmt"
	"net/http"

	"github.com/chunkERR/gowebapp/pkg/handlers"
)

const portNumber = ":8080"



func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting app on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
