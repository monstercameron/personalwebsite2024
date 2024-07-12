package handlers

import (
	"net/http"
	"personalwebsite/utils"
	"personalwebsite/views"
	"strings"
)

func HandleBlogList(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("HX-Request") == "true" {
		views.BlogContent().Render(r.Context(), w)
	} else {
		views.BlogFullPage().Render(r.Context(), w)
	}
}

func HandleBlogPost(w http.ResponseWriter, r *http.Request) {
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
