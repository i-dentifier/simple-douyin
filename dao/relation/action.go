package relationdao

import (
	"gorm.io/gorm"
	"simple-douyin/config"
	"simple-douyin/model"
	"sync"
)

var (
	actionOnce sync.Once
	actionDao  *ActionDao
)

type ActionDao struct {
}

func NewActionDaoInstance() *ActionDao {
	actionOnce.Do(func() {
		actionDao = &ActionDao{}
	})
	return actionDao
}

func (f *ActionDao) IsRelationExisted(user_id uint32, to_user_id uint32) error {
	var r model.Relationship
	res := config.DB.Where("from_user_id = ? and to_user_id = ?", user_id, to_user_id).First(&r)
	return res.Error
}

//判断对方是否关注了自己
func (f *ActionDao) IsRrelationExisted(user_id uint32, to_user_id uint32) error {
	var r model.Relationship
	res := config.DB.Where("from_user_id = ? and to_user_id = ?", to_user_id, user_id).First(&r)
	return res.Error
}

func (f *ActionDao) CreateRelation(user_id uint32, to_user_id uint32, status uint8) (uint32, error) {
	// 开启事务
	tx := config.DB.Begin()

	// 写入relations表
	r := model.Relationship{
		FromUserId: user_id,
		ToUserId:   to_user_id,
		Status:     status,
	}

	if res := tx.Create(&r); res.Error != nil {
		tx.Rollback()
		return r.Id, res.Error
	}

	//修改users表
	tx.Model(&model.User{}).Where("id = ?", user_id).Update("follow_count", gorm.Expr("follow_count + 1"))
	tx.Model(&model.User{}).Where("id = ?", to_user_id).Update("follower_count", gorm.Expr("follower_count + 1"))

	if status == 2 { //对方也关注了我，则修改对方关注status
		tx.Model(&model.Relationship{}).Where("user_id = ?", to_user_id).Update("status", 2)
	}
	// 提交事务
	tx.Commit()
	return r.Id, nil

}

func (f *ActionDao) DeleteRelation(user_id uint32, to_user_id uint32, status uint8) (uint32, error) {
	// 开启事务
	tx := config.DB.Begin()

	// 删除relations表中数据
	tx.Where("from_user_id = ? and to_user_id = ?", user_id, to_user_id).Delete((&model.Relationship{}))

	//修改users表
	tx.Model(&model.User{}).Where("id = ?", user_id).Update("follow_count", gorm.Expr("follow_count - 1"))
	tx.Model(&model.User{}).Where("id = ?", to_user_id).Update("follower_count", gorm.Expr("follower_count - 1"))

	if status == 2 { //对方也关注了我，则修改对方关注status
		tx.Model(&model.Relationship{}).Where("user_id = ?", to_user_id).Update("status", 1)
	}

	// 提交事务
	tx.Commit()
	return 111, nil

}
