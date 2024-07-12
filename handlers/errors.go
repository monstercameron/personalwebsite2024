package handlers

import (
	"net/http"
	"personalwebsite/views"
)

func Handle404(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	views.NotFoundPage().Render(r.Context(), w)
}
