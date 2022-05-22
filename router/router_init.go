package router

import (
	"github.com/gin-gonic/gin"
	"simple-douyin/middleware"
)

func InitRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")
	apiRouter.Use(middleware.VerifyToken())
	UserRouterInit(apiRouter)
	FeedRouterInit(apiRouter)
	PublishRouterInit(apiRouter)
	FavoriteRouterInit(apiRouter)
	CommentRouterInit(apiRouter)
	RelationRouterInit(apiRouter)
}
