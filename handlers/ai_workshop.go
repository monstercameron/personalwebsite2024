package handlers

import (
	"fmt"
	"net/http"
	"personalwebsite/views"
)

var (
	vsCodeUser string
	vsCodePass string
)

func SetCredentials(user, pass string) {
	vsCodeUser = user
	vsCodePass = pass
}

func HandleAIWorkshop(w http.ResponseWriter, r *http.Request) {
	loggedIn := false

	if r.Header.Get("HX-Request") == "true" {
		views.AIWorkshopContent(loggedIn).Render(r.Context(), w)
	} else {
		views.AIWorkshopFullPage(loggedIn).Render(r.Context(), w)
	}
}

func HandleAIWorkshopLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == vsCodeUser && password == vsCodePass {
		views.AIWorkshopDashboard().Render(r.Context(), w)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid credentials")
	}
}

func HandleVSCodeSession(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Launching VS Code session... This is where you'd redirect to your actual remote VS Code instance.")
}
