// routes/router.go

package routes

import (
	"fmt"
	"net/http"
	"personalwebsite/middleware"
	"personalwebsite/openai"
	"personalwebsite/utils"
	"personalwebsite/views"
	"strings"
	"bytes"
)

var (
	vsCodeUser string
	vsCodePass string
)

type statusRecorder struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
	body        *bytes.Buffer
}

func SetupRoutes(user, pass string) http.Handler {
	vsCodeUser = user
	vsCodePass = pass

	mux := http.NewServeMux()

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Set up specific routes
	mux.HandleFunc("/", middleware.LoggerMiddleware(handleHome))
	mux.HandleFunc("/resume", middleware.LoggerMiddleware(handleResume))
	mux.HandleFunc("/projects", middleware.LoggerMiddleware(handleProjects))
	mux.HandleFunc("/blog", middleware.LoggerMiddleware(handleBlogList))
	mux.HandleFunc("/blog/", middleware.LoggerMiddleware(handleBlogPost))
	mux.HandleFunc("/aiworkshop", middleware.LoggerMiddleware(handleAIWorkshop))
	mux.HandleFunc("/aiworkshop/login", middleware.LoggerMiddleware(handleAIWorkshopLogin))
	mux.HandleFunc("/aiworkshop/vscode-session", middleware.LoggerMiddleware(handleVSCodeSession))
	mux.HandleFunc("/generate-text", middleware.LoggerMiddleware(handleGenerateText))
	mux.HandleFunc("/generate-image", middleware.LoggerMiddleware(handleGenerateImage))
	http.HandleFunc("/metrics", handleMetrics)

	// Create a custom handler to wrap ServeMux and handle 404
	wrappedMux := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK, body: new(bytes.Buffer)}
		mux.ServeHTTP(rec, r)
		if rec.status == http.StatusNotFound {
			handle404(w, r)
		} else {
			rec.ResponseWriter.WriteHeader(rec.status)
			rec.body.WriteTo(rec.ResponseWriter)
		}
	})

	return wrappedMux
}

func (rec *statusRecorder) WriteHeader(status int) {
	if rec.wroteHeader {
		return
	}
	rec.status = status
	rec.wroteHeader = true
}

func (rec *statusRecorder) Write(b []byte) (int, error) {
	if !rec.wroteHeader {
		rec.WriteHeader(http.StatusOK)
	}
	return rec.body.Write(b)
}

func handleMetrics(w http.ResponseWriter, r *http.Request) {
	metricsJSON, err := utils.GetMetricsJSON()
	if err != nil {
		http.Error(w, "Error calculating metrics", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(metricsJSON)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	quote, err := openai.GetQuoteOfTheDay()
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

func handle404(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	views.NotFoundPage().Render(r.Context(), w)
}
