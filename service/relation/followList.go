package relationservice

import (
	relationdao "simple-douyin/dao/relation"
	userdao "simple-douyin/dao/user"
	"simple-douyin/model"
)

func FollowList(userId uint32) ([]*model.User, error) {
	return newFollowListFlow(userId).dofollowList()
}

type FollowListFlow struct {
	followListDao *relationdao.FollowListDao
	userInfoDao   *userdao.UserInfoDao
	followList    []*model.User
	userId        uint32
}

func newFollowListFlow(userId uint32) *FollowListFlow {
	return &FollowListFlow{
		followListDao: relationdao.NewFollowListDaoInstance(),
		userInfoDao:   userdao.NewUserInfoDaoInstance(),
		userId:        userId,
	}
}

func (f *FollowListFlow) dofollowList() ([]*model.User, error) {
	var err error
	// 获取关注列表
	if f.followList, err = f.getFollowList(); err != nil {
		return nil, err
	}

	return f.followList, nil
}

func (f *FollowListFlow) getFollowList() ([]*model.User, error) {
	return f.followListDao.GetFollowList(f.userId)
}
