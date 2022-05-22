package main

import (
	"github.com/gin-gonic/gin"
	"simple-douyin/config"
	"simple-douyin/router"
)

func main() {
	config.DBInit()
	r := gin.Default()
	router.InitRouter(r)
	r.Run(":8080")
}
