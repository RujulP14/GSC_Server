// controllers/ngo_controller.go
package controllers

import (
	"context"
	"net/http"

	"Server/db"
	"Server/models"
	"Server/utils"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

const ngosCollection = "ngos"

func CreateNGO(c *gin.Context) {
   var ngo models.NGO
   if err := c.ShouldBindJSON(&ngo); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
   }

   // Hash the NGO's password
   hashedPassword, err := utils.HashPassword(ngo.PasswordHash)
   if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
      return
   }

   // Set the hashed password
   ngo.PasswordHash = hashedPassword

   // Omitting the ID field to let Firebase generate a unique ID
   docRef, _, err := db.FirestoreClient.Collection(ngosCollection).Add(context.Background(), ngo)
   if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create NGO"})
      return
   }

   // Retrieve the generated default UID from the Firestore document reference
   // Retrieve the generated default UID from the Firestore document reference
   uid := docRef.ID

   // Set the UID in the model
   ngo.UID = uid

   // Update the UID in Firestore
	_, err = docRef.Update(context.Background(), []firestore.Update{
		{Path: "UID", Value: uid},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update UID in Firestore"})
		return
	}
   
   c.JSON(http.StatusCreated, gin.H{"message": "NGO created successfully", "uid": uid})
}


func UpdateNGO(c *gin.Context) {
	ngoID := c.Param("id")

	var updatedNGO models.NGO
	if err := c.ShouldBindJSON(&updatedNGO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the NGO exists
	ngoRef := db.FirestoreClient.Collection(ngosCollection).Doc(ngoID)
	docSnapshot, err := ngoRef.Get(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve NGO"})
		return
	}

	// Update the entire NGOProfile with the content from the request
	var existingNGO models.NGO
	docSnapshot.DataTo(&existingNGO)
	existingNGO.Profile = updatedNGO.Profile

	// Update the NGO in Firestore
	_, err = ngoRef.Set(context.Background(), existingNGO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update NGO"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "NGO updated successfully"})
}


func GetNGO(c *gin.Context) {
	ngoID := c.Param("id")

	ngoRef := db.FirestoreClient.Collection(ngosCollection).Doc(ngoID)
	docSnapshot, err := ngoRef.Get(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve NGO"})
		return
	}

	var ngo models.NGO
	docSnapshot.DataTo(&ngo)

	c.JSON(http.StatusOK, ngo)
}

func GetNGOs(c *gin.Context) {
	ngos := []models.NGO{}

	// Retrieve all NGOs from Firestore
	iter := db.FirestoreClient.Collection(ngosCollection).Documents(context.Background())
	for {
		docSnapshot, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve NGOs"})
			return
		}

		var ngo models.NGO
		docSnapshot.DataTo(&ngo)
		ngos = append(ngos, ngo)
	}

	c.JSON(http.StatusOK, ngos)
}

func DeleteNGO(c *gin.Context) {
	ngoID := c.Param("id")

	_, err := db.FirestoreClient.Collection(ngosCollection).Doc(ngoID).Delete(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete NGO"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "NGO deleted successfully"})
}

func LoginNGO(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve NGO by email
	ngo, err := GetNGOByEmail(credentials.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email"})
		return
	}

	// Compare hashed password
	if err := utils.ComparePasswords(ngo.PasswordHash, credentials.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
		return
	}

	// Authentication successful
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "ngo": ngo})
}

func GetNGOByEmail(email string) (models.NGO, error) {
	var ngo models.NGO

	// Retrieve the NGO by email from Firestore
	query := db.FirestoreClient.Collection(ngosCollection).Where("Email", "==", email)
	iter := query.Documents(context.Background())
	for {
		docSnapshot, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return models.NGO{}, err
		}

		if err := docSnapshot.DataTo(&ngo); err != nil {
			return models.NGO{}, err
		}
	}

	return ngo, nil
}
