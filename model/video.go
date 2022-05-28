package model

import "time"

type Video struct {
	Id            uint32 `json:"id,omitempty" gorm:"primaryKey"`
	Author        User   `json:"author" gorm:"-"`
	UserId        uint32
	PlayUrl       string    `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string    `json:"cover_url,omitempty"`
	FavoriteCount uint32    `json:"favorite_count,omitempty"`
	CommentCount  uint32    `json:"comment_count,omitempty"`
	IsFavorite    bool      `json:"is_favorite,omitempty"`
	Title         string    `json:"title"`
	CreateAt      time.Time `gorm:"autoCreateTime;index"`
}

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}
