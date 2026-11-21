package dto

type ShareLinkReq struct {
	CallBackData string `json:"callback_data,omitempty"`
}

type ShareLinkResp struct {
	Data *ShareLinkData `json:"data,omitempty"`
}

type ShareLinkData struct {
	Url string `json:"url,omitempty"`
}
