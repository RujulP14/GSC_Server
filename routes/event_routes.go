// routes/event_routes.go
package routes

import (
	"Server/controllers"

	"github.com/gin-gonic/gin"
)

func SetupEventRoutes(router *gin.Engine) {
	eventRoutes := router.Group("/api/events")
	{
		eventRoutes.POST("/", controllers.CreateEvent)
		eventRoutes.GET("/", controllers.GetEvents)
		eventRoutes.GET("/:id", controllers.GetEventByID)
		eventRoutes.PUT("/:id", controllers.UpdateEvent)
		eventRoutes.DELETE("/:id", controllers.DeleteEvent)
		eventRoutes.POST("/:id/register", controllers.RegisterUserForEvent)
	}
}
