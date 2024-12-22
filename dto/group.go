package dto

type GroupAddRobotEvent struct {
	Timestamp      int    `json:"timestamp,omitempty"`
	GroupOpenId    string `json:"group_openid,omitempty"`
	OpMemberOpenId string `json:"op_member_openid,omitempty"`
}

type GroupDelRobotEvent struct {
	Timestamp      int    `json:"timestamp,omitempty"`
	GroupOpenId    string `json:"group_openid,omitempty"`
	OpMemberOpenId string `json:"op_member_openid,omitempty"`
}

type GroupMsgRejectEvent struct {
	Timestamp      int    `json:"timestamp,omitempty"`
	GroupOpenId    string `json:"group_openid,omitempty"`
	OpMemberOpenId string `json:"op_member_openid,omitempty"`
}

type GroupMsgReceiveEvent struct {
	Timestamp      int    `json:"timestamp,omitempty"`
	GroupOpenId    string `json:"group_openid,omitempty"`
	OpMemberOpenId string `json:"op_member_openid,omitempty"`
}
