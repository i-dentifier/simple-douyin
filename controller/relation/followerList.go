package relationcontroller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-douyin/model"
	relationservice "simple-douyin/service/relation"
)

func FollowerList(c *gin.Context) {

	// token经middleware验证合法后将存入context
	claims, exists := c.Get("user")
	// 如果没有在context中查到用户说明未登录
	if !exists {
		c.JSON(http.StatusOK, model.UserResponse{
			Response: model.Response{
				StatusCode: -1,
				StatusMsg:  "login required to get follower list",
			},
		})
		return
	}
	userClaims := claims.(*model.UserClaims)
	userId := userClaims.UserId
	followerList, err := relationservice.FollowerList(userId)
	if err != nil {
		c.JSON(http.StatusOK, model.UserResponse{
			Response: model.Response{
				StatusCode: -1,
				StatusMsg:  "get follower list error",
			},
		})
		return
	}

	c.JSON(http.StatusOK, model.FollowerListResponse{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		FollowerList: followerList,
	})
}
