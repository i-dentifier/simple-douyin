package model

import "github.com/golang-jwt/jwt/v4"

type User struct {
	// 主键id, 关注与粉丝数量均使用uint32
	// 与数据库`int unsigned`对应
	// 可以支持约40亿(4294967295)用户
	Id            uint32 `json:"id"`
	Name          string `json:"name"`
	FollowCount   uint32 `json:"follow_count"`
	FollowerCount uint32 `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type UserAuth struct {
	// 只读
	Id uint32 `gorm:"->"`
	// 允许读和修改
	Name string `gorm:"<-"`
	// 允许读和修改
	Password string `gorm:"<-"`
}

type UserLoginResponse struct {
	Response
	UserId uint32 `json:"user_id"`
	Token  string `json:"token"`
}

type UserRegisterResponse struct {
	Response
	UserId uint32 `json:"user_id"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

// UserClaims 用于生成和解析token
type UserClaims struct {
	UserId uint32
	// jwt-go提供的标准claims
	jwt.RegisteredClaims
}

// LoginInfoMap 验证用户登录状态的map
// key: 用户的userId
// value: 对应用户的claims(包含userId, userName, expiredTime等)
// var LoginInfoMap = make(map[uint32]UserClaims, 20)
