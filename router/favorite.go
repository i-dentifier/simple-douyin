package router

import "github.com/gin-gonic/gin"

func FavoriteRouterInit(r *gin.RouterGroup) {
	// extra apis - I
	favoriteRouter := r.Group("/favorite")
	{
		favoriteRouter.POST("/action/", nil)
		favoriteRouter.GET("/list/", nil)
	}
}
