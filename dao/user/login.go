package userdao

import (
	"simple-douyin/config"
	"simple-douyin/model"
	"sync"
)

var (
	authOnce sync.Once
	authDao  *AuthDao
)

type AuthDao struct {
}

func NewAuthDaoInstance() *AuthDao {
	authOnce.Do(func() {
		authDao = &AuthDao{}
	})
	return authDao
}

func (f *AuthDao) FindUser(username string) (uint32, error) {
	var ua model.UserAuth
	res := config.DB.Select("id").Where("name = ?", username).First(&ua)
	return ua.Id, res.Error
}

func (f *AuthDao) GetPassword(username string) (string, error) {
	var ua model.UserAuth
	res := config.DB.Select("password").Where("name = ?", username).First(&ua)
	return ua.Password, res.Error
}
