package dto

func init() {
	eventIntentMap = transposeIntentEventMap(intentEventMap)
}

// 事件类型
const (
	EventInteractionCreate     EventType = "INTERACTION_CREATE"
	EventC2CMessageCreate      EventType = "C2C_MESSAGE_CREATE"
	EventGroupATMessageCreate  EventType = "GROUP_AT_MESSAGE_CREATE"
	EventGroupMessageCreate    EventType = "GROUP_MESSAGE_CREATE"
	EventGroupAddRobot         EventType = "GROUP_ADD_ROBOT"
	EventGroupDelRobot         EventType = "GROUP_DEL_ROBOT"
	EventGroupMsgReject        EventType = "GROUP_MSG_REJECT"
	EventGroupMsgReceive       EventType = "GROUP_MSG_RECEIVE"
	EventFriendAdd             EventType = "FRIEND_ADD"
	EventFriendDel             EventType = "FRIEND_DEL"
	EventC2CMsgReject          EventType = "C2C_MSG_REJECT"
	EventC2CMsgReceive         EventType = "C2C_MSG_RECEIVE"
	EventGroupMemberAdd        EventType = "GROUP_MEMBER_ADD"
	EventGroupMemberRemove     EventType = "GROUP_MEMBER_REMOVE"
	EventGroupJoinRequest      EventType = "GROUP_JOIN_REQUEST"
)

// intentEventMap 不同 intent 对应的事件定义
var intentEventMap = map[Intent][]EventType{
	IntentInteraction:  {EventInteractionCreate},
	IntentQQ:           {EventC2CMessageCreate, EventC2CMsgReceive, EventC2CMsgReject, EventGroupATMessageCreate, EventGroupMessageCreate, 
		EventGroupAddRobot, EventGroupDelRobot, EventGroupMsgReceive, EventGroupMsgReject, EventFriendAdd, EventFriendDel},
	IntentMemberChange: {EventGroupMemberAdd, EventGroupMemberRemove, 
		EventGroupJoinRequest},
}

var eventIntentMap = transposeIntentEventMap(intentEventMap)

// transposeIntentEventMap 转置 intent 与 event 的关系，用于根据 event 找到 intent
func transposeIntentEventMap(input map[Intent][]EventType) map[EventType]Intent {
	result := make(map[EventType]Intent)
	for i, eventTypes := range input {
		for _, s := range eventTypes {
			result[s] = i
		}
	}
	return result
}

// EventToIntent 事件转换对应的Intent
func EventToIntent(events ...EventType) Intent {
	var i Intent
	for _, event := range events {
		i = i | eventIntentMap[event]
	}
	return i
}
