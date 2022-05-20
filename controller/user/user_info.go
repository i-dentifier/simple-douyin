package usercontroller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-douyin/common"
	"simple-douyin/service/userservice"
	"strconv"
)

func UserInfo(c *gin.Context) {
	var userId uint32
	uid, err := strconv.ParseUint(c.Query("user_id"), 10, 32)
	userId = uint32(uid)
	// 非法userId引起的error
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: -1, StatusMsg: "invalid user_id",
		})
		return
	}
	// token经middleware验证合法后将存入common.LoginInfoMap
	// 如果没有在common.LoginInfoMap中查到用户说明未登录
	if _, exist := common.LoginInfoMap[userId]; !exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: common.Response{
				StatusCode: -1,
				StatusMsg:  "please login first",
			},
		})
		return
	}
	user, err := userservice.QueryUserInfoById(userId)
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
