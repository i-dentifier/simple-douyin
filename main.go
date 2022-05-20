package main

import (
	"github.com/gin-gonic/gin"
	"simple-douyin/config"
	"simple-douyin/middleware"
)

func main() {
	config.DBInit()
	r := gin.Default()
	// 除个别请求外均需先验证token
	r.Use(middleware.VerifyToken)
	initRouter(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
