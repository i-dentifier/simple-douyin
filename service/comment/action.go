package commentservice

import (
	"errors"
	commentdao "simple-douyin/dao/comment"
	"simple-douyin/model"
	"strconv"
)

func MakeComment(userId int64, videoId int64, actionType string, commentText string,
	commentId string) (*model.Comment,
	error) {
	// 根据actionType选取对应操作
	if actionType == "1" {
		return newCommentActionFlow(userId, videoId, commentText).makeComment()
	} else if actionType == "2" {

		return newCommentActionFlow(userId, videoId, commentText).deleteComment(commentId)
	}

	return nil, errors.New("actionType error")
}

func newCommentActionFlow(userId int64, videoId int64, commentText string) *CommentActionFlow {
	return &CommentActionFlow{
		commentDao: commentdao.NewCommentDaoInstance(),
		comment: &model.Comment{
			UserId:  userId,
			VideoId: videoId,
			Content: commentText,
		},
	}
}

func (f *CommentActionFlow) makeComment() (*model.Comment, error) {
	comment, err := f.commentDao.CreateComment(f.comment)
	if err != nil {
		return nil, err
	}
	// 更新commentdate格式
	comment.CreateDate = comment.CreateTime.Format("01-02")
	return comment, nil
}

func (f *CommentActionFlow) deleteComment(commentId string) (*model.Comment, error) {
	f.comment.Id, _ = strconv.ParseInt(commentId, 10, 32)
	comment, err := f.commentDao.DeleteComment(f.comment)
	if err != nil {
		return nil, err
	}
	// 更新commentdate格式
	comment.CreateDate = comment.CreateTime.Format("01-02")
	return comment, nil
}

type CommentActionFlow struct {
	commentDao *commentdao.CommentDao
	comment    *model.Comment
}
