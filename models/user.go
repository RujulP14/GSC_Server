// models/user.go
package models

type User struct {

   DisplayName  string `json:"displayName"`
   Email        string `json:"email"`
   PasswordHash string `json:"password"`
}
