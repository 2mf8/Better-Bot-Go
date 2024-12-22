package onebot

type Frame struct {
	BotId  string `json:"bot_id,omitempty"`
	Echo   string `json:"echo,omitempty"`
	Ok     bool   `json:"ok,omitempty"`
	Time   int64  `json:"time,omitempty"`
	SelfId string `json:"self_id,omitempty"`
	Data   []byte `json:"data,omitempty"`
}
