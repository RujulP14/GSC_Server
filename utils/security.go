// utils/password_utils.go
package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates a hashed password using bcrypt.
func HashPassword(password string) (string, error) {
   hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
   if err != nil {
      return "", err
   }
   return string(hashedPassword), nil
}

// ComparePasswords compares a hashed password with a plain-text password.
func ComparePasswords(hashedPassword, password string) error {
   return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
