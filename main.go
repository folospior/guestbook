package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/folospior/guestbook/views/components"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", templ.Handler(components.Index()))
	router.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}
