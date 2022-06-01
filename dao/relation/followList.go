package relationdao

import (
	"simple-douyin/config"
	"simple-douyin/model"
	"sync"
)

var (
	followListOnce sync.Once
	followListDao  *FollowListDao
)

type FollowListDao struct {
}

func NewFollowListDaoInstance() *FollowListDao {
	followListOnce.Do(func() {
		followListDao = &FollowListDao{}
	})
	return followListDao
}

func (f *FollowListDao) GetFollowList(userId uint32) ([]*model.User, error) {
	tx := config.DB.Begin()

	//从relations表中查找关注列表
	var follows []model.Relationship
	if res := tx.Where("from_user_id = ?", userId).Find(&follows); res.Error != nil {
		tx.Rollback()
		return nil, res.Error
	}
	// 关注用户id
	followIds := make([]uint32, 0, len(follows))
	for i := 0; i < len(follows); i++ {
		followIds = append(followIds, follows[i].ToUserId)
	}

	var followList []*model.User

	if res := tx.Where(followIds).Find(&followList); res.Error != nil {
		tx.Rollback()
		return nil, res.Error
	}

	for _, follow := range followList {
		follow.IsFollow = true
	}

	tx.Commit()

	return followList, nil
}
