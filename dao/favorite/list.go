package favoritedao

import (
	"simple-douyin/config"
	"simple-douyin/model"
	"sync"
)

var (
	favListOnce sync.Once
	favListDao  *FavListDao
)

type FavListDao struct {
}

func NewFavListDaoInstance() *FavListDao {
	favListOnce.Do(func() {
		favListDao = &FavListDao{}
	})
	return favListDao
}

func (f *FavListDao) GetVideoList(userId uint32) ([]*model.Video, error) {
	var favorites []model.Favorite
	if res := config.DB.Where("user_id = ?", userId).Select("video_id").Find(&favorites); res.Error != nil {
		return nil, res.Error
	}

	// videoIds 该用户所有点赞的视频id
	videoIds := make([]uint32, 0, len(favorites))
	for i := 0; i < len(favorites); i++ {
		videoIds = append(videoIds, favorites[i].VideoId)
	}

	var videoList []*model.Video
	// 添加Preload将video与其发布者关联起来
	config.DB.Preload("Author").Where(videoIds).Find(&videoList) //videoId 主键查询所有视频

	return videoList, nil
}
