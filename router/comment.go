package router

import (
	"github.com/gin-gonic/gin"
	commentcontroller "simple-douyin/controller/comment"
)

func CommentRouterInit(r *gin.RouterGroup) {
	// extra apis - I
	commentRouter := r.Group("/comment")
	{
		commentRouter.POST("/action/", commentcontroller.Action)
		commentRouter.GET("/list/", commentcontroller.List)
	}
}
