package userdao

import (
	"simple-douyin/common"
	"simple-douyin/config"
)

type AuthDao struct {
	ua common.UserAuth
}

func NewAuthDaoInstance() *AuthDao {
	return &AuthDao{}
}

func (f *AuthDao) FindUser(username string) (uint32, error) {
	res := config.DB.Where("name = ?", username).First(&f.ua)
	return f.ua.Id, res.Error
}

func (f *AuthDao) GetPassword() string {
	return f.ua.Password
}
