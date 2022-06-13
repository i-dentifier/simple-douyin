package commentdao

import (
	"errors"
	"gorm.io/gorm"
	"simple-douyin/config"
	"simple-douyin/model"
	"sync"
)

var (
	commentOnce sync.Once
	commentDao  *CommentDao
)

type CommentDao struct {
}

func NewCommentDaoInstance() *CommentDao {
	commentOnce.Do(func() {
		commentDao = &CommentDao{}
	})
	return commentDao
}

func (f *CommentDao) CreateComment(comment *model.Comment) (*model.Comment, error) {
	tx := config.DB.Begin()

	// 插入数据
	if result := tx.Create(&comment); result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	// 添加comment中 User字段
	tx.Preload("User").Where("id = ?", comment.Id).Find(&comment)

	var video model.Video
	// videoId 的 type 待讨论
	video.Id = uint32(comment.VideoId)
	if res := tx.Model(&video).UpdateColumn(
		"comment_count", gorm.Expr("comment_count + ?", 1)); res.Error != nil {
		tx.Rollback()
		return nil, res.Error
	}

	tx.Commit()
	return comment, nil
}

func (f *CommentDao) DeleteComment(comment *model.Comment) (*model.Comment, error) {
	tx := config.DB.Begin()

	// 添加comment中 User 和 comment 字段
	tx.Preload("User").Where("id = ?", comment.Id).Find(&comment)

	// 删除评论 （指定了comment中的id字段）
	if res := tx.Delete(&comment); res.RowsAffected == 0 {
		// 若评论不存在, 结束事务并返回
		tx.Rollback()
		return nil, errors.New("Comment doesn't exist!")
	}

	// 更新videos表中字段
	var video model.Video
	video.Id = uint32(comment.VideoId)
	if res := tx.Model(&video).UpdateColumn(
		"comment_count", gorm.Expr("comment_count - ?", 1)); res.Error != nil {
		tx.Rollback()
		return nil, res.Error
	}

	tx.Commit()
	return comment, nil
}

func (f *CommentDao) GetCommentList(videoId int64) ([]*model.Comment, error) {
	var commentList []*model.Comment
	if res := config.DB.Preload("User").Where("video_id = ? ", videoId).Order("create_time desc").
		Find(&commentList); res.Error != nil {
		return nil, res.Error
	}
	return commentList, nil
}
