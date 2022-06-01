package relationdao

import (
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
	r := model.Relationship{}
	res := config.DB.Where("from_user_id = ? and to_user_id = ?", user_id, to_user_id).First(&r)
	return res.Error
}

//判断对方是否关注了自己
func (f *ActionDao) IsRrelationExisted(user_id uint32, to_user_id uint32) error {
	r := model.Relationship{}
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

	if status == 2 { //对方也关注了我，则修改对方关注status
		relation := model.Relationship{}
		if res := tx.Where("from_user_id = ? and to_user_id + ?", to_user_id, user_id).First(&relation); res.Error != nil {
			tx.Rollback()
			return user_id, res.Error
		}
		relation.Status = 2
		if res := tx.Save(&relation); res.Error != nil {
			tx.Rollback()
			return user_id, res.Error
		}
		//tx.Model(&model.Relationship{}).Where("user_id = ?", to_user_id).Update("status", 2)
	}

	//修改users表
	user := model.User{}
	if res := tx.Where("id = ?", user_id).First(&user); res.Error != nil {
		tx.Rollback()
		return user_id, res.Error
	}
	user.FollowCount++
	if res := tx.Save(&user); res.Error != nil {
		tx.Rollback()
		return user_id, res.Error
	}
	to_user := model.User{}
	if res := tx.Where("id = ?", to_user_id).First(&to_user); res.Error != nil {
		tx.Rollback()
		return to_user_id, res.Error
	}
	to_user.FollowerCount++
	if res := tx.Save(&to_user); res.Error != nil {
		tx.Rollback()
		return to_user_id, res.Error
	}

	// 提交事务
	tx.Commit()
	return r.Id, nil

}

func (f *ActionDao) DeleteRelation(user_id uint32, to_user_id uint32, status uint8) (uint32, error) {
	// 开启事务
	tx := config.DB.Begin()

	// 删除relations表中数据
	if res := tx.Where("from_user_id = ? and to_user_id = ?", user_id, to_user_id).
		Delete((&model.Relationship{})); res.Error != nil {
		tx.Rollback()
		return user_id, res.Error
	}

	if status == 2 { //对方也关注了我，则修改对方关注status
		relation := model.Relationship{}
		if res := tx.Where("from_user_id = ? and to_user_id + ?", to_user_id, user_id).First(&relation); res.Error != nil {
			tx.Rollback()
			return user_id, res.Error
		}
		relation.Status = 1
		if res := tx.Save(&relation); res.Error != nil {
			tx.Rollback()
			return user_id, res.Error
		}
	}

	//修改users表
	user := model.User{}
	if res := tx.Where("id = ?", user_id).First(&user); res.Error != nil {
		tx.Rollback()
		return user_id, res.Error
	}
	user.FollowCount--
	if res := tx.Save(&user); res.Error != nil {
		tx.Rollback()
		return user_id, res.Error
	}
	to_user := model.User{}
	if res := tx.Where("id = ?", to_user_id).First(&to_user); res.Error != nil {
		tx.Rollback()
		return to_user_id, res.Error
	}
	to_user.FollowerCount--
	if res := tx.Save(&to_user); res.Error != nil {
		tx.Rollback()
		return to_user_id, res.Error
	}

	// 提交事务
	tx.Commit()
	return user_id, nil

}
