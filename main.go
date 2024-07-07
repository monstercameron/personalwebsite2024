package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"personalwebsite/openai"
	"personalwebsite/views"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize OpenAI client
	err = openai.Init()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("OpenAI client initialized successfully")

	// Get PORT from .env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port if not specified
	}

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handle routes
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/resume", handleResume)
	http.HandleFunc("/generate-text", handleGenerateText)
	http.HandleFunc("/generate-image", handleGenerateImage)

	// Start the server
	fmt.Printf("Server starting on http://localhost:%s\n", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("HX-Request") == "true" {
		views.HomeContent().Render(r.Context(), w)
	} else {
		views.HomeFullPage().Render(r.Context(), w)
	}
}

func handleResume(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("HX-Request") == "true" {
		views.ResumeContent().Render(r.Context(), w)
	} else {
		views.ResumeFullPage().Render(r.Context(), w)
	}
}

func handleGenerateText(w http.ResponseWriter, r *http.Request) {
	prompt := r.URL.Query().Get("prompt")
	if prompt == "" {
		http.Error(w, "Prompt is required", http.StatusBadRequest)
		return
	}

	completion, err := openai.GenerateCompletion(prompt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, completion)
}

func handleGenerateImage(w http.ResponseWriter, r *http.Request) {
	prompt := r.URL.Query().Get("prompt")
	if prompt == "" {
		http.Error(w, "Prompt is required", http.StatusBadRequest)
		return
	}

	imageURL, err := openai.GenerateImage(prompt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, imageURL)
}