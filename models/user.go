// models/user.go
package models

type User struct {
   UID          string `json:"uid"`
   DisplayName  string `json:"displayName"`
   Email        string `json:"email"`
   PasswordHash string `json:"password"`
}
