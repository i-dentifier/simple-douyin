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
	ua model.UserAuth
}

func NewAuthDaoInstance() *AuthDao {
	authOnce.Do(func() {
		authDao = &AuthDao{}
	})
	return authDao
}

func (f *AuthDao) FindUser(username string) (uint32, error) {
	res := config.DB.Where("name = ?", username).First(&f.ua)
	return f.ua.Id, res.Error
}

func (f *AuthDao) GetPassword() string {
	return f.ua.Password
}
