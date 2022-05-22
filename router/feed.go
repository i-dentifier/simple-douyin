package router

import (
	"github.com/gin-gonic/gin"
)

func FeedRouterInit(r *gin.RouterGroup) {
	// basic apis
	feedRouter := r.Group("/feed")
	{
		//r.GET("/", controller.Feed)
		feedRouter.GET("/", nil)
	}
}
