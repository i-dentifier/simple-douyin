package commentcontroller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-douyin/model"
	commentservice "simple-douyin/service/comment"
	"strconv"
)

func List(c *gin.Context) {

	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 32)

	commentList, err := commentservice.GetCommentList(videoId)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		},
		)
		return
	}

	c.JSON(http.StatusOK, model.CommentListResponse{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "successfully fetch comment_list",
		},
		CommentList: commentList,
	})
}
