package controllers

import (
	"context"
	"net/http"
	"time"

	"Server/db"
	"Server/models"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

// HandleDonation handles donations by storing transaction data in Firestore
func HandleDonation(c *gin.Context) {
	// Parse request body to get transaction data
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set current time as transaction time
	transaction.TransactionTime = time.Now()

	// Add transaction to Firestore and retrieve the auto-generated ID
	docRef, _, err := db.FirestoreClient.Collection("transactions").Add(context.Background(), transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store transaction data"})
		return
	}

	// Update the transaction ID with the auto-generated ID from Firestore
	transactionID := docRef.ID
	transaction.TransactionID = transactionID

	// Update the transaction document in Firestore with the new ID
	if _, err := docRef.Set(context.Background(), transaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update transaction with ID"})
		return
	}

	// Update the campaign's donors with the transaction and sender ID
	campaignID := transaction.CampaignID
	donor := models.Donor{
		DonorID:       transaction.SenderID,
		TransactionID: transactionID,
	}

	_, err = db.FirestoreClient.Collection("campaigns").Doc(campaignID).Update(context.Background(), []firestore.Update{
		{Path: "Donors", Value: firestore.ArrayUnion(donor)},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update campaign donors"})
		return
	}

	// Update the user's donations with the transaction and campaign ID
	userID := transaction.SenderID
	donation := models.Donation{
		TransactionID: transactionID,
		CampaignID:    transaction.CampaignID,
	}

	_, err = db.FirestoreClient.Collection("users").Doc(userID).Update(context.Background(), []firestore.Update{
		{Path: "Donations", Value: firestore.ArrayUnion(donation)},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user donations"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction data stored successfully"})
}
