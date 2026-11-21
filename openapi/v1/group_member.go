package v1

import (
	"context"
	"fmt"
	"strings"

	"github.com/2mf8/Better-Bot-Go/dto"
)

func (o *openAPI) GroupMembersGet(ctx context.Context, groupId string, cursor string) (*dto.GroupMembersListResp, error) {
	url := o.getURL(groupMembersGet)
	if strings.TrimSpace(cursor) != "" {
		url += fmt.Sprintf("?cursor=%s", cursor)
	}
	resp, err := o.request(ctx).
		SetResult(dto.GroupMembersListResp{}).
		SetPathParam("group_openid", groupId).
		Get(url)
	if err != nil {
		return nil, err
	}

	return resp.Result().(*dto.GroupMembersListResp), nil
}

func (o *openAPI) GroupMemberInfoGet(ctx context.Context, groupId, userId string) (*dto.GroupMemberInfoGetResp, error) {
	resp, err := o.request(ctx).
		SetResult(dto.GroupMemberInfoGetResp{}).
		SetPathParam("group_openid", groupId).
		SetPathParam("member_openid", userId).
		Get(o.getURL(groupMemberInfoGetUri))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*dto.GroupMemberInfoGetResp), nil
}

func (o *openAPI) GroupMembersRemove(ctx context.Context, groupId string, req *dto.GroupMemberRemoveReq) (*dto.GroupMemberRemoveResp, error) {
	resp, err := o.request(ctx).
		SetResult(dto.GroupMemberRemoveResp{}).
		SetPathParam("group_openid", groupId).
		SetBody(req).
		Post(o.getURL(groupMembersRemoveUri))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*dto.GroupMemberRemoveResp), nil
}

func (o *openAPI) GroupBlackListGet(ctx context.Context, groupId string, req *dto.GroupBlackListGetOption) (*dto.GroupBlackListGetResp, error) {
	url := o.getURL(groupMemberBlackListGetUri)
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
		SetResult(dto.GroupBlackListGetResp{}).
		SetPathParam("group_openid", groupId).
		Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.GroupBlackListGetResp), nil
}

func (o *openAPI) GroupBlackListSet(ctx context.Context, groupId string, req *dto.BlackListUserSetReq) (*dto.BlackListUserSetResp, error) {
	resp, err := o.request(ctx).
		SetResult(dto.BlackListUserSetResp{}).
		SetPathParam("group_openid", groupId).
		SetBody(req).
		Post(o.getURL(groupMemberBlackListSetUri))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*dto.BlackListUserSetResp), nil
}
