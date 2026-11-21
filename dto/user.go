package dto

// User 用户
type User struct {
	ID               string `json:"id,omitempty"`
	Username         string `json:"username,omitempty"`
	Avatar           string `json:"avatar,omitempty"`
	Bot              bool   `json:"bot,omitempty"`
	UnionOpenID      string `json:"union_openid,omitempty"`       // 特殊关联应用的 openid
	UnionUserAccount string `json:"union_user_account,omitempty"` // 机器人关联的用户信息，与union_openid关联的应用是同一个
	UserOpenId string `json:"user_openid,omitempty"`
	MemberOpenid      string `json:"member_openid,omitempty"`
	MemberRole        string `json:"member_role,omitempty"`
}
