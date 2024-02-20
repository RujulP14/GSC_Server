// routes/video_routes.go
package routes

import (
	"Server/controllers"

	"github.com/gin-gonic/gin"
)

func SetupVideoRoutes(router *gin.Engine) {
	videoGroup := router.Group("/api/videos")
	{
		videoGroup.POST("/", controllers.CreateVideo)
		videoGroup.GET("/", controllers.GetVideos)

		videoGroup.GET("/:id", controllers.GetVideo)
		videoGroup.PUT("/:id", controllers.UpdateVideo)
		videoGroup.DELETE("/:id", controllers.DeleteVideo)

		videoGroup.POST("/:id/comments", controllers.AddCommentOnVideo)
		videoGroup.DELETE("/:id/comments", controllers.RemoveCommentOnVideo)
		videoGroup.POST("/:id/like", controllers.LikeVideo)
	}
}
