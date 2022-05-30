package favoritedao

import (
	"simple-douyin/config"
	"simple-douyin/model"
	"sync"
)

var (
	favActionOnce sync.Once
	favActionDao  *FavActionDao
)

type FavActionDao struct {
}

func NewFavActionDaoInstance() *FavActionDao {
	favActionOnce.Do(func() {
		favActionDao = &FavActionDao{}
	})
	return favActionDao
}

func (f *FavActionDao) CreateFavorite(userId, videoId uint32) error {
	// 开启事务
	tx := config.DB.Begin()
	// videos 的 favorite_count 字段加 1
	video := model.Video{}
	if res := tx.Where("id = ?", videoId).First(&video); res.Error != nil {
		tx.Rollback()
		return res.Error
	}
	video.FavoriteCount++
	if res := tx.Save(&video); res.Error != nil {
		tx.Rollback()
		return res.Error
	}

	// 增加一条favorite记录
	favorite := model.Favorite{
		UserId:  userId,
		VideoId: videoId,
	}
	if result := tx.Create(favorite); result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	// 提交事务
	tx.Commit()
	return nil
}

func (f *FavActionDao) CancelFavorite(userId, videoId uint32) error {
	tx := config.DB.Begin()
	// videos 的 favorite_count 字段减 1
	video := model.Video{}
	if res := tx.Where("id = ?", videoId).First(&video); res.Error != nil {
		tx.Rollback()
		return res.Error
	}
	video.FavoriteCount--
	if res := tx.Save(&video); res.Error != nil {
		tx.Rollback()
		return res.Error
	}

	// 删除 favorite 字段
	if result := config.DB.Where("user_id = ? AND video_id = ?", userId, videoId).Delete(&model.Favorite{}); result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	// 提交事务
	tx.Commit()
	return nil
}

func (f *FavActionDao) CheckIsFavorite(userId, videoId uint32) bool {
	result := config.DB.Where("user_id = ? AND video_id = ?", userId, videoId).First(&model.Favorite{})
	// 没有错误，找到了
	return result.Error == nil
}
