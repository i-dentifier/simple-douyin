package middleware

import "C"
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"simple-douyin/model"
)

func Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password")
		if ok := validateName(username); !ok {
			c.AbortWithStatusJSON(http.StatusOK, model.Response{
				StatusCode: -1,
				StatusMsg: "username should start with letters" +
					" and only use letters and digits with length(8-16)",
			})
			return
		}
		if ok := validatePassword(password); !ok {
			c.AbortWithStatusJSON(http.StatusOK, model.Response{
				StatusCode: -1,
				StatusMsg: "password should only use letters" +
					" and digits with length(8-16)",
			})
			return
		}
	}
}

func validateName(username string) bool {
	if ok, _ := regexp.MatchString("^[a-zA-Z][A-Za-z0-9]{7,15}$", username); !ok {
		fmt.Println(ok)
		return false
	}
	return true

}

func validatePassword(password string) bool {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{8,16}$", password); !ok {
		return false
	}
	return true
}
