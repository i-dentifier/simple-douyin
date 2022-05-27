package usercontroller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-douyin/middleware"
	"simple-douyin/model"
	userservice "simple-douyin/service/user"
)

func Register(c *gin.Context) {

	username := c.Query("username")
	password := c.Query("password")
	userId, err := userservice.Register(username, password)
	if err != nil {
		// 注册失败
		c.JSON(http.StatusOK, model.UserRegisterResponse{
			Response: model.Response{StatusCode: -1, StatusMsg: err.Error()},
		})
		return
	}
	// 注册成功后自动登录
	// 这里用户名和密码一定是正确的
	// 所以忽略error
	// userId, _ := userservice.Login(username, password)

	// 生成token
	token, errToken := middleware.GenerateToken(userId)
	if errToken != nil {
		// token生成失败
		c.JSON(http.StatusOK, model.UserRegisterResponse{
			Response: model.Response{StatusCode: -1, StatusMsg: errToken.Error()},
		})
		return
	}
	//登录
	_, err = userservice.Login(username, password)
	if err != nil {
		// 注册跳转登录失败
		c.JSON(http.StatusOK, model.UserLoginResponse{
			Response: model.Response{StatusCode: -1, StatusMsg: err.Error()},
		})
		return
	}
	c.JSON(http.StatusOK, model.UserRegisterResponse{
		Response: model.Response{StatusCode: 0, StatusMsg: "success"},
		UserId:   userId,
		Token:    token,
	})
}
