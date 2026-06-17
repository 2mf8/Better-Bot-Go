package dto

type GetGroupMembersReq struct {
	Limit      int `json:"limit,omitempty"`
	StartIndex int `json:"start_index,omitempty"`
}
type GetGroupMembersResp struct {
	Members   []*GroupMember `json:"members,omitempty"`
	NextIndex int            `json:"next_index,omitempty"`
}

type GroupMember struct {
	MemberOpenId  string `json:"member_openid,omitempty"`
	JoinTimeStamp int64  `json:"join_timestamp,omitempty"`
}

type GroupMemberChange struct {
	GroupOpenId  string `json:"group_openid,omitempty"`
	MemberOpenId string `json:"member_openid,omitempty"`
	Timestamp    int64  `json:"timestamp,omitempty"`
}
