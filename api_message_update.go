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

// UpdateMessage 更新应用已发送的消息卡片内容。
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/home/develop-a-bot-in-5-minutes/create-an-app)
// - 当前仅支持更新 卡片消息
// - 不支持更新批量消息
// - 只支持对所有人都更新的[「共享卡片」](ukTMukTMukTM/uAjNwUjLwYDM14CM2ATN), 也即需要在卡片的`config`属性中, 显式声明`"update_multi":true`。如果你只想更新特定人的消息卡片, 必须要用户在卡片操作交互后触发, 开发文档参考[「独享卡片」](https://open.feishu.cn/document/ukTMukTMukTM/uYjNwUjL2YDM14iN2ATN#49904b71)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/patch
func (r *MessageService) UpdateMessage(ctx context.Context, request *UpdateMessageReq, options ...MethodOptionFunc) (*UpdateMessageResp, *Response, error) {
	if r.cli.mock.mockMessageUpdateMessage != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Message#UpdateMessage mock enable")
		return r.cli.mock.mockMessageUpdateMessage(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Message",
		API:                   "UpdateMessage",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/messages/:message_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateMessageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMessageUpdateMessage mock MessageUpdateMessage method
func (r *Mock) MockMessageUpdateMessage(f func(ctx context.Context, request *UpdateMessageReq, options ...MethodOptionFunc) (*UpdateMessageResp, *Response, error)) {
	r.mockMessageUpdateMessage = f
}

// UnMockMessageUpdateMessage un-mock MessageUpdateMessage method
func (r *Mock) UnMockMessageUpdateMessage() {
	r.mockMessageUpdateMessage = nil
}

// UpdateMessageReq ...
type UpdateMessageReq struct {
	MessageID string `path:"message_id" json:"-"` // 待更新的消息的ID, 示例值: "om_dc13264520392913993dd051dba21dcf"
	Content   string `json:"content,omitempty"`   // 消息内容 json 格式, [发送消息 content 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json), 参考文档中的卡片格式, 示例值: "参考链接"
}

// UpdateMessageResp ...
type UpdateMessageResp struct {
}

// updateMessageResp ...
type updateMessageResp struct {
	Code int64              `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string             `json:"msg,omitempty"`  // 错误描述
	Data *UpdateMessageResp `json:"data,omitempty"`
}
