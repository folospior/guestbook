package main

import (
	"net/http"

	"github.com/folospior/guestbook/handlers"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", handlers.HandleHome)
	router.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}
