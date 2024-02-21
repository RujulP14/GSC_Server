// models/blog.go
package models

import "time"

type Blog struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Author   string    `json:"author"`
	Image    string    `json:"image"`
	Category string    `json:"category"`
	Date     time.Time `json:"date"`
}
