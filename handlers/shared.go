package handlers_test

import (
	"net/http"

	"github.com/a-h/templ"
)

func Render(w http.ResponseWriter, r *http.Request, c templ.Component) {
	c.Render(r.Context(), w)
}