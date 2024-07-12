package routes

import (
	"bytes"
	"net/http"
	"personalwebsite/handlers"
	"personalwebsite/middleware"
)

// SetupRoutes initializes the routes for the web application
func SetupRoutes(user, pass string) http.Handler {
	// Set credentials for handlers
	handlers.SetCredentials(user, pass)

	// Initialize ServeMux
	mux := http.NewServeMux()

	// Serve static files from the "static" directory
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", handlers.CacheControlHandler(fs)))

	// Add handler for robots.txt
	mux.HandleFunc("/robots.txt", handlers.HandleRobotsTxt)

	// Public routes
	mux.HandleFunc("/", middleware.LoggerMiddleware(handlers.HandleHome))
	mux.HandleFunc("/resume", middleware.LoggerMiddleware(handlers.HandleResume))
	mux.HandleFunc("/projects", middleware.LoggerMiddleware(handlers.HandleProjects))
	mux.HandleFunc("/blog", middleware.LoggerMiddleware(handlers.HandleBlogList))
	mux.HandleFunc("/blog/", middleware.LoggerMiddleware(handlers.HandleBlogPost))

	// AI Workshop routes
	mux.HandleFunc("/aiworkshop", middleware.LoggerMiddleware(handlers.HandleAIWorkshop))
	mux.HandleFunc("/aiworkshop/login", middleware.LoggerMiddleware(handlers.HandleAIWorkshopLogin))
	mux.HandleFunc("/aiworkshop/vscode-session", middleware.LoggerMiddleware(handlers.HandleVSCodeSession))

	// Generate content routes
	mux.HandleFunc("/generate-text", middleware.LoggerMiddleware(handlers.HandleGenerateText))
	mux.HandleFunc("/generate-image", middleware.LoggerMiddleware(handlers.HandleGenerateImage))

	// Metrics route
	mux.HandleFunc("/metrics", handlers.HandleMetrics)

	// Wrap the mux to handle 404 errors
	wrappedMux := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK, body: new(bytes.Buffer)}
		mux.ServeHTTP(rec, r)
		if rec.status == http.StatusNotFound {
			handlers.Handle404(w, r)
		} else {
			rec.ResponseWriter.WriteHeader(rec.status)
			rec.body.WriteTo(rec.ResponseWriter)
		}
	})

	return middleware.GzipMiddleware(wrappedMux) // Apply gzip middleware to all routes
	// return wrappedMux
}
