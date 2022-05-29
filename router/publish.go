package router

import (
	"github.com/gin-gonic/gin"
	"simple-douyin/controller/publish"
)

func PublishRouterInit(r *gin.RouterGroup) {
	// basic apis
	publishRouter := r.Group("/publish")
	{
		publishRouter.POST("/action/", publishcontroller.Publish)
		publishRouter.GET("/list/", publishcontroller.PublishList)
	}
}
