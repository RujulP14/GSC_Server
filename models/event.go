// models/event.go
package models

import "time"

type Event struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Date         time.Time `json:"Date"`
	Location     string    `json:"location"`
	Description  string    `json:"description"`
	Banner       string    `json:"banner"`
	OrganizerID  string    `json:"organizerID"`
	Participants []string  `json:"participants"`
}
