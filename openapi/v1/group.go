package v1

import (
	"context"
	"fmt"
	"strings"

	"github.com/2mf8/Better-Bot-Go/dto"
)

func (o *openAPI) GroupInfoGet(ctx context.Context, groupId string) (*dto.GroupInfoGetResp, error) {
	resp, err := o.request(ctx).
		SetResult(dto.GroupInfoGetResp{}).
		SetPathParam("group_openid", groupId).
		Get(o.getURL(groupInfoUri))

	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.GroupInfoGetResp), nil
}

func (o *openAPI) GroupBotStatusGet(ctx context.Context, groupId string) (*dto.GroupBotStatus, error) {
	resp, err := o.request(ctx).
		SetResult(dto.GroupBotStatus{}).
		SetPathParam("group_openid", groupId).
		Get(o.getURL(groupBotStatusUri))

	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.GroupBotStatus), nil
}

func (o *openAPI) JoinRequestListGet(ctx context.Context, groupId string, req *dto.JoinRequestGetOption) (*dto.GroupJoinRequestListGetResp, error) {
	url := o.getURL(groupJoinRequestListUri)
	if strings.TrimSpace(req.Cursor) != "" && req.Limit > 0 {
		url += fmt.Sprintf("?cursor=%s&limit=%d", req.Cursor, req.Limit)
	} else {
		if strings.TrimSpace(req.Cursor) != "" {
			url += fmt.Sprintf("?cursor=%s", req.Cursor)
		}
		if req.Limit > 0 {
			url += fmt.Sprintf("?limit=%d", req.Limit)
		}
	}
	resp, err := o.request(ctx).
		SetResult(dto.GroupJoinRequestListGetResp{}).
		SetPathParam("group_openid", groupId).
		Get(url)

	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.GroupJoinRequestListGetResp), nil
}

func (o *openAPI) ApprovalJoinRequest(ctx context.Context, groupId, userId string, req *dto.GroupJoinRequestSetReq) (*dto.GroupJoinRequestSetResp, error) {
	resp, err := o.request(ctx).
		SetResult(dto.GroupJoinRequestSetResp{}).
		SetPathParam("group_openid", groupId).
		SetPathParam("member_openid", userId).
		SetBody(req).
		Post(o.getURL(groupJoinRequestApproveUri))

	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.GroupJoinRequestSetResp), nil
}

func (o *openAPI) MuteStatusGet(ctx context.Context, groupId string) (*dto.MuteStatusGetResp, error) {
	resp, err := o.request(ctx).
		SetResult(dto.MuteStatusGetResp{}).
		SetPathParam("group_openid", groupId).
		Get(o.getURL(groupRestrictChatStatusUri))

	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.MuteStatusGetResp), nil
}

func (o *openAPI) MuteStatusSet(ctx context.Context, groupId string, req *dto.MuteSetReq) (*dto.MuteSetResp, error) {
	resp, err := o.request(ctx).
		SetResult(dto.MuteSetResp{}).
		SetPathParam("group_openid", groupId).
		SetBody(req).
		Post(o.getURL(groupRestrictChatSetUri))

	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.MuteSetResp), nil
}
