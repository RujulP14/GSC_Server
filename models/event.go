// models/event.go
package models

import "time"

type Event struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	UploadDate   time.Time `firestore:"upload_date"`
	Location     string    `json:"location"`
	Description  string    `json:"description"`
	Banner       string    `json:"banner"`
	OrganizerID  string    `json:"organizerId"`
	Participants []string  `json:"participants"`
}
