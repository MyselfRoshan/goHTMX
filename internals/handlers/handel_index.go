package handlers

import (
	"net/http"

	"github.com/MyselfRoshan/goHTMX/templates/pages"
)

func HandleGetIndex(w http.ResponseWriter, r *http.Request) {
	pages.Index("This is index page", nil).Render(r.Context(), w)
}

func HandlePostIndex(w http.ResponseWriter, r *http.Request) {
	pages.NewIndex("This is new site").Render(r.Context(), w)
}
