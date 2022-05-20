package userservice

import (
	"errors"
	"simple-douyin/dao/userdao"
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
		return 0, errors.New("user doesn't exit")
	}
	// 如果身份验证失败返回error
	if !f.authentication() {
		return 0, errors.New("incorrect password")
	}
	// 将userId和err返回给controller
	return userId, nil
}

func (f *LoginFlow) checkUserName() (uint32, error) {
	userId, err := f.authDao.FindUser(f.logUsername)
	return userId, err
}

func (f *LoginFlow) authentication() bool {
	return f.logPassword == f.authDao.GetPassword()
}
