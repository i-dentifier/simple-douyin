package userdao

import (
	"simple-douyin/common"
	"simple-douyin/config"
)

type RegisterDao struct {
}

func NewRegisterDaoInstance() *RegisterDao {
	return &RegisterDao{}
}

func (f *RegisterDao) IsUserExisted(username string) error {
	var u common.User
	res := config.DB.Where("name = ?", username).First(&u)
	return res.Error
}

func (f *RegisterDao) CreateUser(userName string, password string) error {
	// 写入users表
	u := common.User{Name: userName}
	res := config.DB.Create(&u)
	if res.Error != nil {
		return res.Error
	}
	// 级联写入userAuths表
	ua := common.UserAuth{
		Name:     userName,
		Password: password,
	}
	res = config.DB.Create(&ua)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
