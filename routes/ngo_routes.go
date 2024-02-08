// routes/ngo_routes.go
package routes

import (
	"Server/controllers"

	"github.com/gin-gonic/gin"
)

func SetupNGORoutes(router *gin.Engine) {
	ngoGroup := router.Group("/api/ngos")
	{
		ngoGroup.POST("/signup", controllers.SignupNGO)
		ngoGroup.GET("/:id", controllers.GetNGO)
		ngoGroup.GET("/", controllers.GetNGOs)
		ngoGroup.PUT("/:id", controllers.UpdateNGO)
		ngoGroup.DELETE("/:id", controllers.DeleteNGO)
		ngoGroup.POST("/login", controllers.LoginNGO)
	}
}
