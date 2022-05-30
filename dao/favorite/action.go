package favoritedao

import (
	"simple-douyin/config"
	"simple-douyin/model"
)

func CreateFavorite(userId, videoId uint32) error {

	// videos 的 favorite_count 字段加 1
	video := model.Video{}
	if res := config.DB.Where("id = ?", videoId).First(&video); res.Error != nil {
		return res.Error
	}
	video.FavoriteCount++
	config.DB.Save(&video)

	// 增加一条favorite记录
	favorite := model.Favorite{
		UserId:  userId,
		VideoId: videoId,
	}
	if result := config.DB.Create(favorite); result.Error != nil {
		return result.Error
	}
	return nil
}

func CancelFavorite(userId, videoId uint32) error {

	// videos 的 favorite_count 字段减 1
	video := model.Video{}
	if res := config.DB.Where("id = ?", videoId).First(&video); res.Error != nil {
		return res.Error
	}
	video.FavoriteCount--
	config.DB.Save(&video)

	// 删除 favorite 字段
	if result := config.DB.Where("user_id = ? AND video_id = ?", userId, videoId).Delete(&model.Favorite{}); result.Error != nil {
		return result.Error
	}
	return nil
}

func CheckIsFavorite(userId, videoId uint32) bool {
	result := config.DB.Where("user_id = ? AND video_id = ?", userId, videoId).First(&model.Favorite{})
	if result.Error == nil {
		// 没有错误，找到了
		return true
	}
	return false
}
