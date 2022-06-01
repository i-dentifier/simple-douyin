package model

type Favorite struct {
	UserId  uint32 `gorm:"index"`
	VideoId uint32
}

type FavoriteListResponse struct {
	Response
	Video
}
