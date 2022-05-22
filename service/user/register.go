package userservice

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	userdao "simple-douyin/dao/user"
)

func Register(username string, password string) (uint32, error) {
	// 密码加密
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	encryptPwd := string(hash)
	return NewRegisterFlow(username, encryptPwd).DoRegister()
}

type RegisterFlow struct {
	registerDao *userdao.RegisterDao
	username    string
	password    string
	userId      uint32
	token       string
}

func NewRegisterFlow(username string, password string) *RegisterFlow {
	return &RegisterFlow{
		registerDao: userdao.NewRegisterDaoInstance(),
		username:    username,
		password:    password,
	}
}

func (f *RegisterFlow) DoRegister() (uint32, error) {
	// 用户名不存在则可以注册
	if ok := f.checkUserName(); !ok {
		return 0, errors.New(fmt.Sprintf("user:%s has been registerd", f.username))
	}

	return f.registerUser()
}

func (f *RegisterFlow) checkUserName() bool {
	// 存在err说明该用户不存在
	if err := f.registerDao.IsUserExisted(f.username); err != nil {
		return true
	}
	return false
}

func (f *RegisterFlow) registerUser() (uint32, error) {
	return f.registerDao.CreateUser(f.username, f.password)
}
