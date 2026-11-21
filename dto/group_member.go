package dto

type GroupMember struct {
	MemberOpenId string `json:"member_openid,omitempty"`
	Username     string `json:"username,omitempty"`
	MemberRole   string `json:"member_role,omitempty"`
	Bot          bool   `json:"bot,omitempty"`
	JoinedAt     int64  `json:"joined_at,omitempty"`
	UnionOpenID  string `json:"union_openid,omitempty"`
}

type GroupMemberChange struct {
	GroupOpenId  string `json:"group_openid,omitempty"`
	MemberOpenId string `json:"member_openid,omitempty"`
	Timestamp    int64  `json:"timestamp,omitempty"`
	UserOpenId   string `json:"user_openid,omitempty"`
}

type GroupMembersListReq struct {
	Cursor string `json:"cursor,omitempty"`
}

type GroupMembersListResp struct {
	Members    []*GroupMember `json:"members,omitempty"`
	NextCursor string         `json:"next_cursor,omitempty"`
}

type GroupMemberInfoGetReq struct{}

type GroupMemberInfoGetResp struct {
	MemberOpenId string `json:"member_openid,omitempty"`
	Username     string `json:"username,omitempty"`
	MemberRole   string `json:"member_role,omitempty"`
	Bot          bool   `json:"bot,omitempty"`
	JoinedAt     string `json:"joined_at,omitempty"`
	UnionOpenId  string `json:"union_openid,omitempty"`
}

type GroupMemberRemoveReq struct {
	MemberOpenids []string `json:"member_openids,omitempty"`
}

type GroupMemberRemoveResp struct {
	RemoveMembersResult string `json:"remove_members_result,omitempty"`
}

type GroupBlackListGetReq struct {
	Cursor string `json:"cursor,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}

type GroupBlackListGetResp struct {
	Users      []BlackListUser `json:"users,omitempty"`
	NextCursor string          `json:"next_cursor,omitempty"`
}

type BlackListUser struct {
	UnionOpenId  string `json:"union_openid,omitempty"`
	MemberOpenId string `json:"member_openid,omitempty"`
	Username     string `json:"username,omitempty"`
	BannedAt     string `json:"banned_at,omitempty"`
	Bot          bool   `json:"bot,omitempty"`
}

type BlackListUserSetReq struct {
	Op            string   `json:"op,omitempty"`
	MemberOpenids []string `json:"member_openids,omitempty"`
}

type BlackListUserSetResp struct {
	FailOpenids []string `json:"fail_openids,omitempty"`
}
