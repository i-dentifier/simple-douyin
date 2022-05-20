package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"simple-douyin/common"
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
	noVerifyToken = []string{"/login", "/register", "/feed"}
)

// GenerateToken 生成token
func GenerateToken(userId uint32) (string, error) {
	// userClaims 用于生成token的自定义claims
	claims := common.UserClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			// 签发时间
			IssuedAt: jwt.At(time.Now()),
			// 过期时间
			ExpiresAt: jwt.At(time.Now().Add(effectTime)),
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
	common.LoginInfoMap[userId] = claims
	return tokenStr, nil
}

// VerifyToken 验证token合法性
// 该方法是大部分请求的前置条件
// 其余请求在该操作后直接查询controller.LoginInfoMap
// 来获得用户信息，如果为空说明未登录
func VerifyToken(c *gin.Context) {
	// token可能是GET或者POST传入
	token := c.Query("token")
	if token == "" {
		token = c.PostForm("token")
	}
	// noVerifyToken中的请求不需要校验token
	for _, str := range noVerifyToken {
		if token == str {
			return
		}
	}

	// 校验token合法性
	claims, err := parseToken(token)
	// token非法
	if err != nil {
		return
	}

	// 校验token过期时间
	if !(claims.ExpiresAt.Time.Unix() > time.Now().Unix()) {
		// 过期删除token
		delete(common.LoginInfoMap, claims.UserId)
		return
	}
	// 刷新token时间
	refreshExp(claims)
	// 将用户信息放入map
	common.LoginInfoMap[claims.UserId] = *claims
}

// parseToken 验证token合法性
func parseToken(tokenStr string) (*common.UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &common.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	// token非法
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*common.UserClaims)
	if !ok {
		return nil, err
	}
	return claims, err
}

// refreshExp 更新token的过期时间
func refreshExp(claims *common.UserClaims) {
	// 下次过期时间设置为当前时间往后推2h
	claims.ExpiresAt = jwt.At(time.Now().Add(2 * time.Hour))
}
