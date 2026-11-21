package dto

import "github.com/2mf8/Better-Bot-Go/dto/keyboard"

type C2CMsgType uint8

const (
	// C2CMsgTypeText 文本
	C2CMsgTypeText C2CMsgType = iota
	// C2CMsgTypeTextPic 图文混排
	C2CMsgTypeTextPic
	// C2CMsgTypeMarkdown markdown
	C2CMsgTypeMarkdown
	// C2CMsgTypeArk ark
	C2CMsgTypeArk
	// C2CMsgTypeEmbed embed
	C2CMsgTypeEmbed
	// C2CMsgTypeMedia media
	C2CMsgTypeMedia C2CMsgType = 7
	// 图文消息
	C2CTuWen C2CMsgType = 8
)

// GroupMessageToCreate 消息类型： 0 是文本，1 图文混排，2  markdown， 3 ark，4 embed 7 富媒体
type GroupMessageToCreate = C2CMessageToCreate

// C2CMessageToCreate 创建人人、群消息的结构体定义
type C2CMessageToCreate struct {
	Content          string                    `json:"content,omitempty"`
	MsgType          C2CMsgType                `json:"msg_type,omitempty"`
	Markdown         *Markdown                 `json:"markdown,omitempty"`
	Keyboard         *keyboard.MessageKeyboard `json:"keyboard,omitempty"`        // 消息按钮组件
	PromptKeyboard   *PromptKeyboard           `json:"prompt_keyboard,omitempty"` // 消息扩展信息
	Media            *FileInfo                 `json:"media,omitempty"`
	Ark              *Ark                      `json:"ark,omitempty"` // 已失效
	Image            string                    `json:"image,omitempty"`
	MessageReference *MessageReference         `json:"message_reference,omitempty"`
	EventID          EventType                 `json:"event_id,omitempty"` // 要回复的事件id, 逻辑同MsgID
	MsgID            string                    `json:"msg_id,omitempty"`
	MsgSeq           uint                      `json:"msg_seq,omitempty"`
	IsWakeUp         bool                      `json:"is_wakeup,omitempty"`
	Card             *MessageCard              `json:"card,omitempty"`
	ActionButton     *ActionButton             `json:"action_button,omitempty"` // 消息操作结构
	InputNotify      *InputNotify              `json:"input_notify,omitempty"`
}

// ActionButton 消息操作按钮
type ActionButton struct {
	TemplateID   int32  `json:"template_id,omitempty"`   // 消息操作栏模块ID，与下面具体具体按钮二选一填写。待废弃字段！！！
	CallbackData string `json:"callback_data,omitempty"` // 用户操作时会回调通过回调事件给到button_data中， 最长不超过128个字符。
	Feedback     bool   `json:"feedback,omitempty"`      // 反馈按钮（赞踩按钮）
	TTS          bool   `json:"tts,omitempty"`           // TTS语音播放按钮
	ReGenerate   bool   `json:"re_generate,omitempty"`   // 重新生成按钮
	StopGenerate bool   `json:"stop_generate,omitempty"` // 停止生成按钮
}

type PromptKeyboard struct {
	Keyboard *keyboard.MessageKeyboard `json:"keyboard,omitempty"` // 消息按钮组件
}

type MessageCard struct {
	Type    string `json:"type,omitempty"`
	Content struct {
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
		Url         string `json:"url,omitempty"`
		PicUrl      string `json:"pic_url,omitempty"`
	}
}

type InputNotify struct {
	InputType   int `json:"input_type,omitempty"`
	InputSecond int `json:"input_second,omitempty"`
}

type StreamMessageCreate struct {
	InputMode   string `json:"input_mode,omitempty"`
	InputStatus int    `json:"input_status,omitempty"`
	Index       int    `json:"index,omitempty"`
	ContentType string `json:"content_type,omitempty"`
	ContentRaw  string `json:"content_raw,omitempty"`
	EventId     string `json:"event_id,omitempty"`
	MsgId       string `json:"msg_id,omitempty"`
	StreamMsgId string `json:"stream_msg_id,omitempty"`
	MsgReq      int    `json:"msg_req,omitempty"`
	IsWakeUp    bool   `json:"is_wakeup,omitempty"`
}

type FileInfo struct {
	FileInfo string `json:"file_info,omitempty"`
}

// GroupRichMediaMessageToCreate 媒体类型：1 图片，2 视频，3 语音，4 文件（暂不开放） 资源格式要求： 图片：png/jpg，视频：mp4，语音：silk，
type GroupRichMediaMessageToCreate = C2CRichMediaMessageToCreate

// C2CRichMediaMessageToCreate 媒体类型：1 图片，2 视频，3 语音，4 文件（暂不开放） 资源格式要求： 图片：png/jpg，视频：mp4，语音：silk，
type C2CRichMediaMessageToCreate struct {
	FileType   int    `json:"file_type,omitempty"`
	Url        string `json:"url,omitempty"`
	SrvSendMsg bool   `json:"srv_send_msg,omitempty"`
	FileData   []byte `json:"file_data,omitempty"`
}

// MessageReference 引用消息
type MessageReference struct {
	MessageID             string `json:"message_id,omitempty"`               // 消息 id
	IgnoreGetMessageError bool   `json:"ignore_get_message_error,omitempty"` // 是否忽律获取消息失败错误
}

// Markdown markdown 消息
type Markdown struct {
	TemplateID               int    `json:"template_id,omitempty"` // 模版 id
	CustomTemplateId         string `json:"custom_template_id,omitempty"`
	Content                  string `json:"content,omitempty"`                     // 原生 markdown
	ForceVerifyImageResource bool   `json:"force_verify_image_resource,omitempty"` // 模版参数
}

// MarkdownParams markdown 模版参数 键值对
type MarkdownParams struct {
	Key    string   `json:"key,omitempty"`
	Values []string `json:"values,omitempty"`
}

type RichMediaMsgResp struct {
	FileUuid string `json:"file_uuid,omitempty"`
	FileInfo string `json:"file_info,omitempty"`
	Ttl      uint   `json:"ttl,omitempty"`
}

type GroupMsgResp = C2CMsgResp

type C2CMsgResp struct {
	Id        string    `json:"id,omitempty"`
	Timestamp Timestamp `json:"timestamp,omitempty"`
	ExtInfo   struct {
		RefIdx string `json:"ref_idx,omitempty"`
	} `json:"ext_info,omitempty"`
	RemainMsgLen int `json:"remain_msg_len,omitempty"`
}
