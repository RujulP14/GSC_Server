// routes/initiative_routes.go
package routes

import (
	"Server/controllers"

	"github.com/gin-gonic/gin"
)

func SetupInitiativeRoutes(router *gin.Engine) {
	initiativeGroup := router.Group("/api/initiatives")
	{
		initiativeGroup.POST("/", controllers.CreateInitiative)
		initiativeGroup.GET("/:id", controllers.GetInitiative)
		initiativeGroup.GET("/", controllers.GetInitiatives)
		initiativeGroup.PUT("/:id", controllers.UpdateInitiative)
		initiativeGroup.DELETE("/:id", controllers.DeleteInitiative)
		initiativeGroup.PATCH("/:id/update-image", controllers.UpdateInitiativeImage)
	}
}
