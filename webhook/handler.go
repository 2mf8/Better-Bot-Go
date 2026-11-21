package webhook

import "github.com/2mf8/Better-Bot-Go/dto"

// ReadyHandler 可以处理 ws 的 ready 事件
var ReadyHandler func(bot *BotHeaderInfo, event *dto.WSPayload, data *dto.WSReadyData)

// ErrorNotifyHandler 当 ws 连接发生错误的时候，会回调，方便使用方监控相关错误
// 比如 reconnect invalidSession 等错误，错误可以转换为 bot.Err
var ErrorNotifyHandler func(err error)

// PlainEventHandler 透传handler
var PlainEventHandler func(bot *BotHeaderInfo, event *dto.WSPayload, message []byte) error

// CheckEventHandler 消息前置检测
var CheckEventHandler func(bot *BotHeaderInfo, event *dto.WSPayload, message []byte) bool

// InteractionEventHandler 互动事件 handler
var InteractionEventHandler func(bot *BotHeaderInfo, event *dto.WSPayload, data *dto.WSInteractionData) error

var GroupAtMessageEventHandler func(bot *BotHeaderInfo, event *dto.WSPayload, data *dto.WSGroupATMessageData) error

var GroupMessageEventHandler func(bot *BotHeaderInfo, event *dto.WSPayload, data *dto.WSGroupMessageData) error

var GroupMemberAddHander func(bot *BotHeaderInfo, event *dto.WSPayload, data *dto.WSGroupMemberAddData) error

var GroupMemberRemoveHander func(bot *BotHeaderInfo, event *dto.WSPayload, data *dto.WSGroupMemberRemoveData) error

var C2CMessageEventHandler func(bot *BotHeaderInfo, event *dto.WSPayload, data *dto.WSC2CMessageData) error

var GroupAddRobotEventHandler func(bot *BotHeaderInfo, event *dto.WSPayload, data *dto.WSGroupAddRobotData) error

var GroupDelRobotEventHandler func(bot *BotHeaderInfo, event *dto.WSPayload, data *dto.WSGroupDelRobotData) error

var GroupMsgRejectEventHandler func(bot *BotHeaderInfo, event *dto.WSPayload, data *dto.WSGroupMsgRejectData) error

var GroupMsgReceiveEventHandler func(bot *BotHeaderInfo, event *dto.WSPayload, data *dto.WSGroupMsgReceiveData) error

var FriendAddEventHandler func(bot *BotHeaderInfo, event *dto.WSPayload, data *dto.WSFriendAddData) error

var FriendDelEventHandler func(bot *BotHeaderInfo, event *dto.WSPayload, data *dto.WSFriendDelData) error

var C2CMsgRejectHandler func(bot *BotHeaderInfo, event *dto.WSPayload, data *dto.WSFriendMsgRejectData) error

var C2CMsgReceiveHandler func(bot *BotHeaderInfo, event *dto.WSPayload, data *dto.WSFriendMsgReveiceData) error

func init() {
	// ReadyHandler 可以处理 ws 的 ready 事件
	ReadyHandler = func(bot *BotHeaderInfo, event *dto.WSPayload, data *dto.WSReadyData) {}

	// ErrorNotifyHandler 当 ws 连接发生错误的时候，会回调，方便使用方监控相关错误
	// 比如 reconnect invalidSession 等错误，错误可以转换为 bot.Err
	ErrorNotifyHandler = func(err error) {}

	// PlainEventHandler 透传handler
	PlainEventHandler = func(bot *BotHeaderInfo, event *dto.WSPayload, message []byte) error {
		return nil
	}

	// CheckEventHandler 消息前置检测
	CheckEventHandler = func(bot *BotHeaderInfo, event *dto.WSPayload, message []byte) bool {
		return false
	}

	// InteractionEventHandler 互动事件 handler
	InteractionEventHandler = func(bot *BotHeaderInfo, event *dto.WSPayload, data *dto.WSInteractionData) error {
		return nil
	}

	GroupAtMessageEventHandler = func(bot *BotHeaderInfo, event *dto.WSPayload, data *dto.WSGroupATMessageData) error {
		return nil
	}

	GroupMessageEventHandler = func(bot *BotHeaderInfo, event *dto.WSPayload, data *dto.WSGroupMessageData) error {
		return nil
	}

	C2CMessageEventHandler = func(bot *BotHeaderInfo, event *dto.WSPayload, data *dto.WSC2CMessageData) error {
		return nil
	}

	GroupAddRobotEventHandler = func(bot *BotHeaderInfo, event *dto.WSPayload, data *dto.WSGroupAddRobotData) error {
		return nil
	}

	GroupDelRobotEventHandler = func(bot *BotHeaderInfo, event *dto.WSPayload, data *dto.WSGroupDelRobotData) error {
		return nil
	}

	GroupMsgRejectEventHandler = func(bot *BotHeaderInfo, event *dto.WSPayload, data *dto.WSGroupMsgRejectData) error {
		return nil
	}

	GroupMsgReceiveEventHandler = func(bot *BotHeaderInfo, event *dto.WSPayload, data *dto.WSGroupMsgReceiveData) error {
		return nil
	}

	FriendAddEventHandler = func(bot *BotHeaderInfo, event *dto.WSPayload, data *dto.WSFriendAddData) error {
		return nil
	}

	FriendDelEventHandler = func(bot *BotHeaderInfo, event *dto.WSPayload, data *dto.WSFriendDelData) error {
		return nil
	}

	C2CMsgRejectHandler = func(bot *BotHeaderInfo, event *dto.WSPayload, data *dto.WSFriendMsgRejectData) error {
		return nil
	}

	C2CMsgReceiveHandler = func(bot *BotHeaderInfo, event *dto.WSPayload, data *dto.WSFriendMsgReveiceData) error {
		return nil
	}
}
