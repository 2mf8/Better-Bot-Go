package dto

// Interaction 互动行为对象
type Interaction struct {
	ID            string           `json:"id,omitempty"`             // 互动行为唯一标识
	ApplicationID string           `json:"application_id,omitempty"` // 应用ID
	Type          InteractionType  `json:"type,omitempty"`           // 互动类型
	Data          *InteractionData `json:"data,omitempty"`           // 互动数据
	/* GuildID           string           `json:"guild_id,omitempty"`       // 频道 ID
	ChannelID         string           `json:"channel_id,omitempty"`     // 子频道 ID */
	GroupOpenID       string `json:"group_openid,omitempty"` // 群 ID
	GroupMemberOpenID string `json:"group_member_openid,omitempty"`
	Version           uint32 `json:"version,omitempty"`   //	版本，默认为 1
	ChatType          uint32 `json:"chat_type,omitempty"` // 0=频道, 1=群聊, 2=单聊
	Scene             string `json:"scene,omitempty"`     // c2c=单聊, group=群聊, guild=频道
	Timestamp         string `json:"timestamp,omitempty"`
	UserOpenId        string `json:"user_openid,omitempty"`
}

// InteractionType 互动类型
type InteractionType uint32

const (
	// InteractionTypePing ping
	//InteractionTypePing InteractionType = 1
	// InteractionTypeCommand 命令
	//InteractionTypeCommand InteractionType = 2

	InteractionInlinkKeyboard       = 11
	InteractionCallbackCommand      = 12
	InteractionMessageFeedback      = 13
	InteractionClearSession         = 14
	InteractionInOutStory           = 15
	InteractionSwitchModel          = 16
	InteractionUserAuthorize        = 18
	InteractionGroupAuthorize       = 19
	InteractionGroupAuthorizeStatus = 20
)

// InteractionData 互动数据
type InteractionData struct {
	Name     string              `json:"name,omitempty"`     // 标题
	Type     InteractionDataType `json:"type,omitempty"`     //	数据类型，不同数据类型对应不同的 resolved 数据
	Resolved ResolvedContent     `json:"resolved,omitempty"` // 跟不同的互动类型和数据类型有关系的数据
}

type ResolvedContent struct {
	ButtonData string `json:"button_data,omitempty"`
	ButtonId   string `json:"button_id,omitempty"`
	// UserId     string `json:"user_id,omitempty"` 仅频道场景有值
	FeatureId     string                  `json:"feature_id,omitempty"`
	MessageId     string                  `json:"message_id,omitempty"`
	FeedBackOpt   string                  `json:"feedback_opt,omitempty"`
	Checked       int                     `json:"checked,omitempty"`
	Action        string                  `json:"action,omitempty"`
	MessageScene  InteractionMessageScene `json:"message_scene,omitempty"`
	AuthorizeData AuthorizeData           `json:"authorize_data,omitempty"`
}

type InteractionMessageScene struct {
	Ext []string `json:"ext,omitempty"`
}

type AuthorizeData struct {
	OptScene string `json:"opt_scene,omitempty"`
	Scope    string `json:"scope,omitempty"`
}

// InteractionDataType 互动数据类型
type InteractionDataType uint32

const (
	// InteractionDataTypeChatSearch 聊天框搜索
	InteractionDataTypeChatSearch InteractionDataType = 9
)

type InteractionRespReq struct {
	Code int `json:"code,omitempty"`
}
