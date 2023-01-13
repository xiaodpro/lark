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

// SendRawMessageOld 为了更好地提升该接口的安全性, 我们对其进行了升级, 请尽快迁移至[新版本>>](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/create)
//
// 给指定用户或者会话发送文本消息, 其中会话包括私聊会话和群会话。
// - 需要启用机器人能力；私聊会话时机器人需要拥有对用户的可见性, 群会话需要机器人在群里
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUjNz4SN2MjL1YzM
func (r *MessageService) SendRawMessageOld(ctx context.Context, request *SendRawMessageOldReq, options ...MethodOptionFunc) (*SendRawMessageOldResp, *Response, error) {
	if r.cli.mock.mockMessageSendRawMessageOld != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Message#SendRawMessageOld mock enable")
		return r.cli.mock.mockMessageSendRawMessageOld(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Message",
		API:                   "SendRawMessageOld",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/message/v4/send/",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(sendRawMessageOldResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMessageSendRawMessageOld mock MessageSendRawMessageOld method
func (r *Mock) MockMessageSendRawMessageOld(f func(ctx context.Context, request *SendRawMessageOldReq, options ...MethodOptionFunc) (*SendRawMessageOldResp, *Response, error)) {
	r.mockMessageSendRawMessageOld = f
}

// UnMockMessageSendRawMessageOld un-mock MessageSendRawMessageOld method
func (r *Mock) UnMockMessageSendRawMessageOld() {
	r.mockMessageSendRawMessageOld = nil
}

// SendRawMessageOldReq ...
type SendRawMessageOldReq struct {
	OpenID  string                       `json:"open_id,omitempty"`  // 给用户发私聊消息, 只需要填 open_id([什么是 Open ID？](https://open.feishu.cn/document/home/user-identity-introduction/open-id))、email（真实邮箱）、user_id([什么是 User ID？](https://open.feishu.cn/document/home/user-identity-introduction/user-id)) 中的一个即可, 向群里发消息使用群的 chat_id。服务端依次读取字段的顺序为 chat_id > open_id > user_id > email   ( user_id 对应V3接口的 employee_id, chat_id 对应V3的 open_chat_id )
	UserID  string                       `json:"user_id,omitempty"`  // 给用户发私聊消息, 只需要填 open_id([什么是 Open ID？](https://open.feishu.cn/document/home/user-identity-introduction/open-id))、email（真实邮箱）、user_id([什么是 User ID？](https://open.feishu.cn/document/home/user-identity-introduction/user-id)) 中的一个即可, 向群里发消息使用群的 chat_id。服务端依次读取字段的顺序为 chat_id > open_id > user_id > email   ( user_id 对应V3接口的 employee_id, chat_id 对应V3的 open_chat_id )
	Email   string                       `json:"email,omitempty"`    // 给用户发私聊消息, 只需要填 open_id([什么是 Open ID？](https://open.feishu.cn/document/home/user-identity-introduction/open-id))、email（真实邮箱）、user_id([什么是 User ID？](https://open.feishu.cn/document/home/user-identity-introduction/user-id)) 中的一个即可, 向群里发消息使用群的 chat_id。服务端依次读取字段的顺序为 chat_id > open_id > user_id > email   ( user_id 对应V3接口的 employee_id, chat_id 对应V3的 open_chat_id )
	ChatID  string                       `json:"chat_id,omitempty"`  // 给用户发私聊消息, 只需要填 open_id([什么是 Open ID？](https://open.feishu.cn/document/home/user-identity-introduction/open-id))、email（真实邮箱）、user_id([什么是 User ID？](https://open.feishu.cn/document/home/user-identity-introduction/user-id)) 中的一个即可, 向群里发消息使用群的 chat_id。服务端依次读取字段的顺序为 chat_id > open_id > user_id > email   ( user_id 对应V3接口的 employee_id, chat_id 对应V3的 open_chat_id )
	RootID  *string                      `json:"root_id,omitempty"`  // 如果需要回复某条消息, 填对应消息的消息 ID
	MsgType MsgType                      `json:"msg_type,omitempty"` // 消息类型, 此处固定填 "text"
	Content *SendRawMessageOldReqContent `json:"content,omitempty"`  // 消息内容
}

// SendRawMessageOldReqContent ...
type SendRawMessageOldReqContent struct {
	Text     string                 `json:"text,omitempty"`      // 文本消息内容, 文本消息中可以 at 个人或全体成员 at 全体成员: <at user_id="all">  </at>   at 个人: <at user_id="ou_xxxxxxx"></at>, user_id 为用户 user_id或者open_id
	ImageKey string                 `json:"image_key,omitempty"` // image_key 可以通过图片上传接口获得
	Post     *MessageContentPostAll `json:"post,omitempty"`      // 富文本消息
}

// SendRawMessageOldResp ...
type SendRawMessageOldResp struct {
	MessageID string `json:"message_id,omitempty"` // 消息 ID
}

// sendRawMessageOldResp ...
type sendRawMessageOldResp struct {
	Code int64                  `json:"code,omitempty"` // 返回码, 非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 返回码描述
	Data *SendRawMessageOldResp `json:"data,omitempty"`
}
