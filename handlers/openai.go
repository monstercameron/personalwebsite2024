package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"personalwebsite/utils"
)

func HandleGenerateText(w http.ResponseWriter, r *http.Request) {
	prompt := r.URL.Query().Get("prompt")
	if prompt == "" {
		http.Error(w, "Prompt is required", http.StatusBadRequest)
		return
	}

	completion, err := utils.GenerateCompletion(prompt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, completion)
}

func HandleGenerateImage(w http.ResponseWriter, r *http.Request) {
	prompt := r.URL.Query().Get("prompt")
	if prompt == "" {
		http.Error(w, "Prompt is required", http.StatusBadRequest)
		return
	}

	imageURL, err := utils.GenerateImage(prompt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := http.Get(imageURL)
	if err != nil {
		http.Error(w, "Failed to download image", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	staticDir := filepath.Join("static", "images")
	if err := os.MkdirAll(staticDir, os.ModePerm); err != nil {
		http.Error(w, "Failed to create directory", http.StatusInternalServerError)
		return
	}

	filePath := filepath.Join(staticDir, "qotd.jpg")
	out, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Failed to create file", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		http.Error(w, "Failed to save image", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "/static/images/qotd.jpg")
}
