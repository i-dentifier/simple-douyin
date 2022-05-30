package favoritecontroller

import (
	"net/http"
	"simple-douyin/model"
	favoriteservice "simple-douyin/service/favorite"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Action(c *gin.Context) {

	videoIdString := c.Query("video_id")
	actionType := c.Query("action_type")

	// token经middleware验证合法后将存入context
	claims, exists := c.Get("user")
	// 如果没有在context中查到用户说明未登录
	if !exists {
		c.JSON(http.StatusOK, model.UserResponse{
			Response: model.Response{
				StatusCode: -1,
				StatusMsg:  "login required to favorite this video",
			},
		})
		return
	}
	userClaims := claims.(*model.UserClaims)

	videoId, err := strconv.ParseUint(videoIdString, 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, model.UserResponse{
			Response: model.Response{
				StatusCode: -1,
				StatusMsg:  "video id error",
			},
		})
		return
	}

	// 执行点赞服务
	if err := favoriteservice.Action(userClaims.UserId, uint32(videoId), actionType); err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		StatusCode: 0,
		StatusMsg:  "success do favorite action",
	})
}
