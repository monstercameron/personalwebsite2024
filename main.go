package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"personalwebsite/openai"
	"personalwebsite/routes"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

var (
	vsCodeUser string
	vsCodePass string
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

	// Get VS Code credentials from .env
	vsCodeUser = os.Getenv("VSCODEUSER")
	vsCodePass = os.Getenv("VSCODEPASS")

	if vsCodeUser == "" || vsCodePass == "" {
		log.Fatal("VSCODEUSER and VSCODEPASS must be set in .env file")
	}

	fmt.Println("OpenAI client and VS Code credentials initialized successfully")

	// Get PORT from .env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port if not specified
	}

	// Setup routes
	routes.SetupRoutes(vsCodeUser, vsCodePass)

	// Create server
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: nil, // use default ServeMux
	}

	// Start server in a goroutine
	go func() {
		fmt.Printf("Server starting on http://localhost:%s\n", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait until the timeout deadline
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}