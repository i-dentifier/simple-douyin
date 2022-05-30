package model

type Favorite struct {
	UserId  uint32
	VideoId uint32
}

type FavoriteListResponse struct {
	Response
	Video
}
