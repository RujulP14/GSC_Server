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
	ID                string      `json:"id"`
	AuthID            string      `json:"authID"`
	Email             string      `json:"email"`
	PasswordHash      string      `json:"password"`
	Profile           UserProfile `json:"profile"`
	Donations         []Donation  `json:"donations"`
	MyEvents          []string    `json:"myEvents"`
	FavouriteArticles []string    `json:"favouriteArticles"`
	FavouriteBlogs    []string    `json:"favouriteBlogs"`
	FavouriteVideos   []string    `json:"favouriteVideos"`
}
