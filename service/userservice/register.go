package userservice

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"simple-douyin/dao/userdao"
)

func Register(username string, password string) error {
	// 密码加密
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
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

func (f *RegisterFlow) DoRegister() error {
	// 用户名不存在则可以注册
	if f.checkUserName() {
		if err := f.registerUser(); err != nil {
			return err
		}
		return nil
	}
	return errors.New("user already exists")
}

func (f *RegisterFlow) checkUserName() bool {
	// 存在err说明该用户不存在
	if err := f.registerDao.IsUserExisted(f.username); err != nil {
		return true
	}
	return false
}

func (f *RegisterFlow) registerUser() error {
	return f.registerDao.CreateUser(f.username, f.password)
}
