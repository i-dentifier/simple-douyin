package userservice

import (
	"simple-douyin/dao/user"
	"simple-douyin/model"
)

// QueryUserInfoById 供controller层调用查询用户信息
// userId 将要查询的用户
// tokenUserId 当前发起操作的用户
func QueryUserInfoById(toUserId uint32, fromUserId uint32) (*model.User, error) {
	return NewQueryUserInfoFlow(toUserId, fromUserId).DoQueryUserInfo()
}

type QueryUserInfoFlow struct {
	userInfoDao *userdao.UserInfoDao
	toUserId    uint32
	fromUserId  uint32
}

func NewQueryUserInfoFlow(userId uint32, tokenUserId uint32) *QueryUserInfoFlow {
	return &QueryUserInfoFlow{
		userInfoDao: userdao.NewUserInfoDaoInstance(),
		toUserId:    userId,
		fromUserId:  tokenUserId,
	}
}

func (f *QueryUserInfoFlow) DoQueryUserInfo() (*model.User, error) {
	user, err := f.queryUserBasicInfo()
	if err != nil {
		return nil, err
	}
	// 如果查询的是自己的信息，默认关注自己
	if f.toUserId == f.fromUserId {
		user.IsFollow = true
		return user, nil
	}
	user.IsFollow, err = f.queryIsFollow()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (f *QueryUserInfoFlow) queryUserBasicInfo() (*model.User, error) {
	return f.userInfoDao.GetUserBasicInfo(f.toUserId)
}

func (f *QueryUserInfoFlow) queryIsFollow() (bool, error) {
	err := f.userInfoDao.GetIsFollow(f.toUserId)
	return err == nil, err
}
