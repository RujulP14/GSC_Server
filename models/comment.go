// Comment model
package models

import "time"

type Comment struct {
	CommentID string    `json:"commentID"`
	UserID    string    `json:"userID"`
	Content   string    `json:"content"`
	Commented time.Time `json:"commented"`
}
