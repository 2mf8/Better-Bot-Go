package dto

// Message 消息结构体定义
type Message struct {
	// 消息ID
	ID string `json:"id"`
	// 子频道ID
	ChannelID string `json:"channel_id"`
	// 频道ID
	GuildID string `json:"guild_id"`
	// 内容
	Content string `json:"content"`
	// 发送时间
	Timestamp Timestamp `json:"timestamp"`
	// 消息编辑时间
	EditedTimestamp Timestamp `json:"edited_timestamp"`
	// 是否@all
	MentionEveryone bool `json:"mention_everyone"`
	// 消息发送方
	Author *User `json:"author"`
	// 消息发送方Author的member属性，只是部分属性
	Member *Member `json:"member"`
	// 附件
	Attachments []*MessageAttachment `json:"attachments"`
	// 结构化消息-embeds
	Embeds []*Embed `json:"embeds"`
	// 消息中的提醒信息(@)列表
	Mentions []*User `json:"mentions"`
	// ark 消息
	Ark *Ark `json:"ark"`
	// 私信消息
	DirectMessage bool `json:"direct_message"`
	// 子频道 seq，用于消息间的排序，seq 在同一子频道中按从先到后的顺序递增，不同的子频道之前消息无法排序
	SeqInChannel string `json:"seq_in_channel"`
	// 引用的消息
	MessageReference *MessageReference `json:"message_reference,omitempty"`
	// 私信场景下，该字段用来标识从哪个频道发起的私信
	SrcGuildID string `json:"src_guild_id"`
}

type GroupMessage struct {
	GroupId     string               `json:"group_id"`
	GroupOpenId string               `json:"group_openid"`
	Content     string               `json:"content"`
	MsgId       string               `json:"id"`
	Author      *Author              `json:"author"`
	Timestamp   Timestamp            `json:"timestamp"`
	Attachments []*MessageAttachment `json:"attachments"`
}

type C2CMessage struct {
	Id          string          `json:"id"`
	Author      *C2CAuthor      `json:"author"`
	Content     string          `json:"content"`
	Timestamp   Timestamp       `json:"timestamp"`
	Attachments []*C2CAtachment `json:"attachments"`
}

type C2CAtachment struct {
	ContentType string `json:"content_type"`
	Filename    string `json:"filename"`
	Height      int    `json:"height"`
	Width       int    `json:"width"`
	Size        int    `json:"size"`
	Url         string `json:"url"`
}

type C2CAuthor struct {
	UserOpenId string `json:"user_openid"`
}

type Author struct {
	UserId     string `json:"id"`
	UserOpenId string `json:"member_openid"`
}

// Embed 结构
type Embed struct {
	Title       string                `json:"title,omitempty"`
	Description string                `json:"description,omitempty"`
	Prompt      string                `json:"prompt"` // 消息弹窗内容，消息列表摘要
	Thumbnail   MessageEmbedThumbnail `json:"thumbnail,omitempty"`
	Fields      []*EmbedField         `json:"fields,omitempty"`
}

// MessageEmbedThumbnail embed 消息的缩略图对象
type MessageEmbedThumbnail struct {
	URL string `json:"url"`
}

// EmbedField Embed字段描述
type EmbedField struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// MessageAttachment 附件定义
type MessageAttachment struct {
	URL string `json:"url"`
}

// MessageReactionUsers 消息表情表态用户列表
type MessageReactionUsers struct {
	Users  []*User `json:"users,omitempty"`
	Cookie string  `json:"cookie,omitempty"`
	IsEnd  bool    `json:"is_end,omitempty"`
}