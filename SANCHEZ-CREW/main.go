package main

import (
	"fmt"
	"log"
	"net/http"

	"SANCHEZ-CREW/internal/templates"

	"github.com/a-h/templ"
)

func main() {

	landing := templates.LandingPage()

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/", templ.Handler(landing))

	fmt.Println("Listening on :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Printf("Error listening: %v", err)
	}
}
