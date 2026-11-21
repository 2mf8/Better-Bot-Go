package dto

type MenuGetResp struct {
	Version int  `json:"version,omitempty"`
	Menu    Menu `json:"menu,omitempty"`
}

type MenuSetReq struct {
	Menu Menu `json:"menu,omitempty"`
}

type MenuSetResp struct {
	Version int `json:"version,omitempty"`
}

type Menu struct {
	Items []*MenuItem `json:"items,omitempty"`
}

type MenuItem struct {
	Name         string      `json:"name,omitempty"`
	Type         string      `json:"type,omitempty"`
	SubMenuItems SubMenuItem `json:"sub_menu_items,omitempty"`
	SendMessage  string      `json:"send_message,omitempty"`
	Link         string      `json:"link,omitempty"`
	Switch       Switch      `json:"switch,omitempty"`
}

type SubMenuItem struct {
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
	SendMessage string `json:"send_message,omitempty"`
	Link        string `json:"link,omitempty"`
}

type Switch struct {
	SwitchId string `json:"switch_id,omitempty"`
	Default  bool   `json:"default,omitempty"`
}

type PanelsGetResp struct {
	Records []struct {
		PanelID    string `json:"panel_id,omitempty"`
		Scope      string `json:"scope,omitempty"`
		TargetType string `json:"target_type,omitempty"`
		Panel      struct {
			Items []struct {
				Type      string `json:"type,omitempty"`
				Name      string `json:"name,omitempty"`
				Desc      string `json:"desc,omitempty"`
				OnlyAdmin bool   `json:"only_admin,omitempty"`
				Link      string `json:"link,omitempty"`
			} `json:"items,omitempty"`
			Remark  string `json:"remark,omitempty"`
			Version int    `json:"version,omitempty"`
		} `json:"panel,omitempty"`
		CreateAt  string `json:"created_at,omitempty"`
		UpdatedAt string `json:"updated_at,omitempty"`
		Version   int    `json:"version,omitempty"`
	} `json:"records,omitempty"`
	NextCursor string `json:"next_cursor,omitempty"`
	IsEnd      bool   `json:"is_end,omitempty"`
}

type PanelsSetReq struct {
	Scope        string   `json:"scope,omitempty"`
	TargetType   string   `json:"target_type,omitempty"`
	UserOpenids  []string `json:"user_openids,omitempty"`
	GroupOpenids []string `json:"group_openids,omitempty"`
	Panel        struct {
		Items []struct {
			Type      string `json:"type,omitempty"`
			Name      string `json:"name,omitempty"`
			Desc      string `json:"desc,omitempty"`
			OnlyAdmin bool   `json:"only_admin,omitempty"`
			Link      string `json:"link,omitempty"`
		} `json:"items,omitempty"`
		Remark  string `json:"remark,omitempty"`
		Version int    `json:"version,omitempty"`
	} `json:"panel,omitempty"`
}

type PanelsSetResp struct {
	PanelId string `json:"panel_id,omitempty"`
}

type PanelGetResp struct {
	PanelID    string `json:"panel_id,omitempty"`
	Scope      string `json:"scope,omitempty"`
	TargetType string `json:"target_type,omitempty"`
	Panel      struct {
		Items []struct {
			Type      string `json:"type,omitempty"`
			Name      string `json:"name,omitempty"`
			Desc      string `json:"desc,omitempty"`
			OnlyAdmin bool   `json:"only_admin,omitempty"`
			Link      string `json:"link,omitempty"`
		} `json:"items,omitempty"`
		Remark  string `json:"remark,omitempty"`
		Version int    `json:"version,omitempty"`
	} `json:"panel,omitempty"`
	CreateAt     string   `json:"created_at,omitempty"`
	UpdatedAt    string   `json:"updated_at,omitempty"`
	Version      int      `json:"version,omitempty"`
	UserOpenids  []string `json:"user_openids,omitempty"`
	GroupOpenids []string `json:"group_openids,omitempty"`
}

type PanelSetReq struct {
	Panel struct {
		Items []struct {
			Type      string `json:"type,omitempty"`
			Name      string `json:"name,omitempty"`
			Desc      string `json:"desc,omitempty"`
			OnlyAdmin bool   `json:"only_admin,omitempty"`
			Link      string `json:"link,omitempty"`
		} `json:"items,omitempty"`
		Remark  string `json:"remark,omitempty"`
		Version int    `json:"version,omitempty"`
	} `json:"panel,omitempty"`
}

type PanelSetResp struct {
	Version int `json:"version,omitempty"`
}

type PanelDelResp struct {
}

type PanelTargetSetReq struct {
	Op           string   `json:"op,omitempty"`
	UserOpenids  []string `json:"user_openids,omitempty"`
	GroupOpenids []string `json:"group_openids,omitempty"`
}

type PanelTargetSetResp struct {
}
