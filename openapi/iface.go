package openapi

import (
	"context"
	"time"

	"github.com/2mf8/Better-Bot-Go/dto"
	"github.com/2mf8/Better-Bot-Go/token"
)

// RetractMessageOption 撤回消息可选参数
type RetractMessageOption int

type PanelsCursorOption string

type PanelsLimitOption int

const (
	// RetractMessageOptionHidetip 撤回消息隐藏小灰条可选参数
	RetractMessageOptionHidetip RetractMessageOption = 1
	PanelsCursorOptionDefault   PanelsCursorOption   = ""
	PanelsLimitOptionDefault    PanelsLimitOption    = 10
)

// OpenAPI openapi 完整实现
type OpenAPI interface {
	Base
	WebsocketAPI
	UserAPI
	MessageAPI
	WebhookAPI
	InteractionAPI
	GroupMembersAPI
	MenuApi
	GroupApi
}

// Base 基础能力接口
type Base interface {
	Version() APIVersion
	Setup(token *token.Token, inSandbox bool) OpenAPI
	// WithTimeout 设置请求接口超时时间
	WithTimeout(duration time.Duration) OpenAPI
	// Transport 透传请求，如果 sdk 没有及时跟进新的接口的变更，可以使用该方法进行透传，openapi 实现时可以按需选择是否实现该接口
	Transport(ctx context.Context, method, url string, body interface{}) ([]byte, error)
	// TraceID 返回上一次请求的 trace id
	TraceID() string
}

// WebsocketAPI websocket 接入地址
type WebsocketAPI interface {
	WS(ctx context.Context, params map[string]string, body string) (*dto.WebsocketAP, error)
}

type MenuApi interface {
	GolbalMenuGet(ctx context.Context) (*dto.MenuGetResp, error)
	GolbalMenuSet(ctx context.Context, req *dto.MenuSetReq) (*dto.MenuSetResp, error)
	PanelsGet(ctx context.Context, scope string, option *dto.PanelsGetOption) (*dto.PanelsGetResp, error)
	PanelsCreate(ctx context.Context, req *dto.PanelsSetReq) (*dto.PanelsSetResp, error)
	PanelDetailGet(ctx context.Context, panelId string) (*dto.PanelGetResp, error)
	PanelSet(ctx context.Context, panelId string, req *dto.PanelSetReq) (*dto.PanelSetResp, error)
	PanelDel(ctx context.Context, panelId string) (*dto.PanelDelResp, error)
	PanelTargetSet(ctx context.Context, panelId string, req *dto.PanelTargetSetReq) (*dto.PanelTargetSetResp, error)
}

// UserAPI 用户相关接口
type UserAPI interface {
	Me(ctx context.Context) (*dto.User, error)
	ShareUrl(ctx context.Context, req *dto.ShareLinkReq) (*dto.ShareLinkResp, error)
	UserInfoGet(ctx context.Context, userId string) (*dto.UserInfoGetResp, error)
}

// MessageAPI 消息相关接口
type MessageAPI interface {
	PostGroupMessage(ctx context.Context, groupId string, msg *dto.GroupMessageToCreate) (*dto.GroupMsgResp, error)
	PostC2CMessage(ctx context.Context, userId string, msg *dto.C2CMessageToCreate) (*dto.C2CMsgResp, error)
	PostGroupRichMediaMessage(ctx context.Context, groupId string, msg *dto.GroupRichMediaMessageToCreate) (*dto.RichMediaMsgResp, error)
	PostC2CRichMediaMessage(ctx context.Context, userId string, msg *dto.C2CRichMediaMessageToCreate) (*dto.RichMediaMsgResp, error)
	RetractC2CMessage(ctx context.Context, userID, msgID string, options ...RetractMessageOption) error
	RetractGroupBotMessage(ctx context.Context, groupID, msgID string, options ...RetractMessageOption) error
	RetractGroupMessage(ctx context.Context, groupID, msgID string, options ...RetractMessageOption) (dto.RetractMessageResp, error)
	PostC2CStreamMessage(ctx context.Context, userId string, req *dto.StreamMessageCreate) (*dto.C2CMsgResp, error)
}

type GroupApi interface {
	GroupInfoGet(ctx context.Context, groupId string) (*dto.GroupInfoGetResp, error)
	GroupBotStatusGet(ctx context.Context, groupId string) (*dto.GroupBotStatus, error)
	JoinRequestListGet(ctx context.Context, groupId string, req *dto.JoinRequestGetOption) (*dto.GroupJoinRequestListGetResp, error)
	ApprovalJoinRequest(ctx context.Context, groupId, userId string, req *dto.GroupJoinRequestSetReq) (*dto.GroupJoinRequestSetResp, error)
	MuteStatusGet(ctx context.Context, groupId string) (*dto.MuteStatusGetResp, error)
	MuteStatusSet(ctx context.Context, groupId string, req *dto.MuteSetReq) (*dto.MuteSetResp, error)
	JoinApprovalStrategyGet(ctx context.Context, req *dto.GroupJoinApprovalStrategyGetReq) (*dto.GroupJoinApprovalStrategyGetResp, error)
	JoinApprovalStrategySet(ctx context.Context, req *dto.GroupJoinApprovalStrategySetReq) (*dto.GroupJoinApprovalStrategySetResp, error)
	JoinApprovalStrategyUpdate(ctx context.Context, strategyId string, req *dto.GroupJoinApprovalStrategyUpdReq) (*dto.GroupJoinApprovalStrategyUpdResp, error)
	JoinApprovalStrategyDelete(ctx context.Context, strategyId string) (*dto.GroupJoinApprovalStrategyDelResp, error)
	JoinApprovalStrategyExecute(ctx context.Context, strategyId string) (*dto.GroupJoinApprovalStrategyExeResp, error)
	JoinApprovalStrategyWhitelistUsersSet(ctx context.Context, strategyId string, req *dto.WhiteListUsersSetReq) (*dto.WhiteListUsersSetResp, error)
}

// InteractionAPI 互动接口
type InteractionAPI interface {
	// PutInteraction 更新互动信息
	PutInteraction(ctx context.Context, interactionID string, req *dto.InteractionRespReq) error
}

// WebhookAPI http 事件网关相关接口
type WebhookAPI interface {
	CreateSession(ctx context.Context, identity dto.HTTPIdentity) (*dto.HTTPReady, error)
	CheckSessions(ctx context.Context) ([]*dto.HTTPSession, error)
	SessionList(ctx context.Context) ([]*dto.HTTPSession, error)
	RemoveSession(ctx context.Context, sessionID string) error
}

type GroupMembersAPI interface {
	GroupMembersGet(ctx context.Context, groupId string, cursor string) (*dto.GroupMembersListResp, error)
	GroupMemberInfoGet(ctx context.Context, groupId, userId string) (*dto.GroupMemberInfoGetResp, error)
	GroupMembersRemove(ctx context.Context, groupId string, req *dto.GroupMemberRemoveReq) (*dto.GroupMemberRemoveResp, error)
	GroupBlackListGet(ctx context.Context, groupId string, req *dto.GroupBlackListGetOption) (*dto.GroupBlackListGetResp, error)
	GroupBlackListSet(ctx context.Context, groupId string, req *dto.BlackListUserSetReq) (*dto.BlackListUserSetResp, error)
}
