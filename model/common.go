package model

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

// Relationship 对应一个关注/粉丝关系
/*type Relationship struct {
	Id         uint64 `json:"id,omitempty"`
	FromUserId uint32 `json:"from_user_id,omitempty"`
	ToUserId   uint32 `json:"to_user_id,omitempty"`
	Status     uint8  `json:"status,omitempty"`
}*/
