package main

import (
	"fmt"
	"net/http"

	"github.com/Anka-Abdullah/Go-Course/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("server running on port %s", portNumber)

	http.ListenAndServe(portNumber, nil)
}

//command to run go run *.go
