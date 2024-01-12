// controllers/user_controller.go
package controllers

import (
	"context"
	"fmt"
	"net/http"

	"Server/db"
	"Server/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/iterator"
)

const usersCollection = "users"

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the user's password
	hashedPassword, err := hashPassword(user.PasswordHash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Set the hashed password
	user.PasswordHash = hashedPassword

	// Omitting the ID field to let Firebase generate a unique ID
	docRef, _, err := db.FirestoreClient.Collection(usersCollection).Add(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Retrieve the generated default UID from the Firestore document reference
	uid := docRef.ID

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "uid": uid})
}

func LoginUser(c *gin.Context) {
	var credentials struct {
	   Email    string `json:"email"`
	   Password string `json:"password"`
	}
 
	if err := c.ShouldBindJSON(&credentials); err != nil {
	   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	   return
	}
 
	// Retrieve user by email
	user, err := getUserByEmail(credentials.Email)
	if err != nil {
	   c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email"})
	   return
	}
 
	// Compare hashed password
	if err := comparePasswords(user.PasswordHash, credentials.Password); err != nil {
		fmt.Println(err)
	   c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
	   return
	}
 
	// Authentication successful
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})
 }
 
 func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
	   return "", err
	}
	return string(hashedPassword), nil
 }
 
 func comparePasswords(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
 }
 
 func getUserByEmail(email string) (*models.User, error) {
	query := db.FirestoreClient.Collection(usersCollection).Where("Email", "==", email)
	iter := query.Documents(context.Background())
 
	doc, err := iter.Next()
	if err == iterator.Done {
	   return nil, err
	}
	if err != nil {
	   return nil, err
	}
 
	var user models.User
	doc.DataTo(&user)
	return &user, nil
 }

func GetUser(c *gin.Context) {
   userID := c.Param("id")

   doc, err := db.FirestoreClient.Collection(usersCollection).Doc(userID).Get(context.Background())
   if err != nil {
      c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
      return
   }

   var user models.User
   doc.DataTo(&user)

   c.JSON(http.StatusOK, user)
}

func GetUsers(c *gin.Context) {
   var users []models.User

   iter := db.FirestoreClient.Collection(usersCollection).Documents(context.Background())
   for {
      doc, err := iter.Next()
      if err == iterator.Done {
         break
      }
      if err != nil {
         c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
         return
      }

      var user models.User
      doc.DataTo(&user)
      users = append(users, user)
   }

   c.JSON(http.StatusOK, users)
}

func UpdateUser(c *gin.Context) {
	userID := c.Param("id")

	var updatedFields map[string]interface{}
	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the user exists
	userRef := db.FirestoreClient.Collection(usersCollection).Doc(userID)
	docSnapshot, err := userRef.Get(context.Background())
	if err != nil {
		if err == iterator.Done {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}

	// Merge the existing user data with the updated fields
	var existingUser models.User
	docSnapshot.DataTo(&existingUser)
	for key, value := range updatedFields {
		switch key {
		case "displayName":
			existingUser.DisplayName = value.(string)
		case "email":
			existingUser.Email = value.(string)
		// Add more cases for other fields you want to update
		}
	}

	// Update the user in Firestore
	_, err = userRef.Set(context.Background(), existingUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUser(c *gin.Context) {
   userID := c.Param("id")

   _, err := db.FirestoreClient.Collection(usersCollection).Doc(userID).Delete(context.Background())
   if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
      return
   }

   c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
