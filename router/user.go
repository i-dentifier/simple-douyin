package router

import (
	"github.com/gin-gonic/gin"
	usercontroller "simple-douyin/controller/user"
	"simple-douyin/middleware"
)

func UserRouterInit(r *gin.RouterGroup) {
	// basic apis
	userRouter := r.Group("/user")
	{
		userRouter.GET("/", usercontroller.UserInfo)
		userRouter.POST("/login", middleware.Validate(), usercontroller.Login)
		userRouter.POST("/register", middleware.Validate(), usercontroller.Register)
	}
}
