package handlers

import (
	"net/http"

	"github.com/MyselfRoshan/goHTMX/templates/pages"
)

func HandleGetTodos(w http.ResponseWriter, r *http.Request) {
	pages.Todos("Todos").Render(r.Context(), w)
}
