// Code generated by lark_sdk_gen. DO NOT EDIT.
/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lark

import (
	"context"
)

// CreateMessageReaction 给指定消息添加指定类型的表情回复（reaction即表情回复, 本说明文档统一用“reaction”代称）。
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/home/develop-a-bot-in-5-minutes/create-an-app)
// - 待添加reaction的消息要真实存在, 不能被撤回
// - 给消息添加reaction, 需要reaction的发送方（机器人或者用户）在消息所在的会话内
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/create
func (r *MessageService) CreateMessageReaction(ctx context.Context, request *CreateMessageReactionReq, options ...MethodOptionFunc) (*CreateMessageReactionResp, *Response, error) {
	if r.cli.mock.mockMessageCreateMessageReaction != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Message#CreateMessageReaction mock enable")
		return r.cli.mock.mockMessageCreateMessageReaction(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Message",
		API:                   "CreateMessageReaction",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/messages/:message_id/reactions",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createMessageReactionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMessageCreateMessageReaction mock MessageCreateMessageReaction method
func (r *Mock) MockMessageCreateMessageReaction(f func(ctx context.Context, request *CreateMessageReactionReq, options ...MethodOptionFunc) (*CreateMessageReactionResp, *Response, error)) {
	r.mockMessageCreateMessageReaction = f
}

// UnMockMessageCreateMessageReaction un-mock MessageCreateMessageReaction method
func (r *Mock) UnMockMessageCreateMessageReaction() {
	r.mockMessageCreateMessageReaction = nil
}

// CreateMessageReactionReq ...
type CreateMessageReactionReq struct {
	MessageID    string                                `path:"message_id" json:"-"`     // 待添加reaction的消息ID, 示例值: "om_a8f2294ba1a38afaac9d"
	ReactionType *CreateMessageReactionReqReactionType `json:"reaction_type,omitempty"` // reaction资源类型
}

// CreateMessageReactionReqReactionType ...
type CreateMessageReactionReqReactionType struct {
	EmojiType string `json:"emoji_type,omitempty"` // emoji类型 [emoji类型列举](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/emojis-introduce), 示例值: "SMILE"
}

// CreateMessageReactionResp ...
type CreateMessageReactionResp struct {
	ReactionID   string                                 `json:"reaction_id,omitempty"`   // reaction资源ID
	Operator     *CreateMessageReactionRespOperator     `json:"operator,omitempty"`      // 添加reaction的操作人
	ActionTime   string                                 `json:"action_time,omitempty"`   // reaction动作的的unix timestamp(单位:ms)
	ReactionType *CreateMessageReactionRespReactionType `json:"reaction_type,omitempty"` // reaction资源类型
}

// CreateMessageReactionRespOperator ...
type CreateMessageReactionRespOperator struct {
	OperatorID   string `json:"operator_id,omitempty"`   // 操作人ID
	OperatorType string `json:"operator_type,omitempty"` // 操作人身份, 用户或应用, 可选值有: `app`: "app", `user`: "user"
}

// CreateMessageReactionRespReactionType ...
type CreateMessageReactionRespReactionType struct {
	EmojiType string `json:"emoji_type,omitempty"` // emoji类型 [emoji类型列举](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/emojis-introduce)
}

// createMessageReactionResp ...
type createMessageReactionResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *CreateMessageReactionResp `json:"data,omitempty"`
}
