package router

import "github.com/gin-gonic/gin"

func RelationRouterInit(r *gin.RouterGroup) {
	// extra apis - II
	commentRouter := r.Group("/relation")
	{
		commentRouter.POST("/action/", nil)
		commentRouter.GET("follow/list/", nil)
		commentRouter.GET("follower/list/", nil)
	}
}
