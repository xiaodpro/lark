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

// UpdateChat 更新群头像、群名称、群描述、群配置、转让群主等。
//
// 注意事项:
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/home/develop-a-bot-in-5-minutes/create-an-app)
// - 若群未开启 [仅群主和群管理员可编辑群信息] 配置:
// - 群主/群管理员 或 创建群组且具备[更新应用所创建群的群信息]权限的机器人, 可更新所有信息
// - 不满足上述条件的群成员或机器人, 仅可更新群头像、群名称、群描述、群国际化名称信息
// - 若群开启了[仅群主和群管理员可编辑群信息]配置:
// - 群主/群管理员 或 创建群组且具备[更新应用所创建群的群信息]权限的机器人, 可更新所有信息
// - 不满足上述条件的群成员或者机器人, 任何群信息都不能修改
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/update
func (r *ChatService) UpdateChat(ctx context.Context, request *UpdateChatReq, options ...MethodOptionFunc) (*UpdateChatResp, *Response, error) {
	if r.cli.mock.mockChatUpdateChat != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Chat#UpdateChat mock enable")
		return r.cli.mock.mockChatUpdateChat(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Chat",
		API:                   "UpdateChat",
		Method:                "PUT",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/chats/:chat_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateChatResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockChatUpdateChat mock ChatUpdateChat method
func (r *Mock) MockChatUpdateChat(f func(ctx context.Context, request *UpdateChatReq, options ...MethodOptionFunc) (*UpdateChatResp, *Response, error)) {
	r.mockChatUpdateChat = f
}

// UnMockChatUpdateChat un-mock ChatUpdateChat method
func (r *Mock) UnMockChatUpdateChat() {
	r.mockChatUpdateChat = nil
}

// UpdateChatReq ...
type UpdateChatReq struct {
	ChatID                 string               `path:"chat_id" json:"-"`                   // 群 ID, 详情参见[群ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description), 示例值: "oc_a0553eda9014c201e6969b478895c230"
	UserIDType             *IDType              `query:"user_id_type" json:"-"`             // 用户 ID 类型, 示例值: "open_id", 可选值有: `open_id`: 用户的 open id, `union_id`: 用户的 union id, `user_id`: 用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	Avatar                 *string              `json:"avatar,omitempty"`                   // 群头像对应的 Image Key, 可通过[上传图片](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/image/create)获取（注意: 上传图片的 [image_type] 需要指定为 [avatar]）, 示例值: "default-avatar_44ae0ca3-e140-494b-956f-78091e348435"
	Name                   *string              `json:"name,omitempty"`                     // 群名称, 示例值: "群聊"
	Description            *string              `json:"description,omitempty"`              // 群描述, 示例值: "测试群描述"
	I18nNames              *I18nNames           `json:"i18n_names,omitempty"`               // 群国际化名称
	AddMemberPermission    *AddMemberPermission `json:"add_member_permission,omitempty"`    // 加 user/bot 入群权限(all_members/only_owner), 示例值: "all_members"
	ShareCardPermission    *ShareCardPermission `json:"share_card_permission,omitempty"`    // 群分享权限(allowed/not_allowed), 示例值: "allowed"
	AtAllPermission        *AtAllPermission     `json:"at_all_permission,omitempty"`        // at 所有人权限(all_members/only_owner), 示例值: "all_members"
	EditPermission         *EditPermission      `json:"edit_permission,omitempty"`          // 群编辑权限(all_members/only_owner), 示例值: "all_members"
	OwnerID                *string              `json:"owner_id,omitempty"`                 // 新群主 ID, 示例值: "4d7a3c6g"
	JoinMessageVisibility  *MessageVisibility   `json:"join_message_visibility,omitempty"`  // 入群消息可见性(only_owner/all_members/not_anyone), 示例值: "only_owner"
	LeaveMessageVisibility *MessageVisibility   `json:"leave_message_visibility,omitempty"` // 出群消息可见性(only_owner/all_members/not_anyone), 示例值: "only_owner"
	MembershipApproval     *MembershipApproval  `json:"membership_approval,omitempty"`      // 加群审批(no_approval_required/approval_required), 示例值: "no_approval_required"
}

// UpdateChatResp ...
type UpdateChatResp struct {
}

// updateChatResp ...
type updateChatResp struct {
	Code int64           `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 错误描述
	Data *UpdateChatResp `json:"data,omitempty"`
}
