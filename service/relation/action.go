package relationservice

import (
	"errors"
	"fmt"
	relationdao "simple-douyin/dao/relation"
)

func Action(user_id uint32, to_user_id uint32, action_type uint8) (uint32, error) {
	return NewActionFlow(user_id, to_user_id, action_type).DoAction()
}

type ActionFlow struct {
	actionDao   *relationdao.ActionDao
	user_id     uint32
	to_user_id  uint32
	action_type uint8
}

func NewActionFlow(user_id uint32, to_user_id uint32, action_type uint8) *ActionFlow {
	return &ActionFlow{
		actionDao:   relationdao.NewActionDaoInstance(),
		user_id:     user_id,
		to_user_id:  to_user_id,
		action_type: action_type,
	}
}

func (f *ActionFlow) DoAction() (uint32, error) {
	var status uint8 = 1
	//对方是否关注自己
	if ok := f.checkRrelation(); !ok { //对方关注了我
		status = 2
	}

	if f.action_type == 1 { //添加关注
		// 之前没有关注对方则可以添加关注
		if ok := f.checkRelation(); !ok {
			return 0, errors.New(fmt.Sprintf("you have followed this user"))
		}
		return f.actionUser(status) //添加关注
	} else {
		if ok := f.checkRelation(); ok {
			return 0, errors.New(fmt.Sprintf("you haven't followed this user yet"))
		}
		return f.uactionUser(status) //取消关注
	}
}
func (f *ActionFlow) checkRelation() bool {
	// 存在err说明该关注关系不存在
	if err := f.actionDao.IsRelationExisted(f.user_id, f.to_user_id); err != nil {
		return true
	}
	return false
}

func (f *ActionFlow) checkRrelation() bool {
	// 存在err说明对方没有关注我
	if err := f.actionDao.IsRrelationExisted(f.user_id, f.to_user_id); err != nil {
		return true
	}
	return false
}

func (f *ActionFlow) actionUser(status uint8) (uint32, error) {
	return f.actionDao.CreateRelation(f.user_id, f.to_user_id, status)
}
func (f *ActionFlow) uactionUser(status uint8) (uint32, error) {
	return f.actionDao.DeleteRelation(f.user_id, f.to_user_id, status)
}
