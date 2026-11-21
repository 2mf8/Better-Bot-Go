package event

import (
	"encoding/json"

	"github.com/2mf8/Better-Bot-Go/dto"
	"github.com/tidwall/gjson" // 由于回包的 d 类型不确定，gjson 用于从回包json中提取 d 并进行针对性的解析
)

var eventParseFuncMap = map[dto.OPCode]map[dto.EventType]eventParseFunc{
	dto.WSDispatchEvent: {
		dto.EventInteractionCreate: interactionHandler,

		dto.EventGroupATMessageCreate: groupAtMessageHandler,
		dto.EventGroupMessageCreate:   groupMessageHandler,
		dto.EventC2CMessageCreate:     c2cMessageHandle,

		dto.EventGroupAddRobot:   groupAddRobotHandle,
		dto.EventGroupDelRobot:   groupDelRobotHandle,
		dto.EventGroupMsgReject:  groupMsgRejectHandle,
		dto.EventGroupMsgReceive: groupMsgReceiveHandle,

		dto.EventGroupMemberAdd:    groupMemberAddHandle,
		dto.EventGroupMemberRemove: groupMemberReomveHandle,
		dto.EventGroupJoinRequest:  groupJoinRequestHandle,

		dto.EventFriendAdd:     friendAddHandle,
		dto.EventFriendDel:     friendDelHandle,
		dto.EventC2CMsgReject:  c2cMsgRejectHandle,
		dto.EventC2CMsgReceive: c2cMsgReceiveHandle,
	},
}

type eventParseFunc func(event *dto.WSPayload, message []byte) error

// ParseAndHandle 处理回调事件
func ParseAndHandle(payload *dto.WSPayload) error {
	// 指定类型的 handler
	if (DefaultHandlers.Check != nil) && DefaultHandlers.Check(payload, payload.RawMessage) == false {
		return nil
	}
	if h, ok := eventParseFuncMap[payload.OPCode][payload.Type]; ok {
		return h(payload, payload.RawMessage)
	}
	// 透传handler，如果未注册具体类型的 handler，会统一投递到这个 handler
	if DefaultHandlers.Plain != nil {
		return DefaultHandlers.Plain(payload, payload.RawMessage)
	}
	return nil
}

// ParseData 解析数据
func ParseData(message []byte, target interface{}) error {
	data := gjson.Get(string(message), "d")
	err := json.Unmarshal([]byte(data.String()), target)
	return err
}

// ParseData 解析数据
func GetId(message []byte) string {
	data := gjson.Get(string(message), "id")
	return data.String()
}

func interactionHandler(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSInteractionData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	data.ID = GetId(message)
	if DefaultHandlers.Interaction != nil {
		return DefaultHandlers.Interaction(payload, data)
	}
	return nil
}

func groupAtMessageHandler(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSGroupATMessageData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.GroupAtMessage != nil {
		return DefaultHandlers.GroupAtMessage(payload, data)
	}
	return nil
}

func groupMessageHandler(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSGroupMessageData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.GroupMessage != nil {
		return DefaultHandlers.GroupMessage(payload, data)
	}
	return nil
}

func groupMemberAddHandle(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSGroupMemberAddData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.GroupMemberAdd != nil {
		return DefaultHandlers.GroupMemberAdd(payload, data)
	}
	return nil
}

func groupMemberReomveHandle(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSGroupMemberRemoveData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.GroupMemberRemove != nil {
		return DefaultHandlers.GroupMemberRemove(payload, data)
	}
	return nil
}

func groupJoinRequestHandle(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSGroupJoinRequestData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.GroupJoinRequest != nil {
		return DefaultHandlers.GroupJoinRequest(payload, data)
	}
	return nil
}

func c2cMessageHandle(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSC2CMessageData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.C2CMessage != nil {
		return DefaultHandlers.C2CMessage(payload, data)
	}
	return nil
}

func groupAddRobotHandle(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSGroupAddRobotData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.GroupAddRobot != nil {
		return DefaultHandlers.GroupAddRobot(payload, data)
	}
	return nil
}

func groupDelRobotHandle(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSGroupDelRobotData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.GroupDelRobot != nil {
		return DefaultHandlers.GroupDelRobot(payload, data)
	}
	return nil
}

func groupMsgRejectHandle(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSGroupMsgRejectData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.GroupMsgReject != nil {
		return DefaultHandlers.GroupMsgReject(payload, data)
	}
	return nil
}

func groupMsgReceiveHandle(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSGroupMsgReceiveData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.GroupMsgReceive != nil {
		return DefaultHandlers.GroupMsgReceive(payload, data)
	}
	return nil
}

func friendAddHandle(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSFriendAddData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.FriendAdd != nil {
		return DefaultHandlers.FriendAdd(payload, data)
	}
	return nil
}

func friendDelHandle(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSFriendDelData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.FriendDel != nil {
		return DefaultHandlers.FriendDel(payload, data)
	}
	return nil
}

func c2cMsgRejectHandle(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSFriendMsgRejectData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.C2CMsgReject != nil {
		return DefaultHandlers.C2CMsgReject(payload, data)
	}
	return nil
}

func c2cMsgReceiveHandle(payload *dto.WSPayload, message []byte) error {
	data := &dto.WSFriendMsgReveiceData{}
	if err := ParseData(message, data); err != nil {
		return err
	}
	if DefaultHandlers.C2CMsgReceive != nil {
		return DefaultHandlers.C2CMsgReceive(payload, data)
	}
	return nil
}
