package handlers

import (
	"net/http"
)

// handleRobotsTxt serves the robots.txt file
func HandleRobotsTxt(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/robots.txt")
}