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

func (f *FeedDao) Fetch(check_time time.Time) ([]model.Video, error) {
	var flow []model.Video
	res := config.DB.Preload("Author").Where("create_at < ?", check_time).Order("create_at desc").Limit(30).Find(&flow)
	return flow, res.Error
}
