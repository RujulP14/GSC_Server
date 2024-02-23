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

		userGroup.POST("/:id/favourite-article", controllers.AddToFavouritesArticles)
		userGroup.DELETE("/:id/favourite-article", controllers.RemoveFromFavouritesArticles)

		userGroup.POST("/:id/favourite-blog", controllers.AddToFavouriteBlogs)
		userGroup.DELETE("/:id/favourite-blog", controllers.RemoveFromFavouriteBlogs)

		userGroup.POST("/:id/favourite-video", controllers.AddToFavouriteVideos)
		userGroup.DELETE("/:id/favourite-video", controllers.RemoveFromFavouriteVideos)

		userGroup.GET("/:id/blogs", controllers.GetUserBlogs)
	}
}
