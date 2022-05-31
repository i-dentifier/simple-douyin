package favoritecontroller

import (
	"net/http"
	"simple-douyin/model"
	favoriteservice "simple-douyin/service/favorite"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	// token经middleware验证合法后将存入context
	claims, exists := c.Get("user")
	// 如果没有在context中查到用户说明未登录
	if !exists {
		c.JSON(http.StatusOK, model.UserResponse{
			Response: model.Response{
				StatusCode: -1,
				StatusMsg:  "login required to get favorite list",
			},
		})
		return
	}
	userClaims := claims.(*model.UserClaims)
	userId := userClaims.UserId
	videoList, err := favoriteservice.List(userId)
	if err != nil {
		c.JSON(http.StatusOK, model.UserResponse{
			Response: model.Response{
				StatusCode: -1,
				StatusMsg:  "get favorite list error",
			},
		})
		return
	}

	c.JSON(http.StatusOK, model.VideoListResponse{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "success to get favorite video list",
		},
		VideoList: videoList,
	})
}
