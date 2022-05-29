package feedservice

import (
	feeddao "simple-douyin/dao/feed"
	"simple-douyin/model"
	"time"
)

type FeedFlow struct {
	feedDao *feeddao.FeedDao
}

func Feed(lastTime, token string) ([]model.Video, error) {
	return NewFeedFlow(lastTime, token).FetchVideos(lastTime, token)
}

func NewFeedFlow(lastTime, token string) (f *FeedFlow) {
	return &FeedFlow{
		feedDao: feeddao.NewFeedDaoInstance(),
	}
}

func (f *FeedFlow) FetchVideos(lastTime, token string) ([]model.Video, error) {
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
	feed, err := f.feedDao.Fetch(search_time)
	if err != nil {
		return nil, nil
	}
	return feed, nil
}
