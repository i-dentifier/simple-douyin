package relationservice

import (
	relationdao "simple-douyin/dao/relation"
	userdao "simple-douyin/dao/user"
	"simple-douyin/model"
)

func FollowerList(userId uint32) ([]*model.User, error) {
	return newFollowerListFlow(userId).dofollowerList()
}

type FollowerListFlow struct {
	followerListDao *relationdao.FollowerListDao
	userInfoDao     *userdao.UserInfoDao
	followerList    []*model.User
	userId          uint32
}

func newFollowerListFlow(userId uint32) *FollowerListFlow {
	return &FollowerListFlow{
		followerListDao: relationdao.NewFollowerListDaoInstance(),
		userInfoDao:     userdao.NewUserInfoDaoInstance(),
		userId:          userId,
	}
}

func (f *FollowerListFlow) dofollowerList() ([]*model.User, error) {
	var err error
	// 获取关注列表
	if f.followerList, err = f.getFollowerList(); err != nil {
		return nil, err
	}

	return f.followerList, nil
}

func (f *FollowerListFlow) getFollowerList() ([]*model.User, error) {
	return f.followerListDao.GetFollowerList(f.userId)
}
