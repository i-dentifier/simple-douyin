package publishdao

import (
	"fmt"
	"simple-douyin/config"
	"simple-douyin/model"
	"sync"
)

var (
	pubListOnce sync.Once
	pubListDao  *PubListDao
)

type PubListDao struct {
}

func NewPubListDaoInstance() *PubListDao {
	pubActionOnce.Do(func() {
		pubListDao = &PubListDao{}
	})
	return pubListDao
}

func (p *PubListDao) GetVideoList(userId uint32) ([]*model.Video, error) {
	// result := config.DB.Take()
	var videoList []*model.Video
	result := config.DB.Where("user_id = ?", userId).Find(&videoList)
	fmt.Println(result.RowsAffected)
	return videoList, result.Error
}
