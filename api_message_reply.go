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

// ReplyRawMessage 回复指定消息, 支持文本、富文本、卡片、群名片、个人名片、图片、视频、文件等多种消息类型。
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/home/develop-a-bot-in-5-minutes/create-an-app)
// - 回复私聊消息, 需要机器人对用户有可用性
// - 回复群组消息, 需要机器人在群中
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/reply
func (r *MessageService) ReplyRawMessage(ctx context.Context, request *ReplyRawMessageReq, options ...MethodOptionFunc) (*ReplyRawMessageResp, *Response, error) {
	if r.cli.mock.mockMessageReplyRawMessage != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Message#ReplyRawMessage mock enable")
		return r.cli.mock.mockMessageReplyRawMessage(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Message",
		API:                   "ReplyRawMessage",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/messages/:message_id/reply",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(replyRawMessageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMessageReplyRawMessage mock MessageReplyRawMessage method
func (r *Mock) MockMessageReplyRawMessage(f func(ctx context.Context, request *ReplyRawMessageReq, options ...MethodOptionFunc) (*ReplyRawMessageResp, *Response, error)) {
	r.mockMessageReplyRawMessage = f
}

// UnMockMessageReplyRawMessage un-mock MessageReplyRawMessage method
func (r *Mock) UnMockMessageReplyRawMessage() {
	r.mockMessageReplyRawMessage = nil
}

// ReplyRawMessageReq ...
type ReplyRawMessageReq struct {
	MessageID string  `path:"message_id" json:"-"` // 待回复的消息的ID, 示例值: "om_dc13264520392913993dd051dba21dcf"
	Content   string  `json:"content,omitempty"`   // 消息内容 json 格式, 格式说明参考: [发送消息content说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json), 示例值: "{\"text\":\"<at user_id=\\\"ou_155184d1e73cbfb8973e5a9e698e74f2\\\">Tom</at> test content\"}"
	MsgType   MsgType `json:"msg_type,omitempty"`  // 消息类型, 包括: text、post、image、file、audio、media、sticker、interactive、share_card、share_user, 示例值: "text"
}

// ReplyRawMessageResp ...
type ReplyRawMessageResp struct {
	MessageID      string       `json:"message_id,omitempty"`       // 消息id, 说明参见: [消息ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/intro#ac79c1c2)
	RootID         string       `json:"root_id,omitempty"`          // 根消息id, 说明参见: [消息ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/intro#ac79c1c2)
	ParentID       string       `json:"parent_id,omitempty"`        // 父消息的id, 说明参见: [消息ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/intro#ac79c1c2)
	MsgType        MsgType      `json:"msg_type,omitempty"`         // 消息类型 包括: text、post、image、file、audio、media、sticker、interactive、share_chat、share_user等, 类型定义请参考[发送消息content说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json)
	CreateTime     string       `json:"create_time,omitempty"`      // 消息生成的时间戳（毫秒）
	UpdateTime     string       `json:"update_time,omitempty"`      // 消息更新的时间戳（毫秒）
	Deleted        bool         `json:"deleted,omitempty"`          // 消息是否被撤回
	Updated        bool         `json:"updated,omitempty"`          // 消息是否被更新
	ChatID         string       `json:"chat_id,omitempty"`          // 所属的群
	Sender         *Sender      `json:"sender,omitempty"`           // 发送者, 可以是用户或应用
	Body           *MessageBody `json:"body,omitempty"`             // 消息内容
	Mentions       []*Mention   `json:"mentions,omitempty"`         // 被@的用户或机器人的id列表
	UpperMessageID string       `json:"upper_message_id,omitempty"` // 合并转发消息中, 上一层级的消息id message_id, 说明参见: [消息ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/intro#ac79c1c2)
}

// replyRawMessageResp ...
type replyRawMessageResp struct {
	Code int64                `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string               `json:"msg,omitempty"`  // 错误描述
	Data *ReplyRawMessageResp `json:"data,omitempty"`
}
