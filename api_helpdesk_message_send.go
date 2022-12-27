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

// SendHelpdeskMessage 通过服务台机器人给指定用户的服务台专属群或私聊发送消息, 支持文本、富文本、卡片、图片。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/bot-message/create
func (r *HelpdeskService) SendHelpdeskMessage(ctx context.Context, request *SendHelpdeskMessageReq, options ...MethodOptionFunc) (*SendHelpdeskMessageResp, *Response, error) {
	if r.cli.mock.mockHelpdeskSendHelpdeskMessage != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#SendHelpdeskMessage mock enable")
		return r.cli.mock.mockHelpdeskSendHelpdeskMessage(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Helpdesk",
		API:                   "SendHelpdeskMessage",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/helpdesk/v1/message",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedHelpdeskAuth:      true,
	}
	resp := new(sendHelpdeskMessageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHelpdeskSendHelpdeskMessage mock HelpdeskSendHelpdeskMessage method
func (r *Mock) MockHelpdeskSendHelpdeskMessage(f func(ctx context.Context, request *SendHelpdeskMessageReq, options ...MethodOptionFunc) (*SendHelpdeskMessageResp, *Response, error)) {
	r.mockHelpdeskSendHelpdeskMessage = f
}

// UnMockHelpdeskSendHelpdeskMessage un-mock HelpdeskSendHelpdeskMessage method
func (r *Mock) UnMockHelpdeskSendHelpdeskMessage() {
	r.mockHelpdeskSendHelpdeskMessage = nil
}

// SendHelpdeskMessageReq ...
type SendHelpdeskMessageReq struct {
	UserIDType  *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	MsgType     MsgType `json:"msg_type,omitempty"`     // 消息类型, 示例值: "post", 可选值有: text: 普通文本, post: 富文本, image: 图片, interactive: 卡片消息
	Content     string  `json:"content,omitempty"`      // 消息内容, json格式结构序列化成string。格式说明参考: [发送消息content说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json), 示例值: "{\"post\":{\"zh_cn\":{\"title\":\"some title\", \"content\":[[{\"tag\":\"text\", \"text\":\"some content\"}]]}}}"
	ReceiverID  string  `json:"receiver_id,omitempty"`  // 接收消息用户id, 示例值: "ou_7346484524"
	ReceiveType *string `json:"receive_type,omitempty"` // 接收消息方式, chat(服务台专属服务群)或user(服务台机器人私聊)。若选择专属服务群, 用户有正在处理的工单将会发送失败。默认以chat方式发送, 示例值: "chat", 可选值有: chat: 通过服务台专属群发送, user: 通过服务台机器人私聊发送
}

// SendHelpdeskMessageResp ...
type SendHelpdeskMessageResp struct {
	MessageID string `json:"message_id,omitempty"` // chat消息open_id
}

// sendHelpdeskMessageResp ...
type sendHelpdeskMessageResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *SendHelpdeskMessageResp `json:"data,omitempty"`
}
