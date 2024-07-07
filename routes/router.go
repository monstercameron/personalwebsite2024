// routes/router.go

package routes

import (
	"fmt"
	"net/http"
	"personalwebsite/openai"
	"personalwebsite/utils"
	"personalwebsite/views"
	"strings"
)

var (
	vsCodeUser string
	vsCodePass string
)

func SetupRoutes(user, pass string) {
	vsCodeUser = user
	vsCodePass = pass

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handle routes
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/resume", handleResume)
	http.HandleFunc("/projects", handleProjects)
	http.HandleFunc("/blog", handleBlogList)
	http.HandleFunc("/blog/", handleBlogPost)
	http.HandleFunc("/aiworkshop", handleAIWorkshop)
	http.HandleFunc("/aiworkshop/login", handleAIWorkshopLogin)
	http.HandleFunc("/aiworkshop/vscode-session", handleVSCodeSession)
	http.HandleFunc("/generate-text", handleGenerateText)
	http.HandleFunc("/generate-image", handleGenerateImage)
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

func handleAIWorkshop(w http.ResponseWriter, r *http.Request) {
	// For simplicity, we're assuming the user is not logged in initially
	loggedIn := false

	if r.Header.Get("HX-Request") == "true" {
		views.AIWorkshopContent(loggedIn).Render(r.Context(), w)
	} else {
		views.AIWorkshopFullPage(loggedIn).Render(r.Context(), w)
	}
}

func handleAIWorkshopLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == vsCodeUser && password == vsCodePass {
		// Login successful
		views.AIWorkshopDashboard().Render(r.Context(), w)
	} else {
		// Login failed
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid credentials")
	}
}

func handleVSCodeSession(w http.ResponseWriter, r *http.Request) {
	// In a real application, you'd want to implement some form of authentication check here
	// For simplicity, we're assuming if they can access this route, they're authenticated
	fmt.Fprintf(w, "Launching VS Code session... This is where you'd redirect to your actual remote VS Code instance.")
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
