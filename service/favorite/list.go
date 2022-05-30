package favoriteservice

import (
	favoritedao "simple-douyin/dao/favorite"
	userdao "simple-douyin/dao/user"
	"simple-douyin/model"
	"time"
)

func List(userId uint32) ([]*model.Video, error) {
	return newFavoriteListFlow(userId).doList()
}

func (f *FavoriteListFlow) doList() ([]*model.Video, error) {
	var err error
	// 获取点赞视频列表
	if f.videoList, err = f.getFavVideoList(); err != nil {
		return nil, err
	}

	// 每个点赞视频获取视频作者
	for _, video := range videoList {
		video.Author, err = getUserInfo(video.UserId)
		if err != nil {
			return nil, err
		}
	}
	return f.videoList, nil
}

func (f *FavoriteListFlow) getFavVideoList() ([]*model.Video, error) {
	return f.favListDao.GetVideoList(f.userId)
}

func (f *FavoriteListFlow) getUserInfo(userId uint32) (*model.User, error) {
	user, err := f.userInfoDao.GetUserBasicInfo(userId)
	return user, err
}

func (f *FavoriteListFlow) prepareFavVideo() (err error) {

	// 每个点赞视频获取视频作者
	for _, video := range f.videoList {
		video.Author, err = f.getUserInfo(video.UserId)
		if err != nil {
			return err
		}
		video.IsFavorite = true
	}
	return nil
}

func newFavoriteListFlow(userId uint32) *FavoriteListFlow {
	return &FavoriteListFlow{
		favListDao:  favoritedao.NewFavListDaoInstance(),
		userInfoDao: userdao.NewUserInfoDaoInstance(),
		userId:      userId,
	}
}

type FavoriteListFlow struct {
	favListDao  *favoritedao.FavListDao
	userInfoDao *userdao.UserInfoDao
	videoList   []*model.Video
	userId      uint32
	modifyTime  time.Time
}
