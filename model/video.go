package model

import "time"

type Video struct {
	Id            uint32    `json:"id" gorm:"primaryKey"`
	Author        *User     `json:"author" gorm:"foreignKey:UserId;references:Id"`
	UserId        uint32    `json:"-" gorm:"index"`
	PlayUrl       string    `json:"play_url"`
	CoverUrl      string    `json:"cover_url"`
	FavoriteCount uint32    `json:"favorite_count"`
	CommentCount  uint32    `json:"comment_count"`
	IsFavorite    bool      `json:"is_favorite"`
	Title         string    `json:"title"`
	CreateAt      time.Time `json:"-" gorm:"autoCreateTime;index"`
}

type VideoListResponse struct {
	Response
	VideoList []*Video `json:"video_list"`
}
