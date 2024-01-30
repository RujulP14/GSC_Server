// routes/user_routes.go
package routes

import (
	"Server/controllers"

	"github.com/gin-gonic/gin"
)

func SetupChatbotRoutes(router *gin.Engine) {
	chatbotGroup := router.Group("/api/chatbot")
	{
		chatbotGroup.POST("/", controllers.GetChatbotResponse)
	}
}
