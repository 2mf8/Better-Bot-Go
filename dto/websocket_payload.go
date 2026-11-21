package dto

// EventType 事件类型
type EventType string

// WSPayload websocket 消息结构
type WSPayload struct {
	WSPayloadBase
	Data       interface{} `json:"d,omitempty"`
	RawMessage []byte      `json:"-"` // 原始的 message 数据
}

// WSPayloadBase 基础消息结构，排除了 data
type WSPayloadBase struct {
	OPCode OPCode    `json:"op"`
	ID     string    `json:"id,omitempty"`
	Seq    uint32    `json:"s,omitempty"`
	Type   EventType `json:"t,omitempty"`
}

// 以下为发送到 websocket 的 data

// WSIdentityData 鉴权数据
type WSIdentityData struct {
	Token      string   `json:"token"`
	Intents    Intent   `json:"intents"`
	Shard      []uint32 `json:"shard"` // array of two integers (shard_id, num_shards)
	Properties struct {
		Os      string `json:"$os,omitempty"`
		Browser string `json:"$browser,omitempty"`
		Device  string `json:"$device,omitempty"`
	} `json:"properties,omitempty"`
}

// WSResumeData 重连数据
type WSResumeData struct {
	Token     string `json:"token"`
	SessionID string `json:"session_id"`
	Seq       uint32 `json:"seq"`
}

// 以下为会收到的事件data

// WSHelloData hello 返回
type WSHelloData struct {
	HeartbeatInterval int `json:"heartbeat_interval"`
}

// WSReadyData ready，鉴权后返回
type WSReadyData struct {
	Version   int    `json:"version"`
	SessionID string `json:"session_id"`
	User      struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		Bot      bool   `json:"bot"`
	} `json:"user"`
	Shard []uint32 `json:"shard"`
}

type WSC2CMessageData C2CMessage

type WSGroupATMessageData GroupMessage

type WSGroupMessageData GroupMessage

type WSGroupMemberAddData GroupMemberChange

type WSGroupMemberRemoveData GroupMemberChange

// WSInteractionData 互动事件
type WSInteractionData Interaction

type WSGroupAddRobotData GroupAddRobotEvent

type WSGroupDelRobotData GroupDelRobotEvent

type WSGroupMsgRejectData GroupMsgRejectEvent

type WSGroupMsgReceiveData GroupMsgReceiveEvent

type WSFriendAddData FriendAddEvent

type WSFriendDelData FriendDelEvent

type WSFriendMsgRejectData FriendMsgRejectEvent

type WSFriendMsgReveiceData FriendMsgReceiveEvent

type WSGroupJoinRequestData GroupJoinRequestEvent
