package relationcontroller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-douyin/model"
	relationservice "simple-douyin/service/relation"
	"strconv"
)

func Action(c *gin.Context) {
	to_user_id, err1 := strconv.ParseUint(c.Query("to_user_id"), 10, 32)
	action_type, err2 := strconv.ParseUint(c.Query("action_type"), 10, 32)

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: -1, StatusMsg: "invalid to_user_id or action_type",
		})
		return
	}

	claims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusOK, model.UserResponse{
			Response: model.Response{
				StatusCode: -1,
				StatusMsg:  "you must login first!",
			},
		})
		return
	}

	UserClaims := claims.(*model.UserClaims)

	relationId, err := relationservice.Action(UserClaims.UserId, uint32(to_user_id), uint8(action_type))

	if err != nil {
		// 关注失败
		c.JSON(http.StatusOK, model.ActionResponse{
			Response: model.Response{StatusCode: -1, StatusMsg: err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, model.ActionResponse{
		Response:   model.Response{StatusCode: 0, StatusMsg: "success"},
		RelationId: relationId,
	})
}
