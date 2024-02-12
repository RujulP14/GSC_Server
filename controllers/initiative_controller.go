// controllers/initiative_controller.go
package controllers

import (
	"context"
	"net/http"

	"Server/db"
	"Server/models"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

const initiativesCollection = "initiatives"

func CreateInitiative(c *gin.Context) {
	var initiative models.Initiative
	if err := c.ShouldBindJSON(&initiative); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	docRef, _, err := db.FirestoreClient.Collection(initiativesCollection).Add(context.Background(), initiative)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create initiative"})
		return
	}

	// Retrieve the generated UID from the Firestore document reference
	initiativeID := docRef.ID

	// Set the initiative ID in the model
	initiative.ID = initiativeID

	// Update the UID in Firestore
	_, err = docRef.Update(context.Background(), []firestore.Update{
		{Path: "ID", Value: initiativeID},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update ID in Firestore"})
		return
	}

	// Add the initiative ID to the NGO's initiatives array
	ngoID := initiative.NGO_ID
	ngoRef := db.FirestoreClient.Collection(ngosCollection).Doc(ngoID)
	err = db.FirestoreClient.RunTransaction(context.Background(), func(ctx context.Context, tx *firestore.Transaction) error {
		// Get the current initiatives array of the NGO
		docSnap, err := tx.Get(ngoRef)
		if err != nil {
			return err
		}

		var ngo models.NGO
		if err := docSnap.DataTo(&ngo); err != nil {
			return err
		}

		// Append the new initiative ID to the array
		ngo.Initiatives = append(ngo.Initiatives, initiativeID)

		// Update the NGO document in Firestore
		if err := tx.Set(ngoRef, ngo); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update NGO with initiative ID"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Initiative created successfully", "id": initiativeID})
}

func GetInitiative(c *gin.Context) {
	initiativeID := c.Param("id")

	doc, err := db.FirestoreClient.Collection(initiativesCollection).Doc(initiativeID).Get(context.Background())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Initiative not found"})
		return
	}

	var initiative models.Initiative
	doc.DataTo(&initiative)

	c.JSON(http.StatusOK, initiative)
}

func GetInitiatives(c *gin.Context) {
	var initiatives []models.Initiative

	iter := db.FirestoreClient.Collection(initiativesCollection).Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve initiatives"})
			return
		}

		var initiative models.Initiative
		doc.DataTo(&initiative)
		initiatives = append(initiatives, initiative)
	}

	c.JSON(http.StatusOK, initiatives)
}

func UpdateInitiative(c *gin.Context) {
	initiativeID := c.Param("id")

	var updatedFields map[string]interface{}
	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the initiative exists
	initiativeRef := db.FirestoreClient.Collection(initiativesCollection).Doc(initiativeID)
	docSnapshot, err := initiativeRef.Get(context.Background())
	if err != nil {
		if err == iterator.Done {
			c.JSON(http.StatusNotFound, gin.H{"error": "Initiative not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve initiative"})
		return
	}

	// Merge the existing initiative data with the updated fields
	var existingInitiative models.Initiative
	docSnapshot.DataTo(&existingInitiative)
	for key, value := range updatedFields {
		switch key {
		case "title":
			existingInitiative.Title = value.(string)
		case "description":
			existingInitiative.Description = value.(string)
			// Add more cases for other fields you want to update
		}
	}

	// Update the initiative in Firestore
	_, err = initiativeRef.Set(context.Background(), existingInitiative)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update initiative"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Initiative updated successfully"})
}

func UpdateInitiativeImage(c *gin.Context) {
	initiativeID := c.Param("id")

	var updatedFields map[string]interface{}
	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the initiative exists
	initiativeRef := db.FirestoreClient.Collection(initiativesCollection).Doc(initiativeID)
	docSnapshot, err := initiativeRef.Get(context.Background())
	if err != nil {
		if err == iterator.Done {
			c.JSON(http.StatusNotFound, gin.H{"error": "Initiative not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve initiative"})
		return
	}

	// Update the image URL if provided in the request
	if imageURL, ok := updatedFields["ImageURL"]; ok {
		docSnapshot.Ref.Update(context.Background(), []firestore.Update{
			{Path: "ImageURL", Value: imageURL},
		})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Initiative image updated successfully"})
}

func DeleteInitiative(c *gin.Context) {
	initiativeID := c.Param("id")

	// Retrieve the initiative document
	initiativeRef := db.FirestoreClient.Collection(initiativesCollection).Doc(initiativeID)
	initiativeDoc, err := initiativeRef.Get(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve initiative"})
		return
	}

	// Get the NGO ID associated with the initiative
	var initiative models.Initiative
	if err := initiativeDoc.DataTo(&initiative); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse initiative data"})
		return
	}
	ngoID := initiative.NGO_ID

	// Delete the initiative document
	_, err = initiativeRef.Delete(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete initiative"})
		return
	}

	// Remove the initiative ID from the NGO's initiatives array
	ngoRef := db.FirestoreClient.Collection(ngosCollection).Doc(ngoID)
	err = db.FirestoreClient.RunTransaction(context.Background(), func(ctx context.Context, tx *firestore.Transaction) error {
		// Get the current initiatives array of the NGO
		docSnap, err := tx.Get(ngoRef)
		if err != nil {
			return err
		}

		var ngo models.NGO
		if err := docSnap.DataTo(&ngo); err != nil {
			return err
		}

		// Find and remove the initiative ID from the array
		var updatedInitiatives []string
		for _, iID := range ngo.Initiatives {
			if iID != initiativeID {
				updatedInitiatives = append(updatedInitiatives, iID)
			}
		}
		ngo.Initiatives = updatedInitiatives

		// Update the NGO document in Firestore
		if err := tx.Set(ngoRef, ngo); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update NGO"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Initiative deleted successfully"})
}
