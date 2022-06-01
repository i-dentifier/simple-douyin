package model

import "time"

type Comment struct {
	Id         int64     `json:"id,omitempty" gorm:"primaryKey"` // 评论自增id
	User       *User     `json:"user" gorm:"foreignKey:UserId;references:Id"`
	UserId     int64     `json:"-"`
	VideoId    int64     `json:"-"` // 视频id
	Content    string    `json:"content,omitempty"`
	CreateTime time.Time `json:"-" gorm:"autoCreateTime"`
	CreateDate string    `json:"create_date,omitempty" gorm:"-"`
}

type CommentActionResponse struct {
	Response
	Comment *Comment `json:"comment"`
}

type CommentListResponse struct {
	Response
	CommentList []*Comment `json:"comment_list"`
}
