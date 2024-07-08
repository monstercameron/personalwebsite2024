// routes/router.go

package routes

import (
	"fmt"
	"strings"
	"net/http"
	"personalwebsite/openai"
	"personalwebsite/utils"
	"personalwebsite/views"
    "personalwebsite/middleware"
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

    // Handle routes with logger middleware
    http.HandleFunc("/", middleware.LoggerMiddleware(handleHome))
    http.HandleFunc("/resume", middleware.LoggerMiddleware(handleResume))
    http.HandleFunc("/projects", middleware.LoggerMiddleware(handleProjects))
    http.HandleFunc("/blog", middleware.LoggerMiddleware(handleBlogList))
    http.HandleFunc("/blog/", middleware.LoggerMiddleware(handleBlogPost))
    http.HandleFunc("/aiworkshop", middleware.LoggerMiddleware(handleAIWorkshop))
    http.HandleFunc("/aiworkshop/login", middleware.LoggerMiddleware(handleAIWorkshopLogin))
    http.HandleFunc("/aiworkshop/vscode-session", middleware.LoggerMiddleware(handleVSCodeSession))
    http.HandleFunc("/generate-text", middleware.LoggerMiddleware(handleGenerateText))
    http.HandleFunc("/generate-image", middleware.LoggerMiddleware(handleGenerateImage))
	http.HandleFunc("/metrics", handleMetrics)
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
