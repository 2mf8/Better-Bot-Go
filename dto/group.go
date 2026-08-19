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

type GroupJoinRequestEvent struct {
	GroupOpenId   string `json:"group_openid,omitempty"`
	JoinRequestId string `json:"join_request_id,omitempty"`
	RiskTips      string `json:"risk_tips,omitempty"` // 安全提示语；可疑消息直接返回 warning_tips；普通消息命中 sec_risk_rules 时返回 top_tips
	UnionOpenId   string `json:"union_openid,omitempty"`
	MemberOpenId  string `json:"member_openid,omitempty"`
	Username      string `json:"username,omitempty"`
	ApplyAt       string `json:"apply_at,omitempty"`
	ApplySource   string `json:"apply_source,omitempty"` // 申请来源：self_apply 主动申请，invited 被邀请
	InvitedBy     string `json:"invited_by,omitempty"`
	Bot           bool   `json:"bot,omitempty"`
	VerifyInfo    *VerifyInfo `json:"verify_info,omitempty"`
	AutoApproved *AutoApproved `json:"auto_approved,omitempty"`
}

type VerifyInfo struct {
	Method        string`json:"method,omitempty"` // 入群验证方式：verify_message / admin_review_qa
	VerifyMessage string`json:"verify_message,omitempty"` // 验证消息内容；仅 auth_type=verify_message 时可能携带
	ReviewQAList []*ReviewQA`json:"review_qa_list,omitempty"` // 问答列表；仅 auth_type=admin_review_qa 时可能携带
}

type ReviewQA struct {
	Question string`json:"question,omitempty"`
	Answer   string`json:"answer,omitempty"`
}

type AutoApproved struct{
	StrategyId string `json:"strategy_id,omitempty"`
}