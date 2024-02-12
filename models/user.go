// models/user.go
package models

type Donation struct {
	TransactionID string `json:"transactionID"`
	CampaignID    string `json:"campaignID"`
}

type UserProfile struct {
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	PhoneNumber  string `json:"phoneNumber"`
	Dob          string `json:"dob"`
	ProfileImage string `json:"profileImage"`
}

type User struct {
	ID           string      `json:"id"`
	Email        string      `json:"email"`
	PasswordHash string      `json:"password"`
	Profile      UserProfile `json:"profile"`
	Donations    []Donation  `json:"donations"`
}
