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

func (f *FeedDao) Fetch(check_time time.Time) ([]*model.Video, error) {
	var flow []*model.Video
	res := config.DB.Preload("Author").Where("create_at < ?", check_time).Order("create_at desc").Limit(30).Find(&flow)
	return flow, res.Error
}

func (f *FeedDao) AddFavorite(feed_list []*model.Video, userId uint32) []*model.Video {
	// 更新videos中的favorite字段
	for _, video := range feed_list {
		var favorite_count, follow_count int64
		// 登录用户已点赞该视频
		if config.DB.Model(&model.Favorite{}).Where("user_id = ? and video_id = ?", userId, video.Id).Count(&favorite_count); favorite_count != 0 {
			video.IsFavorite = true
		}
		// 登录用户已关注该作者
		if config.DB.Model(&model.Relationship{}).Where("from_user_id = ? and to_user_id = ?", userId, video.UserId).Count(&follow_count); follow_count != 0 {
			video.Author.IsFollow = true
		}
	}
	return feed_list
}
