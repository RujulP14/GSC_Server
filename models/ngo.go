// models/ngo.go
package models

type NGOProfile struct {
	RegistrationNumber string `json:"registrationNumber"`
	NGOName            string `json:"ngoName"`
	WorksFor           string `json:"worksFor"`
	Address            string `json:"address"`
	Pincode            string `json:"pincode"`
	City               string `json:"city"`
	State              string `json:"state"`
	Country            string `json:"country"`
	PhoneNumber        string `json:"phoneNumber"`
	Description        string `json:"description"`
	Logo               string `json:"logo"`
}

type NGO struct {
	ID           string     `json:"id"`
	Email        string     `json:"email"`
	PasswordHash string     `json:"password"`
	Profile      NGOProfile `json:"profile"`
	Campaigns    []string   `json:"campaigns"`
	Events       []string   `json:"events"`
}
