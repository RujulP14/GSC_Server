// Article model
package models

import "time"

type Article struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	AuthorName   string    `json:"authorName"`
	UploadDate   time.Time `json:"uploadDate"`
	Content      string    `json:"content"`
	Tags         []string  `json:"tags"`
	Category     string    `json:"category"`
	ThumbnailURL string    `json:"thumbnailUrl"`
	Comments     []Comment `json:"comments"`
	Likes        int       `json:"likes"`
}

// Comment model
type Comment struct {
	CommentID string    `json:"commentID"`
	UserID    string    `json:"userID"`
	Content   string    `json:"content"`
	Commented time.Time `json:"commented"`
}
