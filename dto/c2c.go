package dto

type FriendAddEvent struct {
	Timestamp int    `json:"timestamp,omitempty"`
	OpenId    string `json:"openid"`
}

type FriendDelEvent struct {
	Timestamp int    `json:"timestamp,omitempty"`
	OpenId    string `json:"openid"`
}

type FriendMsgRejectEvent struct {
	Timestamp int    `json:"timestamp,omitempty"`
	OpenId    string `json:"openid"`
}

type FriendMsgReceiveEvent struct {
	Timestamp int    `json:"timestamp,omitempty"`
	OpenId    string `json:"openid"`
}
