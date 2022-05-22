package main

import (
	"github.com/gin-gonic/gin"
	"simple-douyin/config"
	"simple-douyin/router"
)

func main() {
	config.DBInit()
	r := gin.Default()
	// 除个别请求外均需先验证token
	router.InitRouter(r)
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
