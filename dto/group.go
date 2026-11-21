package dto

type GroupInfoGetReq struct{}

type GroupInfoGetResp struct {
	GroupOpenid     string   `json:"group_openid,omitempty"`
	GroupName       string   `json:"group_name,omitempty"`
	GroupFingerMemo string   `json:"group_finger_memo,omitempty"`
	GroupClassText  string   `json:"group_class_text,omitempty"`
	GroupTags       []string `json:"group_tags,omitempty"`
	GroupMemberNum  int      `json:"group_member_num,omitempty"`
}

type GroupBotStatus struct {
	MemberOpenid      string `json:"member_openid,omitempty"`
	JoinedAt          string `json:"joined_at,omitempty"`
	AllowProactiveMsg bool   `json:"allow_proactive_msg,omitempty"`
	RecvMsgSetting    string `json:"recv_msg_setting,omitempty"`
	MemberRole        string `json:"member_role,omitempty"`
}

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

type GroupMemberAdd struct {
	Timestamp    int    `json:"timestamp,omitempty"`
	GroupOpenId  string `json:"group_openid,omitempty"`
	MemberOpenId string `json:"member_openid,omitempty"`
}

type GroupMemberRemove struct {
	Timestamp    int    `json:"timestamp,omitempty"`
	GroupOpenId  string `json:"group_openid,omitempty"`
	MemberOpenId string `json:"member_openid,omitempty"`
	UserOpenid   string `json:"user_openid,omitempty"`
}

type GroupJoinRequestEvent struct {
	GroupOpenId   string        `json:"group_openid,omitempty"`
	JoinRequestId string        `json:"join_request_id,omitempty"`
	RiskTips      string        `json:"risk_tips,omitempty"` // 安全提示语；可疑消息直接返回 warning_tips；普通消息命中 sec_risk_rules 时返回 top_tips
	UnionOpenId   string        `json:"union_openid,omitempty"`
	MemberOpenId  string        `json:"member_openid,omitempty"`
	Username      string        `json:"username,omitempty"`
	ApplyAt       string        `json:"apply_at,omitempty"`
	ApplySource   string        `json:"apply_source,omitempty"` // 申请来源：self_apply 主动申请，invited 被邀请
	InvitedBy     string        `json:"invited_by,omitempty"`
	Bot           bool          `json:"bot,omitempty"`
	VerifyInfo    *VerifyInfo   `json:"verify_info,omitempty"`
	AutoApproved  *AutoApproved `json:"auto_approved,omitempty"`
}

type VerifyInfo struct {
	Method        string      `json:"method,omitempty"`         // 入群验证方式：verify_message / admin_review_qa
	VerifyMessage string      `json:"verify_message,omitempty"` // 验证消息内容；仅 auth_type=verify_message 时可能携带
	ReviewQAList  []*ReviewQA `json:"review_qa_list,omitempty"` // 问答列表；仅 auth_type=admin_review_qa 时可能携带
}

type ReviewQA struct {
	Question string `json:"question,omitempty"`
	Answer   string `json:"answer,omitempty"`
}

type AutoApproved struct {
	StrategyId string `json:"strategy_id,omitempty"`
}

type GroupJoinRequestListGetReq struct {
	Cursor string `json:"cursor,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}

type GroupJoinRequestListGetResp struct {
	List       []*JoinRequest `json:"list,omitempty"`
	NextCursor string         `json:"next_cursor,omitempty"`
}

type JoinRequest struct {
	GroupOpenId   string      `json:"group_openid,omitempty"`
	JoinRequestId string      `json:"join_request_id,omitempty"`
	RiskTips      string      `json:"risk_tips,omitempty"` // 安全提示语；可疑消息直接返回 warning_tips；普通消息命中 sec_risk_rules 时返回 top_tips
	UnionOpenId   string      `json:"union_openid,omitempty"`
	MemberOpenId  string      `json:"member_openid,omitempty"`
	Username      string      `json:"username,omitempty"`
	ApplyAt       string      `json:"apply_at,omitempty"`
	ApplySource   string      `json:"apply_source,omitempty"` // 申请来源：self_apply 主动申请，invited 被邀请
	InvitedBy     string      `json:"invited_by,omitempty"`
	Bot           bool        `json:"bot,omitempty"`
	VerifyInfo    *VerifyInfo `json:"verify_info,omitempty"`
}

type GroupJoinRequestSetReq struct {
	GroupOpenId  string `json:"group_openid,omitempty"`
	MemberOpenId string `json:"member_openid,omitempty"`
}

type GroupJoinRequestSetResp struct {
	Op                   string `json:"op,omitempty"`
	JoinRequestId        string `json:"join_request_id,omitempty"`
	RejectReason         string `json:"reject_reason,omitempty"`
	AddToMemberBlackList bool   `json:"add_to_member_blacklist,omitempty"`
}

type MuteStatusGetReq struct{}

type MuteStatusGetResp struct {
	GlobalRule struct {
		Mode          string `json:"mode,omitempty"`
		ScheduleRules []struct {
			TaskID  string `json:"task_id,omitempty"`
			StartAt string `json:"start_at,omitempty"`
			EndAt   string `json:"end_at,omitempty"`
			Enabled bool   `json:"enabled,omitempty"`
		} `json:"schedule_rules,omitempty"`
		RecurringRules []struct {
			TaskID    string `json:"task_id,omitempty"`
			Weekdays  []int  `json:"weekdays,omitempty"`
			StartTime string `json:"start_time,omitempty"`
			EndTime   string `json:"end_time,omitempty"`
			Enabled   bool   `json:"enabled,omitempty"`
		} `json:"recurring_rules,omitempty"`
	} `json:"global_rule,omitempty"`
	Members []struct {
		MemberOpenid string `json:"member_openid,omitempty"`
		MuteExpireAt string `json:"mute_expire_at,omitempty"`
		Username     string `json:"username,omitempty"`
		UnionOpenid  string `json:"union_openid,omitempty"`
	} `json:"members,omitempty"`
}

type MuteSetReq struct {
	Op           string `json:"op,omitempty"`
	MemberOpenId string `json:"member_openid,omitempty"`
	MuteExpireAt string `json:"mute_expire_at,omitempty"`
}

type GroupJoinApprovalStrategyGetReq struct {
	Cursor string `json:"cursor,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}

type GroupJoinApprovalStrategyGetResp struct {
	Strategies []struct {
		StrategyID         string   `json:"strategy_id,omitempty"`
		GroupOpenids       []string `json:"group_openids,omitempty"`
		GroupIds           []uint64 `json:"group_ids,omitempty"`
		WhitelistUserCount int      `json:"whitelist_user_count,omitempty"`
		IsEnable           string   `json:"is_enable,omitempty"`
		ExpireAt           string   `json:"expire_at,omitempty"`
		CreatedAt          string   `json:"created_at,omitempty"`
		UpdatedAt          string   `json:"updated_at,omitempty"`
	} `json:"strategies,omitempty"`
	NextCursor string `json:"next_cursor,omitempty"`
}

type GroupJoinApprovalStrategySetReq struct {
	GroupOpenids []string `json:"group_openids,omitempty"`
	GroupIds     []uint64 `json:"group_ids,omitempty"`
	IsEnable     string   `json:"is_enable,omitempty"`
	ExpireAt     string   `json:"expire_at,omitempty"`
	Remark       string   `json:"remark,omitempty"`
}

type GroupJoinApprovalStrategySetResp struct {
	StrategyID string `json:"strategy_id,omitempty"`
	IsEnable   string `json:"is_enable,omitempty"`
	ExpireAt   string `json:"expire_at,omitempty"`
}

type GroupJoinApprovalStrategyUpdReq struct {
	IsEnable    string       `json:"is_enable,omitempty"`
	ExpireAt    string       `json:"expire_at,omitempty"`
	GroupAction *GroupAction `json:"group_action,omitempty"`
	Remark      string       `json:"remark,omitempty"`
}

type GroupAction struct {
	Op           string   `json:"op,omitempty"`
	GroupOpenids []string `json:"group_openids,omitempty"`
	GroupIds     []uint64 `json:"group_ids,omitempty"`
}

type GroupJoinApprovalStrategyUpdResp struct {
	IsEnable string `json:"is_enable,omitempty"`
	ExpireAt string `json:"expire_at,omitempty"`
}

type GroupJoinApprovalStrategyDelReq struct{}

type GroupJoinApprovalStrategyDelResp struct{}

type GroupJoinApprovalStrategyExeReq struct{}

type GroupJoinApprovalStrategyExeResp struct{}

type WhiteListUsersSetReq struct {
	Op             string   `json:"op,omitempty"`
	WhiteListUsers []string `json:"whitelist_users,omitempty"`
}

type WhiteListUsersSetResp struct {
	StrategyID         string `json:"strategy_id,omitempty"`
	WhitelistUserCount int    `json:"whitelist_user_count,omitempty"`
	UpdatedAt          string `json:"updated_at,omitempty"`
}
