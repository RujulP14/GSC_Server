// Article model
package models

import "time"

type Article struct {
	ID                string    `json:"id"`
	Title             string    `json:"title"`
	AuthorName        string    `json:"authorName"`
	UploadDate        time.Time `json:"uploadDate"`
	Content           string    `json:"content"`
	EstimatedReadTime int       `json:"estimatedReadTime"`
	Tags              []string  `json:"tags"`
	Category          string    `json:"category"`
	ThumbnailURL      string    `json:"thumbnailUrl"`
	Comments          []Comment `json:"comments"`
	Likes             int       `json:"likes"`
}
