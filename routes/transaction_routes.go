package routes

import (
	"Server/controllers"

	"github.com/gin-gonic/gin"
)

func SetupStripeRoutes(router *gin.Engine) {
	transactionGroup := router.Group("/api/transactions")
	{
		transactionGroup.POST("/donate", controllers.HandleDonation)
	}
}
