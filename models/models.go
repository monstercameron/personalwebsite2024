package models

import (
	"time"
)

type Project struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	TechStack   []string `json:"techStack"`
	GitHubLink  string   `json:"githubLink"`
	DemoLink    string   `json:"demoLink"`
}

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


