package publishservice

import (
	"simple-douyin/dao/publish"
	"simple-douyin/dao/user"
	"simple-douyin/model"
)

type PublishListFlow struct {
	userInfoDao *userdao.UserInfoDao
	pubListDao  *publishdao.PubListDao
	videoList   []*model.Video
	author      *model.User
}

func PublishList(userId uint32) ([]*model.Video, error) {
	return newPublishListFlow(userId).do()
}

func newPublishListFlow(userId uint32) *PublishListFlow {
	return &PublishListFlow{userInfoDao: userdao.NewUserInfoDaoInstance(),
		pubListDao: publishdao.NewPubListDaoInstance(),
		author:     &model.User{Id: userId},
	}
}

func (p *PublishListFlow) do() ([]*model.Video, error) {
	var err error
	// 查询用户信息
	p.author, err = p.getUserBasicInfo()
	if err != nil {
		return nil, err
	}
	// 查询video列表
	p.videoList, err = p.getvideoList()
	if err != nil {
		return nil, err
	}
	// 组装
	p.prepareVideoList()
	return p.videoList, nil
}

func (p *PublishListFlow) getUserBasicInfo() (*model.User, error) {
	return p.userInfoDao.GetUserBasicInfo(p.author.Id)
}

func (p *PublishListFlow) getvideoList() ([]*model.Video, error) {
	return p.pubListDao.GetVideoList(p.author.Id)
}

func (p *PublishListFlow) prepareVideoList() {
	for _, video := range p.videoList {
		video.Author = p.author
	}
}
