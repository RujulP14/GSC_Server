// controllers/campaign_controller.go
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

const campaignsCollection = "campaigns"

func CreateCampaign(c *gin.Context) {
	var campaign models.Campaign
	if err := c.ShouldBindJSON(&campaign); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	docRef, _, err := db.FirestoreClient.Collection(campaignsCollection).Add(context.Background(), campaign)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create campaign"})
		return
	}

	// Retrieve the generated UID from the Firestore document reference
	campaignID := docRef.ID

	// Set the campaign ID in the model
	campaign.ID = campaignID

	// Update the UID in Firestore
	_, err = docRef.Update(context.Background(), []firestore.Update{
		{Path: "ID", Value: campaignID},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update ID in Firestore"})
		return
	}

	// Add the campaign ID to the NGO's campaigns array
	ngoID := campaign.NGO_ID
	ngoRef := db.FirestoreClient.Collection(ngosCollection).Doc(ngoID)
	err = db.FirestoreClient.RunTransaction(context.Background(), func(ctx context.Context, tx *firestore.Transaction) error {
		// Get the current campaigns array of the NGO
		docSnap, err := tx.Get(ngoRef)
		if err != nil {
			return err
		}

		var ngo models.NGO
		if err := docSnap.DataTo(&ngo); err != nil {
			return err
		}

		// Append the new campaign ID to the array
		ngo.Campaigns = append(ngo.Campaigns, campaignID)

		// Update the NGO document in Firestore
		if err := tx.Set(ngoRef, ngo); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update NGO with campaign ID"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Campaign created successfully", "id": campaignID})
}

func GetCampaign(c *gin.Context) {
	campaignID := c.Param("id")

	doc, err := db.FirestoreClient.Collection(campaignsCollection).Doc(campaignID).Get(context.Background())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Campaign not found"})
		return
	}

	var campaign models.Campaign
	doc.DataTo(&campaign)

	c.JSON(http.StatusOK, campaign)
}

func GetCampaigns(c *gin.Context) {
	var campaigns []models.Campaign

	iter := db.FirestoreClient.Collection(campaignsCollection).Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve campaigns"})
			return
		}

		var campaign models.Campaign
		doc.DataTo(&campaign)
		campaigns = append(campaigns, campaign)
	}

	c.JSON(http.StatusOK, campaigns)
}

func UpdateCampaign(c *gin.Context) {
	campaignID := c.Param("id")

	var updatedFields map[string]interface{}
	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the campaign exists
	campaignRef := db.FirestoreClient.Collection(campaignsCollection).Doc(campaignID)
	docSnapshot, err := campaignRef.Get(context.Background())
	if err != nil {
		if err == iterator.Done {
			c.JSON(http.StatusNotFound, gin.H{"error": "Campaign not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve campaign"})
		return
	}

	// Merge the existing campaign data with the updated fields
	var existingCampaign models.Campaign
	docSnapshot.DataTo(&existingCampaign)
	for key, value := range updatedFields {
		switch key {
		case "title":
			existingCampaign.Title = value.(string)
		case "description":
			existingCampaign.Description = value.(string)
			// Add more cases for other fields you want to update
		}
	}

	// Update the campaign in Firestore
	_, err = campaignRef.Set(context.Background(), existingCampaign)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update campaign"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Campaign updated successfully"})
}

func UpdateCampaignImage(c *gin.Context) {
	campaignID := c.Param("id")

	var updatedFields map[string]interface{}
	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the campaign exists
	campaignRef := db.FirestoreClient.Collection(campaignsCollection).Doc(campaignID)
	docSnapshot, err := campaignRef.Get(context.Background())
	if err != nil {
		if err == iterator.Done {
			c.JSON(http.StatusNotFound, gin.H{"error": "Campaign not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve campaign"})
		return
	}

	// Update the image URL if provided in the request
	if imageURL, ok := updatedFields["ImageURL"]; ok {
		docSnapshot.Ref.Update(context.Background(), []firestore.Update{
			{Path: "ImageURL", Value: imageURL},
		})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Campaign image updated successfully"})
}

func DeleteCampaign(c *gin.Context) {
	campaignID := c.Param("id")

	// Retrieve the campaign document
	campaignRef := db.FirestoreClient.Collection(campaignsCollection).Doc(campaignID)
	campaignDoc, err := campaignRef.Get(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve campaign"})
		return
	}

	// Get the NGO ID associated with the campaign
	var campaign models.Campaign
	if err := campaignDoc.DataTo(&campaign); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse campaign data"})
		return
	}
	ngoID := campaign.NGO_ID

	// Delete the campaign document
	_, err = campaignRef.Delete(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete campaign"})
		return
	}

	// Remove the campaign ID from the NGO's campaigns array
	ngoRef := db.FirestoreClient.Collection(ngosCollection).Doc(ngoID)
	err = db.FirestoreClient.RunTransaction(context.Background(), func(ctx context.Context, tx *firestore.Transaction) error {
		// Get the current campaigns array of the NGO
		docSnap, err := tx.Get(ngoRef)
		if err != nil {
			return err
		}

		var ngo models.NGO
		if err := docSnap.DataTo(&ngo); err != nil {
			return err
		}

		// Find and remove the campaign ID from the array
		var updatedCampaigns []string
		for _, iID := range ngo.Campaigns {
			if iID != campaignID {
				updatedCampaigns = append(updatedCampaigns, iID)
			}
		}
		ngo.Campaigns = updatedCampaigns

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

	c.JSON(http.StatusOK, gin.H{"message": "Campaign deleted successfully"})
}
