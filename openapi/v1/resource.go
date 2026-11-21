package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/2mf8/Better-Bot-Go/dto"
)

const domain = "api.bot.qq.com"
const sandBoxDomain = "sandbox.api.sgroup.qq.com"
const getAppAccessTokenDomain = "https://bots.qq.com"

const scheme = "https"

type uri string

// 目前提供的接口的 uri
const (
	getAppAccessTokenUri uri = "/app/getAppAccessToken"

	privateMessageUri          uri = "/v2/users/{openid}/messages"
	privateStreamMessageUri    uri = "/v2/users/{user_openid}/stream_messages"
	groupMessageUri            uri = "/v2/groups/{group_openid}/messages"
	privateRichMediaMessageUri uri = "/v2/users/{openid}/files"
	groupRichMediaMessageUri   uri = "/v2/groups/{group_openid}/files"
	groupMembersGet            uri = "/v2/groups/{group_openid}/members"
	groupMemberInfoGetUri      uri = "/v2/groups/{group_openid}/members"
	groupMembersRemoveUri      uri = "/v2/groups/{group_openid}/batch_remove_members"
	groupMemberBlackListGetUri uri = "/v2/groups/{group_openid}/member_blacklist"
	groupMemberBlackListSetUri uri = "/v2/groups/{group_openid}/member_blacklist"
	groupBotMessageDelUri      uri = "/v2/groups/{group_openid}/messages/{message_id}"
	privateBotMessageDelUri    uri = "/v2/users/{openid}/messages/{message_id}"

	groupInfoUri                             uri = "/v2/groups/{group_openid}/info"
	groupBotStatusUri                        uri = "/v2/groups/{group_openid}/bot_state"
	groupJoinRequestListUri                  uri = "/v2/groups/{group_openid}/join_request_list"
	groupJoinRequestApproveUri               uri = "/v2/groups/{group_openid}/approval_join_request/{member_openid}"
	groupRestrictChatStatusUri               uri = "/v2/groups/{group_openid}/restrict_chat_setting"
	groupRestrictChatSetUri                  uri = "/v2/groups/{group_openid}/restrict_chat_setting"
	groupJoinApproveStategyGetUri            uri = "/v2/groups/join_approval_strategy"
	groupJoinApproveStategySetUri            uri = "/v2/groups/join_approval_strategy"
	groupJoinApproveStategyUpdUri            uri = "/v2/groups/join_approval_strategy/{strategy_id}"
	groupJoinApproveStategyDelUri            uri = "/v2/groups/join_approval_strategy/{strategy_id}"
	groupJoinApproveStategyExeUri            uri = "/v2/groups/join_approval_strategy/{strategy_id}/execute"
	groupJoinApproveStategyWhitelistUsersUri uri = "/v2/groups/join_approval_strategy/{strategy_id}/whitelist_users"

	menuUri              uri = "/v2/menu"
	panelsGetUri         uri = "/v2/panels?scope={scope}"
	panelsSetUri         uri = "/v2/panels"
	panelGetDetailUri    uri = "/v2/panels/{panel_id}"
	panelSetDetailUri    uri = "/v2/panels/{panel_id}"
	panelDelUri          uri = "/v2/panels/{panel_id}"
	panelTargetUpdateUri uri = "/v2/panels/{panel_id}/target"

	userMeURI          uri = "/users/@me"
	generateUrlLinkUri uri = "/v2/generate_url_link"

	gatewayURI    uri = "/gateway" // nolint
	gatewayBotURI uri = "/gateway/bot"

	interactionsURI = "/interactions/{interaction_id}"

	httpSessionsURI uri = "/gateway/webhook/sessions"
	httpSessionURI  uri = "/gateway/webhook/sessions/{session_id}"
)

// getURL 获取接口地址，会处理沙箱环境判断
func (o *openAPI) getURL(endpoint uri) string {
	d := domain
	if o.sandbox {
		d = sandBoxDomain
	}
	return fmt.Sprintf("%s://%s%s", scheme, d, endpoint)
}

func (o *openAPI) getQQURL(endpoint uri) string {
	d := getAppAccessTokenDomain
	return fmt.Sprintf("%s://%s%s", scheme, d, endpoint)
}

func GetAccessToken(appId string, clientsecret string) *dto.GetAccessTokenResp {
	d := getAppAccessTokenDomain
	url := fmt.Sprintf("%s%s", d, getAppAccessTokenUri)
	req := &dto.GetAccessTokenReq{
		AppID:        appId,
		ClientSecret: clientsecret,
	}
	_resp := &dto.GetAccessTokenResp{}
	b, err := json.Marshal(req)
	if err == nil {
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(b))
		if err == nil {
			defer resp.Body.Close()
			body, _ := io.ReadAll(resp.Body)
			json.Unmarshal(body, _resp)
		}
	}
	return _resp
}
