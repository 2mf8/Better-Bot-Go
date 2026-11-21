package dto

// Message 消息结构体定义

type GroupMessage struct {
	GroupId     string               `json:"group_id,omitempty"`
	GroupOpenId string               `json:"group_openid,omitempty"`
	Content     string               `json:"content,omitempty"`
	MsgId       string               `json:"id,omitempty"`
	MsgScene    *MsgSceneContent     `json:"message_scene,omitempty"`
	Author      *Author              `json:"author,omitempty"`
	Timestamp   Timestamp            `json:"timestamp,omitempty"`
	MsgType     int                  `json:"message_type,omitempty"`
	Attachments []*MessageAttachment `json:"attachments,omitempty"`
	Mentions    []*User              `json:"mentions,omitempty"`
	ArkData     *ArkData             `json:"ark_data,omitempty"`
	MsgElements []*MsgElement        `json:"msg_elements,omitempty"`
}
type MsgSceneContent struct {
	Source string   `json:"source,omitempty"`
	Ext    []string `json:"ext,omitempty"`
}
type C2CMessage struct {
	Id          string           `json:"id,omitempty"`
	Author      *C2CAuthor       `json:"author,omitempty"`
	Content     string           `json:"content,omitempty"`
	Timestamp   Timestamp        `json:"timestamp,omitempty"`
	MsgScene    *MsgSceneContent `json:"message_scene,omitempty"`
	MsgType     uint32           `json:"msg_type,omitempty"`
	Attachments []*C2CAtachment  `json:"attachments,omitempty"`
	ArkData     *ArkData         `json:"ark_data,omitempty"`
	MsgElements []*MsgElement    `json:"msg_elements,omitempty"`
}

type C2CAtachment struct {
	ContentType  string `json:"content_type,omitempty"`
	Filename     string `json:"filename,omitempty"`
	Height       int    `json:"height,omitempty"`
	Width        int    `json:"width,omitempty"`
	Size         int    `json:"size,omitempty"`
	Url          string `json:"url,omitempty"`
	VoiceWavUrl  string `json:"voice_wav_url,omitempty"`
	AsrReferText string `json:"asr_refer_text,omitempty"`
}

type ArkData struct {
	Prompt  string `json:"prompt,omitempty"`
	ArkType string `json:"ark_type,omitempty"`
	ArkName string `json:"ark_name,omitempty"`
	Fields  struct {
		Tag        string   `json:"tag,omitempty"`
		Tags       []string `json:"tags,omitempty"`
		Title      string   `json:"title,omitempty"`
		Desc       string   `json:"desc,omitempty"`
		JumpUrl    string   `json:"jump_url,omitempty"`
		Preview    string   `json:"preview,omitempty"`
		Source     string   `json:"source,omitempty"`
		SourceLogo string   `json:"source_logo,omitempty"`
		TagIcon    string   `json:"tag_icon,omitempty"`
		NickName   string   `json:"nickname,omitempty"`
		Avatar     string   `json:"avatar,omitempty"`
		Address    string   `json:"address,omitempty"`
	}
}

type MsgElement struct {
	MsgIdx      string          `json:"msg_idx,omitempty"`
	Author      *C2CAuthor      `json:"author,omitempty"`
	MessageType int             `json:"message_type,omitempty"`
	Content     string          `json:"content,omitempty"`
	Attachments []*C2CAtachment `json:"attachments,omitempty"`
	ArkData     *ArkData        `json:"ark_data,omitempty"`
	MsgElements []*MsgElement   `json:"msg_elements,omitempty"`
}

type C2CAuthor struct {
	UserOpenId       string `json:"user_openid,omitempty"`
	Bot              bool   `json:"bot,omitempty"`
	UserId           string `json:"id,omitempty"`
	UnionOpenId      string `json:"union_openid,omitempty"`
	UnionUserAccount string `json:"union_user_account,omitempty"`
	MemberOpenId     string `json:"member_openid,omitempty"`
	Member_role      string `json:"member_role,omitempty"`
	UserName         string `json:"username,omitempty"`
}

type Author struct {
	Bot              bool   `json:"bot,omitempty"`
	UserId           string `json:"id,omitempty"`
	UserOpenId       string `json:"user_openid,omitempty"`
	MemberOpenId     string `json:"member_openid,omitempty"`
	MemberRole       string `json:"member_role,omitempty"`
	UnionOpenId      string `json:"union_openid,omitempty"`
	UnionUserAccount string `json:"union_user_account,omitempty"`
	UserName         string `json:"username,omitempty"`
}

// Embed 结构
type Embed struct {
	Title       string                `json:"title,omitempty"`
	Description string                `json:"description,omitempty"`
	Prompt      string                `json:"prompt,omitempty"` // 消息弹窗内容，消息列表摘要
	Thumbnail   MessageEmbedThumbnail `json:"thumbnail,omitempty"`
	Fields      []*EmbedField         `json:"fields,omitempty"`
}

// MessageEmbedThumbnail embed 消息的缩略图对象
type MessageEmbedThumbnail struct {
	URL string `json:"url,omitempty"`
}

// EmbedField Embed字段描述
type EmbedField struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// MessageAttachment 附件定义
type MessageAttachment struct {
	URL          string `json:"url,omitempty"`
	Filename     string `json:"filename,omitempty"`
	Height       int    `json:"height,omitempty"`
	Width        int    `json:"width,omitempty"`
	Size         int    `json:"size,omitempty"`
	ContentType  string `json:"content_type,omitempty"`
	VoiceWavUrl  string `json:"voice_wav_url,omitempty"`
	AsrReferText string `json:"asr_refer_text,omitempty"`
}

// MessageReactionUsers 消息表情表态用户列表
type MessageReactionUsers struct {
	Users  []*User `json:"users,omitempty"`
	Cookie string  `json:"cookie,omitempty"`
	IsEnd  bool    `json:"is_end,omitempty"`
}

type RetractMessageResp struct {
	Message string `json:"message,omitempty"`
	Code    int64  `json:"code,omitempty"`
	ErrCode int64  `json:"err_code,omitempty"`
	TraceId string `json:"trace_id,omitempty"`
}
