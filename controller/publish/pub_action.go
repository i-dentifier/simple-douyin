package publishcontroller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-douyin/model"
	"simple-douyin/service/publish"
)

func Publish(c *gin.Context) {
	// 1.data
	// data := c.PostForm("data")
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	// 2.获取已鉴权后的claims
	// token经middleware验证合法后将存入context
	claims, exist := c.Get("user")
	// 如果没有在context中查到用户说明未登录
	if !exist {
		c.JSON(http.StatusOK, model.UserResponse{
			Response: model.Response{
				StatusCode: -1,
				StatusMsg:  "login required to publish videos",
			},
		})
		return
	}
	userClaims := claims.(*model.UserClaims)
	// 3.title
	title := c.PostForm("title")
	// 4.调用service处理
	if err := publishservice.Publish(data, title, userClaims.UserId); err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		StatusCode: 0,
		StatusMsg:  "success publish video",
	})
}
