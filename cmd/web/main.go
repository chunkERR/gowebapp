package main

import (
	"fmt"
	"net/http"

	"github.com/chunkERR/gowebapp/pkg/handlers"
)

const portNumber = ":8090"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting app on port %s", portNumber))
	fmt.Println(http.ListenAndServe(portNumber, nil))
}
