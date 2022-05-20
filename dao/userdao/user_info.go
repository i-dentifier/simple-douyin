package userdao

import (
	"simple-douyin/common"
	"simple-douyin/config"
)

type UserInfoDao struct {
}

func NewUserInfoDaoInstance() *UserInfoDao {
	return &UserInfoDao{}
}

func (f *UserInfoDao) GetUserInfo(userId uint32) (*common.User, error) {
	var user common.User
	config.DB.Where("id = ?", userId).First(&user)
	return &user, nil
}

// func (f *UserInfoDao) GetFollowStatus(userId)
