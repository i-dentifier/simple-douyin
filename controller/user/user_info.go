package usercontroller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-douyin/model"
	userservice "simple-douyin/service/user"
	"strconv"
)

func UserInfo(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Query("user_id"), 10, 32)
	// 非法userId引起的error
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: -1, StatusMsg: "invalid user id",
		})
		return
	}
	// token经middleware验证合法后将存入context
	claims, exist := c.Get("user")
	// 如果没有在context中查到用户说明未登录
	if !exist {
		c.JSON(http.StatusOK, model.UserResponse{
			Response: model.Response{
				StatusCode: -1,
				StatusMsg:  "login required",
			},
		})
		return
	}
	userClaims := claims.(*model.UserClaims)
	user, err := userservice.QueryUserInfoById(uint32(userId), userClaims.UserId)
	if err != nil {
		c.JSON(http.StatusOK, model.UserResponse{
			Response: model.Response{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	c.JSON(http.StatusOK, model.UserResponse{
		Response: model.Response{StatusCode: 0, StatusMsg: "success"},
		User:     *user,
	})
}
