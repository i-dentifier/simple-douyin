package publishcontroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-douyin/model"
	"simple-douyin/service/publish"
	"strconv"
)

func PublishList(c *gin.Context) {
	// 1.获取userid
	userIdStr := c.Query("user_id")
	userIdInt, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusOK, model.UserResponse{
			Response: model.Response{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	userId := uint32(userIdInt)
	// 2.获取已鉴权后的claims
	// token经middleware验证合法后将存入context
	claims, exist := c.Get("user")
	userClaims := claims.(*model.UserClaims)
	// 如果没有在context中查到用户说明未登录
	if !exist {
		c.JSON(http.StatusOK, model.UserResponse{
			Response: model.Response{
				StatusCode: -1,
				StatusMsg:  fmt.Sprintf("user: %v,token time expired", userClaims.UserId),
			},
		})
		return
	}
	// 3.调用service处理
	videoList, err := publishservice.PublishList(userId)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.VideoListResponse{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "success to get video list",
		},
		VideoList: videoList,
	})
}
