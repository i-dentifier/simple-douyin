package feedservice

import (
	feeddao "simple-douyin/dao/feed"
	"simple-douyin/model"
	"strconv"
	"time"
)

type FeedFlow struct {
	feedDao *feeddao.FeedDao
}

func Feed(latestTime string, userId uint32, isLogin bool) ([]*model.Video, error) {
	return NewFeedFlow().FetchVideos(latestTime, userId, isLogin)
}

func NewFeedFlow() (f *FeedFlow) {
	return &FeedFlow{
		feedDao: feeddao.NewFeedDaoInstance(),
	}
}

func (f *FeedFlow) FetchVideos(latestTime string, userId uint32, isLogin bool) ([]*model.Video, error) {
	// 默认为现在时间
	searchTime := time.Now()
	if latestTime != "0" {
		var err error
		// 若给定具体时间戳，将其转换为东八区时间
		latestTime, err := strconv.ParseInt(latestTime, 10, 64)
		if err != nil {
			return nil, err
		}
		t := time.Unix(latestTime, 0).Format("2006-01-02 15:04:05")
		shanghaiZone, _ := time.LoadLocation("Asia/Shanghai")
		searchTime, err = time.ParseInLocation("2006-01-02 15:04:05", t, shanghaiZone)
		if err != nil {
			return nil, err
		}
	}
	feedList, err := f.feedDao.Fetch(searchTime)
	if err != nil {
		return nil, err
	}

	// 若用户已经登录，添加点每个视频的点赞/关注信息
	if isLogin {
		feedList = f.feedDao.AddFavorite(feedList, userId)
	}
	return feedList, nil
}
