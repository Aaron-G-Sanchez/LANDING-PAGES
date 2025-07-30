package main

import (
	"fmt"
	"log"
	"net/http"

	"SANCHEZ-CREW/internal/templates"

	"github.com/a-h/templ"
)

func main() {
	fmt.Println("Hello")

	landing := templates.LandingPage()

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/", templ.Handler(landing))

	fmt.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Printf("Error listening: %v", err)
	}
}
