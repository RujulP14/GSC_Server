package controllers

import (
	"net/http"

	"Server/models" // Import the Transaction model

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
)

// HandleStripeDonation handles donations using Stripe
func HandleStripeDonation(c *gin.Context) {
	// Set up Stripe API key
	stripe.Key = "sk_test_51Ogl1oSEhw22DH1HkwaWv9XUQNfhnBxgt3CU8Hm8kzxbxyy1PmNlkk4t0f1wJQiaVUkL6XDPFk9jBOMmULvFqTFS002DTkLvpz"

	// Parse request body to get donation amount and currency
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a Stripe checkout session
	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String(transaction.Currency),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("Donation"),
					},
					UnitAmount: stripe.Int64(transaction.Amount), // Amount in cents
				},
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String("https://yourwebsite.com/success"),
		CancelURL:  stripe.String("https://yourwebsite.com/cancel"),
	}
	session, err := session.New(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create checkout session"})
		return
	}

	// Return the session ID to the client
	c.JSON(http.StatusOK, gin.H{"sessionId": session.ID})
}