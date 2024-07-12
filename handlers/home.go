package handlers

import (
	"fmt"
	"net/http"
	"personalwebsite/utils"
	"personalwebsite/views"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	quote, err := utils.GetQuoteOfTheDay()
	if err != nil {
		fmt.Printf("Error getting quote of the day: %v\n", err)
		quote = "Probably out of OAI tokens"
	}

	if r.Header.Get("HX-Request") == "true" {
		views.HomeContent(quote).Render(r.Context(), w)
	} else {
		views.HomeFullPage(quote).Render(r.Context(), w)
	}
}
