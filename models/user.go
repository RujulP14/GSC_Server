// models/user.go
package models

type UserProfile struct {
	FirstName    string     `json:"firstName"`
	LastName     string     `json:"lastName"`
	Dob          string     `json:"dob"`
	ProfileImage string     `json:"profileImage"`
	Donations    []Donation `json:"donations"`
}

type Donation struct {
	CampaignID    string `json:"campaignID"`
	TransactionID string `json:"transactionID"`
}

type User struct {
	UID          string      `json:"uid"`
	Email        string      `json:"email"`
	PasswordHash string      `json:"password"`
	Profile      UserProfile `json:"profile"`
}
