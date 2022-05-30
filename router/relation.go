package router

import (
	"github.com/gin-gonic/gin"
	relationcontroller "simple-douyin/controller/relation"
)

func RelationRouterInit(r *gin.RouterGroup) {
	// extra apis - II
	commentRouter := r.Group("/relation")
	{
		commentRouter.POST("/action/", relationcontroller.Action)
		commentRouter.GET("follow/list/", relationcontroller.FollowList)
		commentRouter.GET("follower/list/", nil)
	}
}
