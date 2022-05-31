package model

// Relationship 对应一个关注/粉丝关系
type Relationship struct {
	Id         uint32 `json:"id,omitempty"`
	FromUserId uint32 `json:"from_user_id,omitempty"`
	ToUserId   uint32 `json:"to_user_id,omitempty"`
	Status     uint8  `json:"status,omitempty"`
}

type ActionResponse struct {
	Response
	RelationId uint32 `json:"relation_id"`
}

type FollowListResponse struct {
	Response
	FollowList []*User `json:"user_list"`
}

type FollowerListResponse struct {
	Response
	FollowerList []*User `json:"user_list"`
}
