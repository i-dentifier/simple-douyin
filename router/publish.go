package router

import (
	"github.com/gin-gonic/gin"
)

func PublishRouterInit(r *gin.RouterGroup) {
	// basic apis
	publishRouter := r.Group("/publish")
	{
		publishRouter.POST("/action/", nil)
		publishRouter.GET("/list/", nil)
	}
}
