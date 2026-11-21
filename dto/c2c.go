package dto

type FriendAddEvent struct {
	Timestamp  int          `json:"timestamp,omitempty"`
	OpenId     string       `json:"openid,omitempty"`
	Scene      int          `json:"scene,omitempty"`
	SceneParam string       `json:"scene_param,omitempty"`
	Author     FriendAuthor `json:"author,omitempty"`
	ShortCode  string       `json:"short_code,omitempty"`
}

type FriendAuthor struct {
	UnionOpenid string `json:"union_openid,omitempty"`
}

type FriendDelEvent struct {
	Timestamp int          `json:"timestamp,omitempty"`
	OpenId    string       `json:"openid,omitempty"`
	Author    FriendAuthor `json:"author,omitempty"`
}

type FriendMsgRejectEvent struct {
	Timestamp int    `json:"timestamp,omitempty"`
	OpenId    string `json:"openid,omitempty"`
}

type FriendMsgReceiveEvent struct {
	Timestamp int    `json:"timestamp,omitempty"`
	OpenId    string `json:"openid,omitempty"`
}
