package openapi

import (
	"context"
	"time"

	"github.com/2mf8/Better-Bot-Go/dto"
	"github.com/2mf8/Better-Bot-Go/token"
)

// RetractMessageOption 撤回消息可选参数
type RetractMessageOption int

const (
	// RetractMessageOptionHidetip 撤回消息隐藏小灰条可选参数
	RetractMessageOptionHidetip RetractMessageOption = 1
)

// OpenAPI openapi 完整实现
type OpenAPI interface {
	Base
	WebsocketAPI
	UserAPI
	MessageAPI
	WebhookAPI
	InteractionAPI
	GetGroupMembersAPI
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

// UserAPI 用户相关接口
type UserAPI interface {
	Me(ctx context.Context) (*dto.User, error)
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
}

// InteractionAPI 互动接口
type InteractionAPI interface {
	// PutInteraction 更新互动信息
	PutInteraction(ctx context.Context, interactionID string, body string) error
}

// WebhookAPI http 事件网关相关接口
type WebhookAPI interface {
	CreateSession(ctx context.Context, identity dto.HTTPIdentity) (*dto.HTTPReady, error)
	CheckSessions(ctx context.Context) ([]*dto.HTTPSession, error)
	SessionList(ctx context.Context) ([]*dto.HTTPSession, error)
	RemoveSession(ctx context.Context, sessionID string) error
}

type GetGroupMembersAPI interface {
	GetGroupMembers(ctx context.Context, groupId string, limit, startIndex int) (*dto.GetGroupMembersResp, error)
}
