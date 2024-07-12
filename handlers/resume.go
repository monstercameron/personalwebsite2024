package handlers

import (
    "net/http"
    "personalwebsite/views"
)

func HandleResume(w http.ResponseWriter, r *http.Request) {
    if r.Header.Get("HX-Request") == "true" {
        views.ResumeContent().Render(r.Context(), w)
    } else {
        views.ResumeFullPage().Render(r.Context(), w)
    }
}