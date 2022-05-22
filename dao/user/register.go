package userdao

import (
	"simple-douyin/config"
	"simple-douyin/model"
	"sync"
)

var (
	registerOnce sync.Once
	registerDao  *RegisterDao
)

type RegisterDao struct {
}

func NewRegisterDaoInstance() *RegisterDao {
	registerOnce.Do(func() {
		registerDao = &RegisterDao{}
	})
	return registerDao
}

func (f *RegisterDao) IsUserExisted(username string) error {
	var u model.User
	res := config.DB.Where("name = ?", username).First(&u)
	return res.Error
}

func (f *RegisterDao) CreateUser(userName string, password string) (uint32, error) {
	// 开启事务
	tx := config.DB.Begin()

	// 写入users表
	u := model.User{Name: userName}

	if res := tx.Create(&u); res.Error != nil {
		tx.Rollback()
		return u.Id, res.Error
	}
	// 级联写入userAuths表
	ua := model.UserAuth{
		Name:     userName,
		Password: password,
		Id:       u.Id,
	}

	if res := tx.Create(&ua); res.Error != nil {
		tx.Rollback()
		return ua.Id, res.Error
	}
	// 提交事务
	tx.Commit()
	return u.Id, nil
}
