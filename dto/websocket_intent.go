package dto

// Intent 类型
type Intent int

// websocket intent 声明
const (
	// IntentMemberChange 包含
	// - GROUP_MEMBER_ADD
	// - GROUP_MEMBER_REMOVE
	IntentMemberChange Intent = 1 << 24
	// IntentQQ 包含
	// - C2C_MESSAGE_CREATE
	// - GROUP_AT_MESSAGE_CREATE
	// - GROUP_MESSAGE_CREATE
	// - GROUP_JOIN_REQUEST
	IntentQQ Intent = 1 << 25

	IntentInteraction Intent = 1 << 26 // 互动事件

	IntentNone Intent = 0
)
