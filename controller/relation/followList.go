package relationcontroller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-douyin/model"
	relationservice "simple-douyin/service/relation"
)

func FollowList(c *gin.Context) {

	// token经middleware验证合法后将存入context
	claims, exists := c.Get("user")
	// 如果没有在context中查到用户说明未登录
	if !exists {
		c.JSON(http.StatusOK, model.UserResponse{
			Response: model.Response{
				StatusCode: -1,
				StatusMsg:  "login required to get follow list",
			},
		})
		return
	}
	userClaims := claims.(*model.UserClaims)
	userId := userClaims.UserId
	followList, err := relationservice.FollowList(userId)
	if err != nil {
		c.JSON(http.StatusOK, model.UserResponse{
			Response: model.Response{
				StatusCode: -1,
				StatusMsg:  "get follow list error",
			},
		})
		return
	}

	c.JSON(http.StatusOK, model.FollowListResponse{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		FollowList: followList,
	})
}
