package userdao

import (
	"simple-douyin/common"
	"simple-douyin/config"
)

type UserInfoDao struct {
}

func NewUserInfoDaoInstance() *UserInfoDao {
	return &UserInfoDao{}
}

func (f *UserInfoDao) GetUserBasicInfo(userId uint32) (*common.User, error) {
	var user common.User
	res := config.DB.Where("id = ?", userId).First(&user)
	return &user, res.Error
}

func (f *UserInfoDao) GetIsFollow(userId uint32) error {
	var r common.Relationship
	res := config.DB.Where("to_user_id = ?", userId).Find(&r)
	if res.RowsAffected == 0 {
		res = config.DB.Find(&r, "from_user_id = ? AND status = ?", userId, 1)
	}
	return res.Error
}
