package usercontroller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-douyin/common"
	"simple-douyin/service/userservice"
	"strconv"
)

func UserInfo(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Query("user_id"), 10, 32)
	// 非法userId引起的error
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: -1, StatusMsg: "invalid user id",
		})
		return
	}
	// token经middleware验证合法后将存入context
	claims, exist := c.Get("user")
	// 如果没有在context中查到用户说明未登录
	if !exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: common.Response{
				StatusCode: -1,
				StatusMsg:  "login required",
			},
		})
		return
	}
	userClaims := claims.(*common.UserClaims)
	user, err := userservice.QueryUserInfoById(uint32(userId), userClaims.UserId)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: common.Response{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	c.JSON(http.StatusOK, UserResponse{
		Response: common.Response{StatusCode: 0, StatusMsg: "success"},
		User:     *user,
	})
}
