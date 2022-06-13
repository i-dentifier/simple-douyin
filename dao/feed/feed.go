package feeddao

import (
	"simple-douyin/config"
	"simple-douyin/model"
	"sync"
	"time"
)

var (
	feedOnce sync.Once
	feedDao  *FeedDao
)

type FeedDao struct {
}

func NewFeedDaoInstance() *FeedDao {
	feedOnce.Do(func() {
		feedDao = &FeedDao{}
	})
	return feedDao
}

func (f *FeedDao) Fetch(checkTime time.Time) ([]*model.Video, error) {
	var flow []*model.Video
	res := config.DB.Preload("Author").Where("create_at < ?", checkTime).
		Order("create_at desc").Limit(30).Find(&flow)
	return flow, res.Error
}

func (f *FeedDao) AddFavorite(feedList []*model.Video, userId uint32) []*model.Video {
	// 更新videos中的favorite/comment字段
	for _, video := range feedList {
		var favoriteCount, followCount int64
		// 登录用户已点赞该视频
		if config.DB.Model(&model.Favorite{}).Where("user_id = ? and video_id = ?", userId, video.Id).
			Count(&favoriteCount); favoriteCount != 0 {
			video.IsFavorite = true
		}
		// 登录用户已关注该作者
		if config.DB.Model(&model.Relationship{}).Where(
			"from_user_id = ? and to_user_id = ?", userId, video.UserId).Count(&followCount); followCount != 0 {
			video.Author.IsFollow = true
		}
	}
	return feedList
}
