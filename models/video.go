// models/video.go
package models

import "time"

type Video struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	VideoURL     string    `json:"videoUrl"`
	Uploader     string    `json:"uploader"`
	UploadDate   time.Time `json:"uploadDate"`
	Description  string    `json:"description"`
	Tags         []string  `json:"tags"`
	Category     string    `json:"category"`
	ThumbnailURL string    `json:"thumbnailUrl"`
	Comments     []Comment `json:"comments"`
	Likes        int       `json:"likes"`
	Transcripts  string    `json:"transcripts"`
}
