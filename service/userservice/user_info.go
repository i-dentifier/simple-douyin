package userservice

import (
	"simple-douyin/common"
	"simple-douyin/dao/userdao"
)

func QueryUserInfoById(userId uint32) (*common.User, error) {
	return NewQueryUserInfoFlow(userId).Do()
}

type QueryUserInfoFlow struct {
	userInfoDao *userdao.UserInfoDao
	userId      uint32
}

func NewQueryUserInfoFlow(userId uint32) *QueryUserInfoFlow {
	return &QueryUserInfoFlow{
		userInfoDao: userdao.NewUserInfoDaoInstance(),
		userId:      userId,
	}
}

func (f *QueryUserInfoFlow) Do() (*common.User, error) {
	user, err := f.queryUserBasicInfo()
	if err != nil {
		return nil, err
	}
	// isFollow, err := f.queryIsFollow()
	if err != nil {
		return nil, err
	}
	// user.IsFollow = false
	return user, nil
}

func (f *QueryUserInfoFlow) queryUserBasicInfo() (*common.User, error) {
	return f.userInfoDao.GetUserInfo(f.userId)
}

// func (f *QueryUserInfoFlow) queryIsFollow() (bool, error) {
//
// }
