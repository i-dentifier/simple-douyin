package usercontroller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-douyin/common"
	"simple-douyin/middleware"
	"simple-douyin/service/userservice"
)

type UserRegisterResponse struct {
	common.Response
	UserId uint32 `json:"user_id"`
	Token  string `json:"token"`
}

type UserResponse struct {
	common.Response
	User common.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	err := userservice.Register(username, password)
	if err != nil {
		// 注册失败
		c.JSON(http.StatusOK, UserRegisterResponse{
			Response: common.Response{StatusCode: -1, StatusMsg: err.Error()},
		})
		return
	}
	// 注册成功后自动登录
	// 这里用户名和密码一定是正确的
	// 所以忽略error
	userId, _ := userservice.Login(username, password)
	// 生成token
	token, errToken := middleware.GenerateToken(userId)
	if errToken != nil {
		// token生成失败
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{StatusCode: -1, StatusMsg: errToken.Error()},
		})
		return
	}
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: common.Response{StatusCode: 0, StatusMsg: "success"},
		UserId:   userId,
		Token:    token,
	})
}
