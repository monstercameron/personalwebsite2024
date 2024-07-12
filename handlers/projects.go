package handlers

import (
    "encoding/json"
    "net/http"
    "os"
    "personalwebsite/views"
    "personalwebsite/models"
)

func HandleProjects(w http.ResponseWriter, r *http.Request) {
    // Read the JSON file
    jsonFile, err := os.ReadFile("static/data/projects.json")
    if err != nil {
        http.Error(w, "Error reading projects data", http.StatusInternalServerError)
        return
    }

    // Unmarshal JSON data
    var projects []models.Project
    err = json.Unmarshal(jsonFile, &projects)
    if err != nil {
        http.Error(w, "Error parsing projects data", http.StatusInternalServerError)
        return
    }

    // Render the appropriate template based on the request type
    if r.Header.Get("HX-Request") == "true" {
        views.ProjectsContent(projects).Render(r.Context(), w)
    } else {
        views.ProjectsFullPage(projects).Render(r.Context(), w)
    }
}