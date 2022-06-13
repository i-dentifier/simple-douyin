package commentservice

import (
	commentdao "simple-douyin/dao/comment"
	"simple-douyin/model"
)

type CommentListFlow struct {
	commentDao *commentdao.CommentDao
	videoId    int64
}

func GetCommentList(videoId int64) ([]*model.Comment, error) {
	return newCommentListFlow(videoId).getCommentList()
}

func newCommentListFlow(videoId int64) *CommentListFlow {
	return &CommentListFlow{
		commentDao: commentdao.NewCommentDaoInstance(),
		videoId:    videoId,
	}
}

func (f *CommentListFlow) getCommentList() ([]*model.Comment, error) {
	commentList, err := f.commentDao.GetCommentList(f.videoId)
	if err != nil {
		return nil, err
	}
	changeDateFoamat(commentList)
	return commentList, nil
}

func changeDateFoamat(commentList []*model.Comment) {
	for _, comment := range commentList {
		comment.CreateDate = comment.CreateTime.Format("01-02")
	}
}
