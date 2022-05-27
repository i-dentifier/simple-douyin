package router

import (
	"github.com/gin-gonic/gin"
	feedcontroller "simple-douyin/controller/feed"
)

func FeedRouterInit(r *gin.RouterGroup) {
	// basic apis
	feedRouter := r.Group("/feed")
	{
		feedRouter.GET("/", feedcontroller.Feed)
	}
}
