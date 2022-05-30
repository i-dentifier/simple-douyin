package favoriteservice

import (
	"errors"
	favoritedao "simple-douyin/dao/favorite"
	"time"
)

func Action(userId, videoId uint32, actionType string) error {

	// actionType == "1" 点赞
	// actionType == "2" 取消点赞
	// actionType == 其它 actionType error

	if actionType == "1" {
		return newFavoriteActionFlow(userId, videoId).doFavorite()
	} else if actionType == "2" {
		return newFavoriteActionFlow(userId, videoId).doCancel()
	} else {
		return errors.New("actionType error")
	}
	return nil
}

func (f *FavoriteActionFlow) doFavorite() error {

	// 检查是否已点赞
	exists := f.favActionDao.CheckIsFavorite(f.userId, f.videoId)
	if exists {
		return errors.New("video has been favorited")
	}

	// 点赞
	err := f.favActionDao.CreateFavorite(f.userId, f.videoId)
	if err != nil {
		return err
	}

	return nil
}

func (f *FavoriteActionFlow) doCancel() error {

	// 检查是否没点赞
	exists := f.favActionDao.CheckIsFavorite(f.userId, f.videoId)
	if !exists {
		return errors.New("video has not been favorited")
	}

	// 取消点赞
	err := f.favActionDao.CancelFavorite(f.userId, f.videoId)
	if err != nil {
		return err
	}

	return nil
}

func newFavoriteActionFlow(userId, videoId uint32) *FavoriteActionFlow {
	return &FavoriteActionFlow{
		favActionDao: favoritedao.NewFavActionDaoInstance(),
		userId:       userId,
		videoId:      videoId,
	}
}

type FavoriteActionFlow struct {
	favActionDao *favoritedao.FavActionDao
	userId       uint32
	videoId      uint32
	modifyTime   time.Time
}
