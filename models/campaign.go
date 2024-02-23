package models

import "time"

type Donor struct {
	DonorID       string `json:"donorID"`
	TransactionID string `json:"transactionID"`
}

type Campaign struct {
	ID          string    `json:"id"`
	NGO_ID      string    `json:"ngoID"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	EndDate     time.Time `json:"endDate"`
	ImageURL    string    `json:"imageUrl"`
	RaisedMoney float64   `json:"raisedMoney"`
	TotalGoal   float64   `json:"totalGoal"`
	Donors      []Donor   `json:"donors"`
}
