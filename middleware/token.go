package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"simple-douyin/model"
	"strings"
	"time"
)

// 所有变量和方法仅供包内访问

// 生成token所需内容
var (
	// 自定义token密钥
	secret = []byte("32974doq380dbm38")
	// token有效时间(nanosecond)
	effectTime = 1 * time.Hour
	// 不需要校验token的路由
	// noVerifyToken = []string{"/douyin/user/login", "/douyin/user/register", "/douyin/feed"}
	noVerifyToken = map[string]bool{"/douyin/user/login/": true,
		"/douyin/user/register/": true, "/douyin/feed/": true}
)

// GenerateToken 生成token
func GenerateToken(userId uint32) (string, error) {
	// userClaims 用于生成token的自定义claims
	claims := model.UserClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			// 签发时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			// 过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(effectTime)),
		},
	}
	// 生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(secret)
	if err != nil {
		return "", fmt.Errorf("token generation failed, caused by %w", err)
	}
	// 将签发的(userId, claims)键值对
	// 存储在controller.LoginInfoMap中
	//common.LoginInfoMap[userId] = claims
	return tokenStr, nil
}

// VerifyToken 验证token合法性
// 该方法是大部分请求的前置条件
func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// noVerifyToken中的请求不需要校验token
		req := c.Request.RequestURI
		//截取原始的URL，修改重定向
		req = strings.Split(req, "?")[0]
		fmt.Println(req)
		if noVerifyToken[req] {
			return
		}
		// for _, str := range noVerifyToken {
		//	if req == str {
		//		return
		//	}
		// }

		// token可能是GET或者POST传入
		token := c.Query("token")
		if token == "" {
			token = c.PostForm("token")
		}

		// 校验token合法性
		claims, err := parseToken(token)
		// token非法
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.Response{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			})
			return
		}
		// 校验token过期时间
		if !(claims.ExpiresAt.Time.Unix() > time.Now().Unix()) {
			// 过期删除token
			// delete(common.LoginInfoMap, claims.UserId)
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.Response{
				StatusCode: -1,
				StatusMsg:  "login required",
			})
			return
		}

		// 刷新token时间
		refreshToken(claims)
		// 放入上下文方便controller读取
		c.Set("user", claims)
	}
}

// parseToken 验证token合法性
func parseToken(tokenStr string) (*model.UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &model.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	// token非法
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*model.UserClaims)
	if !ok {
		return nil, err
	}
	return claims, err
}

// refreshToken 更新token的过期时间
func refreshToken(claims *model.UserClaims) {
	// 下次过期时间设置为当前时间往后推2h
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(2 * time.Hour))
}
