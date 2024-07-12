// utils/blog_utils.go

package utils

import (
	"encoding/json"
	"os"
	"personalwebsite/models"
)

func LoadBlogPosts() ([]models.BlogPost, error) {
	file, err := os.ReadFile("static/data/blog_posts.json")
	if err != nil {
		return nil, err
	}

	var blogPosts models.BlogPosts
	err = json.Unmarshal(file, &blogPosts)
	if err != nil {
		return nil, err
	}

	return blogPosts.Posts, nil
}
