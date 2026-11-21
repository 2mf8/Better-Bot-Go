package v1

import (
	"context"
	"encoding/json"

	"github.com/2mf8/Better-Bot-Go/dto"
	"github.com/2mf8/Better-Bot-Go/openapi"
)

func (o *openAPI) PostGroupMessage(ctx context.Context, groupID string, msg *dto.GroupMessageToCreate) (*dto.GroupMsgResp, error) {
	resp, err := o.request(ctx).
		SetResult(dto.GroupMsgResp{}).
		SetPathParam("group_openid", groupID).
		SetBody(msg).
		Post(o.getURL(groupMessageUri))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*dto.GroupMsgResp), nil
}

func (o *openAPI) PostC2CMessage(ctx context.Context, userId string, msg *dto.C2CMessageToCreate) (*dto.C2CMsgResp, error) {
	resp, err := o.request(ctx).
		SetResult(dto.C2CMsgResp{}).
		SetPathParam("openid", userId).
		SetBody(msg).
		Post(o.getURL(privateMessageUri))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*dto.C2CMsgResp), nil
}

func (o *openAPI) PostC2CRichMediaMessage(ctx context.Context, userId string, msg *dto.C2CRichMediaMessageToCreate) (*dto.RichMediaMsgResp, error) {
	resp, err := o.request(ctx).
		SetResult(dto.RichMediaMsgResp{}).
		SetPathParam("openid", userId).
		SetBody(msg).
		Post(o.getURL(privateRichMediaMessageUri))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*dto.RichMediaMsgResp), nil
}

func (o *openAPI) PostGroupRichMediaMessage(ctx context.Context, groupID string, msg *dto.GroupRichMediaMessageToCreate) (*dto.RichMediaMsgResp, error) {
	resp, err := o.request(ctx).
		SetResult(dto.RichMediaMsgResp{}).
		SetPathParam("group_openid", groupID).
		SetBody(msg).
		Post(o.getURL(groupRichMediaMessageUri))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*dto.RichMediaMsgResp), nil
}

func (o *openAPI) PostC2CStreamMessage(ctx context.Context, userId string, req *dto.StreamMessageCreate) (*dto.C2CMsgResp, error) {
	resp, err := o.request(ctx).
		SetResult(dto.C2CMsgResp{}).
		SetPathParam("user_openid", userId).
		SetBody(req).
		Post(o.getURL(privateStreamMessageUri))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*dto.C2CMsgResp), nil
}

// DelC2CMessage 撤回C2C机器人2分钟内发出的消息
func (o *openAPI) RetractC2CMessage(ctx context.Context,
	userID, msgID string, options ...openapi.RetractMessageOption) error {
	request := o.request(ctx).
		SetPathParam("openid", userID).
		SetPathParam("message_id", string(msgID))
	for _, option := range options {
		if option == openapi.RetractMessageOptionHidetip {
			request = request.SetQueryParam("hidetip", "true")
		}
	}
	_, err := request.Delete(o.getURL(privateBotMessageDelUri))
	return err
}

// DelGroupBotMessage 撤回机器人2分钟内发出的群聊消息
func (o *openAPI) RetractGroupBotMessage(ctx context.Context,
	groupID, msgID string, options ...openapi.RetractMessageOption) error {
	request := o.request(ctx).
		SetPathParam("group_openid", groupID).
		SetPathParam("message_id", string(msgID))
	for _, option := range options {
		if option == openapi.RetractMessageOptionHidetip {
			request = request.SetQueryParam("hidetip", "true")
		}
	}
	_, err := request.Delete(o.getURL(groupBotMessageDelUri))
	return err
}

// DelGroupMessage 撤回群成员发出的群聊消息（机器人需为管理）
func (o *openAPI) RetractGroupMessage(ctx context.Context,
	groupID, msgID string, options ...openapi.RetractMessageOption) (dto.RetractMessageResp, error) {
	dmr := &dto.RetractMessageResp{}
	request := o.request(ctx).
		SetPathParam("group_openid", groupID).
		SetPathParam("message_id", string(msgID))
	for _, option := range options {
		if option == openapi.RetractMessageOptionHidetip {
			request = request.SetQueryParam("hidetip", "true")
		}
	}
	resp, err := request.Delete(o.getURL(groupBotMessageDelUri))
	json.Unmarshal(resp.Body(), &dmr)
	return *dmr, err
}