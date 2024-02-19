// Article model
package models

import "time"

type Article struct {
	ID           string    `firestore:"id"`
	Title        string    `firestore:"title"`
	AuthorName   string    `firestore:"author_name"`
	UploadDate   time.Time `firestore:"upload_date"`
	Content      string    `firestore:"content"`
	Tags         []string  `firestore:"tags"`
	Category     string    `firestore:"category"`
	ThumbnailURL string    `firestore:"thumbnail_url"`
	Comments     []Comment `firestore:"comments"`
	Likes        int       `firestore:"likes"`
}

// Comment model
type Comment struct {
	UserID    string    `firestore:"user_id"`
	Content   string    `firestore:"content"`
	Commented time.Time `firestore:"commented"`
}
