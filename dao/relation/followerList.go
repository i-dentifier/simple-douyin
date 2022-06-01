package relationdao

import (
	"simple-douyin/config"
	"simple-douyin/model"
	"sync"
)

var (
	followerListOnce sync.Once
	followerListDao  *FollowerListDao
)

type FollowerListDao struct {
}

func NewFollowerListDaoInstance() *FollowerListDao {
	followerListOnce.Do(func() {
		followerListDao = &FollowerListDao{}
	})
	return followerListDao
}

func (f *FollowerListDao) GetFollowerList(userId uint32) ([]*model.User, error) {
	tx := config.DB.Begin()

	//从relations表中查找关注列表
	var followers []model.Relationship
	if res := tx.Where("to_user_id = ?", userId).Find(&followers); res.Error != nil {
		tx.Rollback()
		return nil, res.Error
	}
	// 关注用户id
	followerIds := make([]uint32, 0, len(followers))
	for i := 0; i < len(followers); i++ {
		followerIds = append(followerIds, followers[i].FromUserId)
	}

	var followerList []*model.User

	if res := tx.Where(followerIds).Find(&followerList); res.Error != nil {
		tx.Rollback()
		return nil, res.Error
	}

	for i, follower := range followerList {
		if followers[i].Status == 2 {
			follower.IsFollow = true
		}
	}

	tx.Commit()

	return followerList, nil
}
