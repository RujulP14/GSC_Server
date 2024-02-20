// routes/article_routes.go
package routes

import (
	"Server/controllers"

	"github.com/gin-gonic/gin"
)

func SetupArticleRoutes(router *gin.Engine) {
	articleGroup := router.Group("/api/articles")
	{
		articleGroup.POST("/", controllers.CreateArticle)
		articleGroup.GET("/", controllers.GetAllArticles)

		articleGroup.GET("/:id", controllers.GetArticle)
		articleGroup.PUT("/:id", controllers.UpdateArticle)
		articleGroup.DELETE("/:id", controllers.DeleteArticle)

		articleGroup.POST("/:id/comments", controllers.AddCommentOnArticle)
		articleGroup.DELETE("/:id/comments", controllers.RemoveCommentOnArticle)
	}
}
