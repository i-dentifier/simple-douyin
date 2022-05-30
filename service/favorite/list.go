package favoriteservice

import (
	favoritedao "simple-douyin/dao/favorite"
	"simple-douyin/model"
)

func List(userId uint32) (*model.VideoListResponse, error) {

	// 获取点赞视频列表
	videoList, err := favoritedao.GetVideoList(userId)
	if err != nil {
		return nil, err
	}

	// 每个点赞视频获取视频作者
	for _, video := range videoList {
		video.IsFavorite = true
	}

	// 返回视频列表响应
	response := model.VideoListResponse{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "get favorite video list success",
		},
		VideoList: videoList,
	}

	return &response, nil
}
