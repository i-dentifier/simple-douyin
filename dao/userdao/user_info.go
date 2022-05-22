package userdao

import (
	"simple-douyin/config"
	"simple-douyin/model"
	"sync"
)

var (
	userInfoOnce sync.Once
	userInfoDao  *UserInfoDao
)

type UserInfoDao struct {
}

func NewUserInfoDaoInstance() *UserInfoDao {
	userInfoOnce.Do(func() {
		userInfoDao = &UserInfoDao{}
	})
	return userInfoDao
}

func (f *UserInfoDao) GetUserBasicInfo(userId uint32) (*model.User, error) {
	var user model.User
	res := config.DB.Where("id = ?", userId).First(&user)
	return &user, res.Error
}

func (f *UserInfoDao) GetIsFollow(userId uint32) error {
	var r model.Relationship
	res := config.DB.Where("to_user_id = ?", userId).Find(&r)
	if res.RowsAffected == 0 {
		res = config.DB.Find(&r, "from_user_id = ? AND status = ?", userId, 1)
	}
	return res.Error
}
