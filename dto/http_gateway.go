package dto

// HTTPIdentity 鉴权数据
type HTTPIdentity struct {
	Intents  Intent    `json:"intents,omitempty"`
	Shards   [2]uint32 `json:"shards,omitempty"` // array of two integers (shard_id, num_shards)
	Callback string    `json:"callback_url,omitempty"`
}

// HTTPReady ready，鉴权后返回
type HTTPReady struct {
	Version   int    `json:"version,omitempty"`
	SessionID string `json:"session_id,omitempty"`
	Bot       struct {
		ID       string `json:"id,omitempty"`
		Username string `json:"username,omitempty"`
	} `json:"bot,omitempty"`
	Shard [2]uint32 `json:"shard,omitempty"`
}

// HTTPSession session 对象
type HTTPSession struct {
	AppID             int64    `json:"app_id,omitempty"`
	SessionID         string   `json:"session_id,omitempty"`
	CallbackURL       string   `json:"callback_url,omitempty"`
	Env               string   `json:"env,omitempty"`
	Intents           int64    `json:"intents,omitempty"`
	LastHeartbeatTime string   `json:"last_heartbeat_time,omitempty"`
	State             string   `json:"state,omitempty"`
	Shards            [2]int64 `json:"shards,omitempty"`
}
