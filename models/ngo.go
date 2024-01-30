// models/ngo.go
package models

type NGOProfile struct {
   NGOName     string `json:"ngoName"`
   WorksFor    string `json:"worksFor"`
   Address     string `json:"address"`
   Pincode     string `json:"pincode"`
   City        string `json:"city"`
   State       string `json:"state"`
   Country     string `json:"country"`
   PhoneNumber string `json:"phoneNumber"`
   Description string `json:"description"`
   Logo        string `json:"logo"`
}

type NGO struct {
   UID          string `json:"uid"`
   Email        string `json:"email"`
   PasswordHash string `json:"password"`
   Profile      NGOProfile `json:"profile"`
}