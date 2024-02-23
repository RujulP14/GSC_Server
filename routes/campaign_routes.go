// routes/campaign_routes.go
package routes

import (
	"Server/controllers"

	"github.com/gin-gonic/gin"
)

func SetupCampaignRoutes(router *gin.Engine) {
	campaignGroup := router.Group("/api/campaigns")
	{
		campaignGroup.POST("/", controllers.CreateCampaign)
		campaignGroup.GET("/", controllers.GetCampaigns)

		campaignGroup.GET("/:id", controllers.GetCampaign)
		campaignGroup.PUT("/:id", controllers.UpdateCampaign)
		campaignGroup.DELETE("/:id", controllers.DeleteCampaign)

		campaignGroup.PATCH("/:id/update-image", controllers.UpdateCampaignImage)
	}
}
