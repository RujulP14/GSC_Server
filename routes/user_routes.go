// routes/user_routes.go
package routes

import (
	"Server/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/api/users")
	{
		userGroup.GET("/:id", controllers.GetUser)
		userGroup.GET("/", controllers.GetUsers)
		userGroup.PUT("/:id", controllers.UpdateUser)
		userGroup.DELETE("/:id", controllers.DeleteUser)

		userGroup.POST("/signup", controllers.SignupUser)
		userGroup.POST("/login", controllers.LoginUser)

		userGroup.POST("/:id/favourite-article", controllers.AddToFavorites)
		userGroup.DELETE("/:id/favourite-article", controllers.RemoveFromFavorites)
	}
}
