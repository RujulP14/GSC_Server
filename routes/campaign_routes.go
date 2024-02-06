// routes/campaign_routes.go
package routes

import (
	"Server/controllers"

	"github.com/gin-gonic/gin"
)

func SetupCampaignRoutes(router *gin.Engine) {
	campaignGroup := router.Group("/api/campaign")
	{
		campaignGroup.POST("/", controllers.CreateCampaign)
		campaignGroup.GET("/:id", controllers.GetCampaign)
		campaignGroup.GET("/", controllers.GetCampaigns)
		campaignGroup.PUT("/:id", controllers.UpdateCampaign)
		campaignGroup.DELETE("/:id", controllers.DeleteCampaign)
		campaignGroup.PATCH("/:id/update-image", controllers.UpdateCampaignImage)
		campaignGroup.PATCH("/:id/add-donor", controllers.AddDonorToCampaign)
		campaignGroup.PATCH("/:id/remove-donor", controllers.RemoveDonorFromCampaign)
	}
}
