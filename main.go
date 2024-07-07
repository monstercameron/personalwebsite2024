package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"personalwebsite/openai"
	"personalwebsite/views"
	"personalwebsite/utils"
	"strings"

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
	http.HandleFunc("/projects", handleProjects)
	http.HandleFunc("/blog", handleBlogList)
	http.HandleFunc("/blog/", handleBlogPost)
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
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
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

func handleProjects(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("HX-Request") == "true" {
		views.ProjectsContent().Render(r.Context(), w)
	} else {
		views.ProjectsFullPage().Render(r.Context(), w)
	}
}

func handleBlogList(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("HX-Request") == "true" {
		views.BlogContent().Render(r.Context(), w)
	} else {
		views.BlogFullPage().Render(r.Context(), w)
	}
}

func handleBlogPost(w http.ResponseWriter, r *http.Request) {
	slug := strings.TrimPrefix(r.URL.Path, "/blog/")
	posts, err := utils.LoadBlogPosts()
	if err != nil {
		http.Error(w, "Error loading blog posts", http.StatusInternalServerError)
		return
	}

	for _, post := range posts {
		if post.Slug == slug {
			views.BlogPostPage(post).Render(r.Context(), w)
			return
		}
	}

	http.NotFound(w, r)
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
