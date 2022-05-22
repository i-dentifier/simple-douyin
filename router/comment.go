package router

import "github.com/gin-gonic/gin"

func CommentRouterInit(r *gin.RouterGroup) {
	// extra apis - I
	commentRouter := r.Group("/comment")
	{
		commentRouter.POST("/action", nil)
		commentRouter.GET("/list", nil)
	}
}
