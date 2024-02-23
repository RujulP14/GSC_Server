// models/blog.go
package models

import "time"

type Blog struct {
	ID                string    `json:"id"`
	Title             string    `json:"title"`
	Content           string    `json:"content"`
	EstimatedReadTime int       `json:"estimatedReadTime"`
	AuthorID          string    `json:"authorID"`
	AuthorName        string    `json:"authorName"`
	Image             string    `json:"image"`
	Category          string    `json:"category"`
	Date              time.Time `json:"date"`
}
