package feedcontroller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-douyin/model"
	"time"
)

type FeedResponse struct {
	model.Response
	VideoList []model.Video `json:"video_list,omitempty"`
	NextTime  int64         `json:"next_time,omitempty"`
}

func Feed(c *gin.Context) {

	_ = c.PostForm("last_time")

	c.JSON(http.StatusOK, FeedResponse{
		Response:  model.Response{StatusCode: 0, StatusMsg: "success"},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
