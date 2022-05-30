package publishservice

import (
	favoritedao "simple-douyin/dao/favorite"
	"simple-douyin/dao/publish"
	"simple-douyin/dao/user"
	"simple-douyin/model"
)

type PublishListFlow struct {
	userInfoDao  *userdao.UserInfoDao
	pubListDao   *publishdao.PubListDao
	favActionDao *favoritedao.FavActionDao
	videoList    []*model.Video
	author       *model.User // 要访问的用户的author，只查询一次后面用拷贝的方式修改
	loginUserId  uint32
}

func PublishList(loginUserId, otherUserId uint32) ([]*model.Video, error) {
	return newPublishListFlow(loginUserId, otherUserId).do()
}

func newPublishListFlow(loginUserId, otherUserId uint32) *PublishListFlow {
	return &PublishListFlow{userInfoDao: userdao.NewUserInfoDaoInstance(),
		pubListDao:   publishdao.NewPubListDaoInstance(),
		favActionDao: favoritedao.NewFavActionDaoInstance(),
		author:       &model.User{Id: otherUserId},
		loginUserId:  loginUserId,
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
		// 喜好查询
		if p.favActionDao.CheckIsFavorite(p.loginUserId, video.Id) {
			video.IsFavorite = true
		}
		// 加入author信息
		video.Author = p.author
	}
}
