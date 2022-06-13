package commentcontroller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-douyin/model"
	commentservice "simple-douyin/service/comment"
	"strconv"
)

func Action(c *gin.Context) {

	// 官方接口使用的是int64 待讨论
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 32)
	videoId, err := strconv.ParseInt(c.Query("video_id"), 10, 32)
	actionType := c.Query("action_type")
	commentText := c.Query("comment_text")
	commentId := c.Query("comment_id")

	comment, err := commentservice.MakeComment(userId, videoId, actionType, commentText, commentId)

	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		},
		)
		return
	}

	var statusMsg string
	if actionType == "1" {
		statusMsg = "successfully make comment"
	} else {
		statusMsg = "successfully delete comment"
	}
	c.JSON(http.StatusOK, model.CommentActionResponse{
		model.Response{
			StatusCode: 0,
			StatusMsg:  statusMsg,
		},
		comment,
	})
}
