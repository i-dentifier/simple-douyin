package publishdao

import (
	"simple-douyin/config"
	"simple-douyin/model"
	"sync"
)

var (
	pubActionOnce sync.Once
	pubActionDao  *PubActionDao
)

type PubActionDao struct {
}

func NewPubActionDaoInstance() *PubActionDao {
	pubActionOnce.Do(func() {
		pubActionDao = &PubActionDao{}
	})
	return pubActionDao
}

func (p *PubActionDao) CreateVideo(video *model.Video) error {
	if result := config.DB.Create(&video); result.Error != nil {
		return result.Error
	}
	return nil
}
