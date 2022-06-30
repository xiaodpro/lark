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

// DeleteChatManager 删除指定的群管理员（用户或机器人）
//
// 注意事项:
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/home/develop-a-bot-in-5-minutes/create-an-app)
// - 仅有群主可以删除群管理员
// - 每次请求最多指定 50 个用户或者 5 个机器人
// - 删除机器人类型的管理员请使用 [app_id]
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-managers/delete_managers
func (r *ChatService) DeleteChatManager(ctx context.Context, request *DeleteChatManagerReq, options ...MethodOptionFunc) (*DeleteChatManagerResp, *Response, error) {
	if r.cli.mock.mockChatDeleteChatManager != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Chat#DeleteChatManager mock enable")
		return r.cli.mock.mockChatDeleteChatManager(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Chat",
		API:                   "DeleteChatManager",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/chats/:chat_id/managers/delete_managers",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(deleteChatManagerResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockChatDeleteChatManager mock ChatDeleteChatManager method
func (r *Mock) MockChatDeleteChatManager(f func(ctx context.Context, request *DeleteChatManagerReq, options ...MethodOptionFunc) (*DeleteChatManagerResp, *Response, error)) {
	r.mockChatDeleteChatManager = f
}

// UnMockChatDeleteChatManager un-mock ChatDeleteChatManager method
func (r *Mock) UnMockChatDeleteChatManager() {
	r.mockChatDeleteChatManager = nil
}

// DeleteChatManagerReq ...
type DeleteChatManagerReq struct {
	ChatID       string   `path:"chat_id" json:"-"`         // 群 ID, 详情参见[群ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description), 示例值: "oc_a0553eda9014c201e6969b478895c230"
	MemberIDType *IDType  `query:"member_id_type" json:"-"` // 群成员 id 类型 open_id/user_id/union_id/app_id, 示例值: "open_id", 可选值有: `user_id`: 以 user_id 来识别成员, 需要有获取用户 UserID 的权限 ([什么是 User ID？](https://open.feishu.cn/document/home/user-identity-introduction/user-id)), `union_id`: 以 union_id 来识别成员([什么是 Union ID？](https://open.feishu.cn/document/home/user-identity-introduction/union-id)), `open_id`: 以 open_id 来识别成员([什么是 Open ID？](https://open.feishu.cn/document/home/user-identity-introduction/open-id)), `app_id`: 以 app_id 来识别成员([获取应用身份访问凭证](https://open.feishu.cn/document/ukTMukTMukTM/ukDNz4SO0MjL5QzM/g))
	ManagerIDs   []string `json:"manager_ids,omitempty"`    // 要删除的 manager_id, 示例值: ["ou_9204a37300b3700d61effaa439f34295"], 最大长度: `50`
}

// DeleteChatManagerResp ...
type DeleteChatManagerResp struct {
	ChatManagers    []string `json:"chat_managers,omitempty"`     // 群目前用户类型的管理员 id
	ChatBotManagers []string `json:"chat_bot_managers,omitempty"` // 群目前机器人类型的管理员 id
}

// deleteChatManagerResp ...
type deleteChatManagerResp struct {
	Code int64                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *DeleteChatManagerResp `json:"data,omitempty"`
}
