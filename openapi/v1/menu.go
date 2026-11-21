package v1

import (
	"context"
	"fmt"
	"strings"

	"github.com/2mf8/Better-Bot-Go/dto"
)

func (o *openAPI) GolbalMenuGet(ctx context.Context) (*dto.MenuGetResp, error) {
	resp, err := o.request(ctx).
		SetResult(dto.MenuGetResp{}).
		Get(o.getURL(menuUri))

	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.MenuGetResp), nil
}

func (o *openAPI) GolbalMenuSet(ctx context.Context, req *dto.MenuSetReq) (*dto.MenuSetResp, error) {
	resp, err := o.request(ctx).
		SetResult(dto.MenuSetResp{}).
		SetBody(req).
		Put(o.getURL(menuUri))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.MenuSetResp), nil
}

// scope 可为 c2c（单聊）、group（群聊）、channel（文字子频道）、dm（频道私信）
func (o *openAPI) PanelsGet(ctx context.Context, scope string, option *dto.PanelsGetOption) (*dto.PanelsGetResp, error) {
	url := o.getURL(panelsGetUri) + fmt.Sprintf("?scope=%s", scope)
	if strings.TrimSpace(option.Cursor) != "" {
		url += fmt.Sprintf("&cursor=%s", option.Cursor)
	}
	if option.Limit > 0 {
		url += fmt.Sprintf("&limit=%d", option.Limit)
	}
	resp, err := o.request(ctx).
		SetResult(dto.PanelsGetResp{}).
		Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.PanelsGetResp), nil
}

// req 里 scope 必填，panel必填
func (o *openAPI) PanelsCreate(ctx context.Context, req *dto.PanelsSetReq) (*dto.PanelsSetResp, error) {
	resp, err := o.request(ctx).
		SetResult(dto.PanelsGetResp{}).
		SetBody(req).
		Post(o.getURL(panelsSetUri))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.PanelsSetResp), nil
}

func (o *openAPI) PanelDetailGet(ctx context.Context, panelId string) (*dto.PanelGetResp, error) {
	resp, err := o.request(ctx).
		SetResult(dto.PanelGetResp{}).
		SetPathParam("panel_id", panelId).
		Get(o.getURL(panelGetDetailUri))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.PanelGetResp), nil
}

func (o *openAPI) PanelSet(ctx context.Context, panelId string, req *dto.PanelSetReq) (*dto.PanelSetResp, error) {
	resp, err := o.request(ctx).
		SetResult(dto.PanelSetResp{}).
		SetPathParam("panel_id", panelId).
		SetBody(req).
		Put(o.getURL(panelSetDetailUri))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.PanelSetResp), nil
}

func (o *openAPI) PanelDel(ctx context.Context, panelId string) (*dto.PanelDelResp, error) {
	resp, err := o.request(ctx).
		SetResult(dto.PanelDelResp{}).
		SetPathParam("panel_id", panelId).
		Delete(o.getURL(panelDelUri))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.PanelDelResp), nil
}

func (o *openAPI) PanelTargetSet(ctx context.Context, panelId string, req *dto.PanelTargetSetReq) (*dto.PanelTargetSetResp, error) {
	resp, err := o.request(ctx).
		SetResult(dto.PanelTargetSetResp{}).
		SetPathParam("panel_id", panelId).
		SetBody(req).
		Put(o.getURL(panelTargetUpdateUri))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.PanelTargetSetResp), nil
}
