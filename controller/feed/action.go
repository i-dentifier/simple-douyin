package feedcontroller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-douyin/model"
	feedservice "simple-douyin/service/feed"
	"time"
)

type FeedResponse struct {
	model.Response
	VideoList []model.Video `json:"video_list,omitempty"`
	NextTime  int64         `json:"next_time,omitempty"`
}

func Feed(c *gin.Context) {

	last_time := c.Query("last_time")
	token := c.Query("token")

	Videos, _ := feedservice.Feed(last_time, token)

	c.JSON(http.StatusOK, FeedResponse{
		Response:  model.Response{StatusCode: 0, StatusMsg: "success"},
		VideoList: Videos,
		NextTime:  time.Now().Unix(),
	})
}
