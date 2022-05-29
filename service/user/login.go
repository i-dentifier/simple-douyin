package userservice

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"simple-douyin/dao/user"
	"simple-douyin/model"
)

// LoginFlow 控制登录流程
type LoginFlow struct {
	authDao     *userdao.AuthDao
	logUsername string
	logPassword string
	// userId      uint32
	userAuth *model.UserAuth
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
	err := f.checkUserName()
	// 如果用户不存在返回error
	if err != nil {
		return 0, errors.New(fmt.Sprintf("user:%s doesn't exit", f.logUsername))
	}
	fmt.Printf("%#v\n", f.userAuth)
	// 如果身份验证失败返回error
	if err := f.authentication(); err != nil {
		return 0, errors.New("incorrect user name or password")
	}
	// 将userId和err返回给controller
	return f.userAuth.Id, nil
}

func (f *LoginFlow) checkUserName() error {
	var err error
	f.userAuth, err = f.authDao.FindUser(f.logUsername)
	return err
}

func (f *LoginFlow) authentication() error {
	err := bcrypt.CompareHashAndPassword([]byte(f.userAuth.Password), []byte(f.logPassword))
	return err
}
