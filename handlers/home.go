package handlers

import (
	"net/http"

	"github.com/folospior/guestbook/views/components"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	Render(w, r, components.Index())
}