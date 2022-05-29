package model

import "time"

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64     `json:"id,omitempty"`
	Author        User      `json:"author" gorm:"foreignKey:Id;references:Id"`
	PlayUrl       string    `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string    `json:"cover_url,omitempty"`
	FavoriteCount int64     `json:"favorite_count,omitempty"`
	CommentCount  int64     `json:"comment_count,omitempty"`
	IsFavorite    bool      `json:"is_favorite,omitempty"`
	Title         string    `json:"title,omitempty"`
	CreateAt      time.Time `json:"-" gorm:"autoCreateTime;index"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

// Relationship 对应一个关注/粉丝关系
type Relationship struct {
	Id         uint64 `json:"id,omitempty"`
	FromUserId uint32 `json:"from_user_id,omitempty"`
	ToUserId   uint32 `json:"to_user_id,omitempty"`
	Status     uint8  `json:"status,omitempty"`
}
