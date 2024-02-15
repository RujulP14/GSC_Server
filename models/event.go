// models/event.go
package models

type Event struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	Date         string   `json:"date"`
	Day          string   `json:"day"`
	Time         string   `json:"time"`
	Location     string   `json:"location"`
	Description  string   `json:"description"`
	Banner       string   `json:"banner"`
	OrganizerID  string   `json:"organizerId"`
	Participants []string `json:"participants"`
}
