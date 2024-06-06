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

// GetMessageReactionList 获取指定消息的特定类型表情回复列表（reaction即表情回复, 本文档统一用“reaction”代称）。
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
// - 待获取reaction信息的消息要真实存在, 不能被撤回
// - 获取消息的reaction, 需要request的授权主体（机器人或者用户）在消息所在的会话内
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/list
// new doc: https://open.feishu.cn/document/server-docs/im-v1/message-reaction/list
func (r *MessageService) GetMessageReactionList(ctx context.Context, request *GetMessageReactionListReq, options ...MethodOptionFunc) (*GetMessageReactionListResp, *Response, error) {
	if r.cli.mock.mockMessageGetMessageReactionList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Message#GetMessageReactionList mock enable")
		return r.cli.mock.mockMessageGetMessageReactionList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Message",
		API:                   "GetMessageReactionList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/messages/:message_id/reactions",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getMessageReactionListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMessageGetMessageReactionList mock MessageGetMessageReactionList method
func (r *Mock) MockMessageGetMessageReactionList(f func(ctx context.Context, request *GetMessageReactionListReq, options ...MethodOptionFunc) (*GetMessageReactionListResp, *Response, error)) {
	r.mockMessageGetMessageReactionList = f
}

// UnMockMessageGetMessageReactionList un-mock MessageGetMessageReactionList method
func (r *Mock) UnMockMessageGetMessageReactionList() {
	r.mockMessageGetMessageReactionList = nil
}

// GetMessageReactionListReq ...
type GetMessageReactionListReq struct {
	MessageID    string  `path:"message_id" json:"-"`     // 待获取reaction的消息ID, 详情参见[消息ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/intro#ac79c1c2), 示例值: "om_8964d1b4*2b31383276113"
	ReactionType *string `query:"reaction_type" json:"-"` // 待查询消息reaction的类型[emoji类型列举](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/emojis-introduce), 注意: 不传入该参数, 表示拉取所有类型reaction, 示例值: LAUGH
	PageToken    *string `query:"page_token" json:"-"`    // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: YhljsPiGfUgnVAg9urvRFd-BvSqRL20wMZNAWfa9xXkud6UKCybPuUgQ1vM26dj6
	PageSize     *int64  `query:"page_size" json:"-"`     // 分页大小, 示例值: 10, 最大值: `50`
	UserIDType   *IDType `query:"user_id_type" json:"-"`  // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetMessageReactionListResp ...
type GetMessageReactionListResp struct {
	Items     []*GetMessageReactionListRespItem `json:"items,omitempty"`      // 查询指定reaction_type返回的reaction列表
	HasMore   bool                              `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                            `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
}

// GetMessageReactionListRespItem ...
type GetMessageReactionListRespItem struct {
	ReactionID   string                                      `json:"reaction_id,omitempty"`   // reaction资源ID
	Operator     *GetMessageReactionListRespItemOperator     `json:"operator,omitempty"`      // 添加reaction的操作人
	ActionTime   string                                      `json:"action_time,omitempty"`   // reaction动作的的unix timestamp(单位:ms)
	ReactionType *GetMessageReactionListRespItemReactionType `json:"reaction_type,omitempty"` // reaction资源类型
}

// GetMessageReactionListRespItemOperator ...
type GetMessageReactionListRespItemOperator struct {
	OperatorID   string `json:"operator_id,omitempty"`   // 操作人ID
	OperatorType string `json:"operator_type,omitempty"` // 操作人身份, 用户或应用, 可选值有: app: "app", user: "user"
}

// GetMessageReactionListRespItemReactionType ...
type GetMessageReactionListRespItemReactionType struct {
	EmojiType string `json:"emoji_type,omitempty"` // emoji类型 [emoji类型列举](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/emojis-introduce)
}

// getMessageReactionListResp ...
type getMessageReactionListResp struct {
	Code  int64                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                      `json:"msg,omitempty"`  // 错误描述
	Data  *GetMessageReactionListResp `json:"data,omitempty"`
	Error *ErrorDetail                `json:"error,omitempty"`
}
