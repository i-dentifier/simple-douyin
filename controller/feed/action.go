package feedcontroller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-douyin/model"
	feedservice "simple-douyin/service/feed"
)

type FeedResponse struct {
	model.Response
	VideoList []*model.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

func Feed(c *gin.Context) {

	latestTime := c.Query("latest_time")

	var userId uint32
	var isLogin bool
	if token := c.Query("token"); token != "" {
		claims, _ := c.Get("user")
		userId = claims.(*model.UserClaims).UserId
		isLogin = true
	}

	Videos, err := feedservice.Feed(latestTime, userId, isLogin)

	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response:  model.Response{StatusCode: 0, StatusMsg: "success"},
		VideoList: Videos,
		NextTime:  Videos[len(Videos)-1].CreateAt.Unix(),
	})
}
