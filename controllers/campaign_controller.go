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
	uid := docRef.ID

	// Set the UID in the model
	campaign.UID = uid

	// Update the UID in Firestore
	_, err = docRef.Update(context.Background(), []firestore.Update{
		{Path: "UID", Value: uid},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update UID in Firestore"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Campaign created successfully", "uid": uid})
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

func DeleteCampaign(c *gin.Context) {
	campaignID := c.Param("id")

	_, err := db.FirestoreClient.Collection(campaignsCollection).Doc(campaignID).Delete(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete campaign"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Campaign deleted successfully"})
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

func AddDonorToCampaign(c *gin.Context) {
	campaignID := c.Param("id")

	var donor models.Donor
	if err := c.ShouldBindJSON(&donor); err != nil {
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

	// Add the donor to the campaign
	docSnapshot.Ref.Update(context.Background(), []firestore.Update{
		{Path: "Donors", Value: firestore.ArrayUnion(donor)},
	})

	c.JSON(http.StatusOK, gin.H{"message": "Donor added to the campaign successfully"})
}

func RemoveDonorFromCampaign(c *gin.Context) {
	campaignID := c.Param("id")

	var donor models.Donor
	if err := c.ShouldBindJSON(&donor); err != nil {
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

	// Remove the donor from the campaign
	docSnapshot.Ref.Update(context.Background(), []firestore.Update{
		{Path: "Donors", Value: firestore.ArrayRemove(donor)},
	})

	c.JSON(http.StatusOK, gin.H{"message": "Donor removed from the campaign successfully"})
}
