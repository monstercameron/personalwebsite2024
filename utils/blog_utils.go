// utils/blog_utils.go

package utils

import (
	"encoding/json"
	"os"
	"time"
)

type BlogPost struct {
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Tags        []string  `json:"tags"`
	Content     string    `json:"content"`
}

type BlogPosts struct {
	Posts []BlogPost `json:"posts"`
}

func LoadBlogPosts() ([]BlogPost, error) {
	file, err := os.ReadFile("static/blog_posts.json")
	if err != nil {
		return nil, err
	}

	var blogPosts BlogPosts
	err = json.Unmarshal(file, &blogPosts)
	if err != nil {
		return nil, err
	}

	return blogPosts.Posts, nil
}