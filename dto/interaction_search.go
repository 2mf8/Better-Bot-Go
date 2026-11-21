package dto

// SearchInputResolved 搜索类型的输入数据
type SearchInputResolved struct {
	Keyword string `json:"keyword,omitempty"`
}

// SearchRsp 搜索返回数据
type SearchRsp struct {
	Layouts []SearchLayout `json:"layouts,omitempty"`
}

// SearchLayout 搜索结果的布局
type SearchLayout struct {
	LayoutType LayoutType
	ActionType ActionType
	Title      string
	Records    []SearchRecord
}

// LayoutType 布局类型
type LayoutType uint32

const (
	// LayoutTypeImageText 左图右文
	LayoutTypeImageText LayoutType = 0
)

// ActionType 每行数据的点击行为
type ActionType uint32

const (
	// ActionTypeSendARK 发送 ark 消息
	ActionTypeSendARK ActionType = 0
)

// SearchRecord 每一条搜索结果
type SearchRecord struct {
	Cover string `json:"cover,omitempty"`
	Title string `json:"title,omitempty"`
	Tips  string `json:"tips,omitempty"`
	URL   string `json:"url,omitempty"`
}
