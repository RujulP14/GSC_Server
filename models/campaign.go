package models

type Donor struct {
	DonorUID      string `json:"donorUid"`
	TransactionID string `json:"transactionId"`
}

type Campaign struct {
	UID         string  `json:"uid"`
	NGO_ID      string  `json:"ngoUid"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ImageURL    string  `json:"imageUrl"`
	RaisedMoney float64 `json:"raisedMoney"`
	TotalGoal   float64 `json:"totalGoal"`
	Donors      []Donor `json:"donors"`
}
