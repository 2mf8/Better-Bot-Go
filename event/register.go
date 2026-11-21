package event

import (
	"github.com/2mf8/Better-Bot-Go/dto"
)

// DefaultHandlers 默认的 handler 结构，管理所有支持的 handler 类型
var DefaultHandlers struct {
	Ready       ReadyHandler
	ErrorNotify ErrorNotifyHandler
	Plain       PlainEventHandler
	Check       CheckEventHandler

	Interaction InteractionEventHandler

	GroupAtMessage GroupAtMessageEventHandler
	GroupMessage   GroupMessageEventHandler
	C2CMessage     C2CMessageEventHandler

	GroupAddRobot   GroupAddRobotEventHandle
	GroupDelRobot   GroupDelRobotEventHandle
	GroupMsgReject  GroupMsgRejectEventHandle
	GroupMsgReceive GroupMsgReceiveEventHandle

	FriendAdd     FriendAddEventHandle
	FriendDel     FriendDelEventHandle
	C2CMsgReject  C2CMsgRejectHandle
	C2CMsgReceive C2CMsgReceiveHandle

	GroupMemberAdd    GroupMemberAddHandle
	GroupMemberRemove GroupMemberRemoveHandle
	GroupJoinRequest  GroupJoinRequestHandle
}

// ReadyHandler 可以处理 ws 的 ready 事件
type ReadyHandler func(event *dto.WSPayload, data *dto.WSReadyData)

// ErrorNotifyHandler 当 ws 连接发生错误的时候，会回调，方便使用方监控相关错误
// 比如 reconnect invalidSession 等错误，错误可以转换为 bot.Err
type ErrorNotifyHandler func(err error)

// PlainEventHandler 透传handler
type PlainEventHandler func(event *dto.WSPayload, message []byte) error

// CheckEventHandler 消息前置检测
type CheckEventHandler func(event *dto.WSPayload, message []byte) bool

// InteractionEventHandler 互动事件 handler
type InteractionEventHandler func(event *dto.WSPayload, data *dto.WSInteractionData) error

type GroupAtMessageEventHandler func(event *dto.WSPayload, data *dto.WSGroupATMessageData) error

type GroupMessageEventHandler func(event *dto.WSPayload, data *dto.WSGroupMessageData) error

type C2CMessageEventHandler func(event *dto.WSPayload, data *dto.WSC2CMessageData) error

type GroupAddRobotEventHandle func(event *dto.WSPayload, data *dto.WSGroupAddRobotData) error

type GroupDelRobotEventHandle func(event *dto.WSPayload, data *dto.WSGroupDelRobotData) error

type GroupMsgRejectEventHandle func(event *dto.WSPayload, data *dto.WSGroupMsgRejectData) error

type GroupMsgReceiveEventHandle func(event *dto.WSPayload, data *dto.WSGroupMsgReceiveData) error

type FriendAddEventHandle func(event *dto.WSPayload, data *dto.WSFriendAddData) error

type FriendDelEventHandle func(event *dto.WSPayload, data *dto.WSFriendDelData) error

type C2CMsgRejectHandle func(event *dto.WSPayload, data *dto.WSFriendMsgRejectData) error

type C2CMsgReceiveHandle func(event *dto.WSPayload, data *dto.WSFriendMsgReveiceData) error

type GroupMemberAddHandle func(event *dto.WSPayload, data *dto.WSGroupMemberAddData) error

type GroupMemberRemoveHandle func(event *dto.WSPayload, data *dto.WSGroupMemberRemoveData) error

type GroupJoinRequestHandle func(event *dto.WSPayload, data *dto.WSGroupJoinRequestData) error

// RegisterHandlers 注册事件回调，并返回 intent 用于 websocket 的鉴权
func RegisterHandlers(handlers ...interface{}) dto.Intent {
	var i dto.Intent
	for _, h := range handlers {
		switch handle := h.(type) {
		case ReadyHandler:
			DefaultHandlers.Ready = handle
		case ErrorNotifyHandler:
			DefaultHandlers.ErrorNotify = handle
		case CheckEventHandler:
			DefaultHandlers.Check = handle
		case PlainEventHandler:
			DefaultHandlers.Plain = handle
		case GroupAtMessageEventHandler:
			DefaultHandlers.GroupAtMessage = handle
			i = i | dto.EventToIntent(dto.EventGroupATMessageCreate)
		case GroupMessageEventHandler:
			DefaultHandlers.GroupMessage = handle
			i = i | dto.EventToIntent(dto.EventGroupMessageCreate)
		case C2CMessageEventHandler:
			DefaultHandlers.C2CMessage = handle
			i = i | dto.EventToIntent(dto.EventC2CMessageCreate)
		case InteractionEventHandler:
			DefaultHandlers.Interaction = handle
			i = i | dto.EventToIntent(dto.EventInteractionCreate)
		case GroupAddRobotEventHandle:
			DefaultHandlers.GroupAddRobot = handle
			i = i | dto.EventToIntent(dto.EventGroupAddRobot)
		case GroupDelRobotEventHandle:
			DefaultHandlers.GroupDelRobot = handle
			i = i | dto.EventToIntent(dto.EventGroupDelRobot)
		case GroupMsgRejectEventHandle:
			DefaultHandlers.GroupMsgReject = handle
			i = i | dto.EventToIntent(dto.EventGroupMsgReject)
		case GroupMsgReceiveEventHandle:
			DefaultHandlers.GroupMsgReceive = handle
			i = i | dto.EventToIntent(dto.EventGroupMsgReceive)
		case FriendAddEventHandle:
			DefaultHandlers.FriendAdd = handle
			i = i | dto.EventToIntent(dto.EventFriendAdd)
		case FriendDelEventHandle:
			DefaultHandlers.FriendDel = handle
			i = i | dto.EventToIntent(dto.EventFriendDel)
		case C2CMsgRejectHandle:
			DefaultHandlers.C2CMsgReject = handle
			i = i | dto.EventToIntent(dto.EventC2CMsgReject)
		case C2CMsgReceiveHandle:
			DefaultHandlers.C2CMsgReceive = handle
			i = i | dto.EventToIntent(dto.EventC2CMsgReceive)
		case GroupMemberAddHandle:
			DefaultHandlers.GroupMemberAdd = handle
			i = i | dto.EventToIntent(dto.EventGroupMemberAdd)
		case GroupMemberRemoveHandle:
			DefaultHandlers.GroupMemberRemove = handle
			i = i | dto.EventToIntent(dto.EventGroupMemberRemove)
		case GroupJoinRequestHandle:
			DefaultHandlers.GroupJoinRequest = handle
			i = i | dto.EventToIntent(dto.EventGroupJoinRequest)
		default:
		}
	}
	return i
}
