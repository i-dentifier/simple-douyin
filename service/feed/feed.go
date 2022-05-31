package feedservice

import (
	feeddao "simple-douyin/dao/feed"
	"simple-douyin/model"
	"time"
)

type FeedFlow struct {
	feedDao *feeddao.FeedDao
}

func Feed(lastTime string, userId uint32, isLogin bool) ([]*model.Video, error) {
	return NewFeedFlow().FetchVideos(lastTime, userId, isLogin)
}

func NewFeedFlow() (f *FeedFlow) {
	return &FeedFlow{
		feedDao: feeddao.NewFeedDaoInstance(),
	}
}

func (f *FeedFlow) FetchVideos(lastTime string, userId uint32, isLogin bool) ([]*model.Video, error) {
	// 默认为现在时间
	search_time := time.Now()
	if lastTime != "" {
		var err error
		// 若给定具体时间，将其转换为东八区时间
		shanghaiZone, _ := time.LoadLocation("Asia/Shanghai")
		search_time, err = time.ParseInLocation("2006-01-02 15:04:05", lastTime, shanghaiZone)
		if err != nil {
			return nil, err
		}
	}
	feed_list, err := f.feedDao.Fetch(search_time)
	if err != nil {
		return nil, err
	}

	// 若用户已经登录，添加点每个视频的点赞/关注信息
	if isLogin {
		feed_list = f.feedDao.AddFavorite(feed_list, userId)
	}
	return feed_list, nil
}
