package userservice

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	userdao "simple-douyin/dao/user"
)

// LoginFlow 控制登录流程
type LoginFlow struct {
	authDao     *userdao.AuthDao
	logUsername string
	logPassword string
	userId      uint32
}

// Login controller通过调用Login
// 将具体实现交由service处理
func Login(username string, password string) (uint32, error) {
	return NewLoginFlow(username, password).DoLogin()
}

func NewLoginFlow(logUsername string, logPassword string) *LoginFlow {
	return &LoginFlow{
		authDao:     userdao.NewAuthDaoInstance(),
		logUsername: logUsername,
		logPassword: logPassword,
	}
}

func (f *LoginFlow) DoLogin() (uint32, error) {
	userId, err := f.checkUserName()
	// 如果用户不存在返回error
	if err != nil {
		return 0, errors.New(fmt.Sprintf("user:%s doesn't exit", f.logUsername))
	}
	// 如果身份验证失败返回error
	if err := f.authentication(); err != nil {
		return 0, err
	}
	// 将userId和err返回给controller
	return userId, nil
}

func (f *LoginFlow) checkUserName() (uint32, error) {
	userId, err := f.authDao.FindUser(f.logUsername)
	return userId, err
}

func (f *LoginFlow) authentication() error {
	password, err := f.authDao.GetPassword(f.logUsername)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(f.logPassword))
	return err
}
