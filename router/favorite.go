package router

import (
	favoritecontroller "simple-douyin/controller/favorite"

	"github.com/gin-gonic/gin"
)

func FavoriteRouterInit(r *gin.RouterGroup) {
	// extra apis - I
	favoriteRouter := r.Group("/favorite")
	{
		favoriteRouter.POST("/action/", favoritecontroller.Action)
		favoriteRouter.GET("/list/", favoritecontroller.List)
	}
}
