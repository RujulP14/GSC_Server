// routes/blog_routes.go
package routes

import (
	"Server/controllers"

	"github.com/gin-gonic/gin"
)

func SetupBlogRoutes(router *gin.Engine) {
	blogsGroup := router.Group("/api/blogs")
	{
		blogsGroup.POST("/", controllers.CreateBlog)
		blogsGroup.GET("/", controllers.GetAllBlogs)

		blogsGroup.GET("/:id", controllers.GetBlog)
		blogsGroup.PUT("/:id", controllers.UpdateBlog)
		blogsGroup.DELETE("/:id", controllers.DeleteBlog)
	}
}
