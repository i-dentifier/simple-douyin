package common

import (
	"github.com/dgrijalva/jwt-go/v4"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type User struct {
	// 主键id, 关注与粉丝数量均使用uint32
	// 与数据库`int unsigned`对应
	// 可以支持约40亿(4294967295)用户
	Id            uint32 `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   uint32 `json:"follow_count,omitempty"`
	FollowerCount uint32 `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

// Relationship 对应一个关注/粉丝关系
type Relationship struct {
	Id         uint64 `json:"id,omitempty"`
	FromUserId uint32 `json:"from_user_id,omitempty"`
	ToUserId   uint32 `json:"to_user_id,omitempty"`
	Status     uint8  `json:"status,omitempty"`
}

type UserAuth struct {
	// 只读
	Id uint32 `gorm:"->"`
	// 允许读和修改
	Name string `gorm:"<-"`
	// 允许读和修改
	Password string `gorm:"<-"`
}

// UserClaims 用于生成和解析token
type UserClaims struct {
	UserId uint32
	// jwt-go提供的标准claims
	jwt.StandardClaims
}

// LoginInfoMap 验证用户登录状态的map
// key: 用户的userId
// value: 对应用户的claims(包含userId, userName, expiredTime等)
// var LoginInfoMap = make(map[uint32]UserClaims, 20)
