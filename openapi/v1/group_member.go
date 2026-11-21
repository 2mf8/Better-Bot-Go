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
