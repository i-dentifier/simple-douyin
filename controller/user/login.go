package usercontroller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-douyin/middleware"
	"simple-douyin/model"
	"simple-douyin/service/user"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
// var usersLoginInfo = map[string]controller.User{
// 	"zhangleidouyin": {
// 		Id:            1,
// 		Name:          "zhanglei",
// 		FollowCount:   10,
// 		FollowerCount: 5,
// 		IsFollow:      true,
// 	},
// }

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	userId, errLogin := userservice.Login(username, password)
	if errLogin != nil {
		// 登录失败
		c.JSON(http.StatusOK, model.UserLoginResponse{
			Response: model.Response{StatusCode: -1, StatusMsg: errLogin.Error()},
		})
		return
	}
	// 生成token
	token, errToken := middleware.GenerateToken(userId)
	if errToken != nil {
		// token生成失败
		c.JSON(http.StatusOK, model.UserLoginResponse{
			Response: model.Response{StatusCode: -1, StatusMsg: errToken.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, model.UserLoginResponse{
		Response: model.Response{StatusCode: 0, StatusMsg: "success"},
		UserId:   userId,
		Token:    token,
	})
}
