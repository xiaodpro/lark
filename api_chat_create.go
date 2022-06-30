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

// CreateChat 创建群并设置群头像、群名、群描述等。
//
// 注意事项:
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/home/develop-a-bot-in-5-minutes/create-an-app)
// - 本接口只支持创建群, 如果需要拉用户或者机器人入群参考 [将用户或机器人拉入群聊](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/create)接口
// - 每次请求, 最多拉 50 个用户或者 5 个机器人, 并且群组最多容纳 15 个机器人
// - 拉机器人入群请使用 [app_id]
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/create
func (r *ChatService) CreateChat(ctx context.Context, request *CreateChatReq, options ...MethodOptionFunc) (*CreateChatResp, *Response, error) {
	if r.cli.mock.mockChatCreateChat != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Chat#CreateChat mock enable")
		return r.cli.mock.mockChatCreateChat(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Chat",
		API:                   "CreateChat",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/chats",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createChatResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockChatCreateChat mock ChatCreateChat method
func (r *Mock) MockChatCreateChat(f func(ctx context.Context, request *CreateChatReq, options ...MethodOptionFunc) (*CreateChatResp, *Response, error)) {
	r.mockChatCreateChat = f
}

// UnMockChatCreateChat un-mock ChatCreateChat method
func (r *Mock) UnMockChatCreateChat() {
	r.mockChatCreateChat = nil
}

// CreateChatReq ...
type CreateChatReq struct {
	UserIDType             *IDType             `query:"user_id_type" json:"-"`             // 用户 ID 类型, 示例值: "open_id", 可选值有: `open_id`: 用户的 open id, `union_id`: 用户的 union id, `user_id`: 用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	SetBotManager          *bool               `query:"set_bot_manager" json:"-"`          // 如果选择了设置群主为指定用户, 可以选择是否同时设置创建此群的机器人为管理员, 此标志位用于标记是否设置创建群的机器人为管理员, 示例值: false
	Avatar                 *string             `json:"avatar,omitempty"`                   // 群头像对应的 Image Key, 可通过[上传图片](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/image/create)获取（注意: 上传图片的 [image_type] 需要指定为 [avatar]）, 示例值: "default-avatar_44ae0ca3-e140-494b-956f-78091e348435"
	Name                   *string             `json:"name,omitempty"`                     // 群名称, 示例值: "测试群名称"
	Description            *string             `json:"description,omitempty"`              // 群描述, 示例值: "测试群描述"
	I18nNames              *I18nNames          `json:"i18n_names,omitempty"`               // 群国际化名称
	OwnerID                *string             `json:"owner_id,omitempty"`                 // 创建群时指定的群主, 不填时指定建群的机器人为群主, 群主 ID, ID值与查询参数中的 user_id_type 对应, 不同 ID 的说明参见 [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction), 示例值: "4d7a3c6g"
	UserIDList             []string            `json:"user_id_list,omitempty"`             // 创建群时邀请的群成员, id 类型为 user_id_type, 示例值: ["4d7a3c6g"], 最大长度: `50`
	BotIDList              []string            `json:"bot_id_list,omitempty"`              // 创建群时邀请的群机器人, 示例值: ["cli_a10fbf7e94b8d01d"], 最大长度: `5`
	ChatMode               *ChatMode           `json:"chat_mode,omitempty"`                // 群模式, 可选值有: `group`: 群组, 示例值: "group"
	ChatType               *ChatType           `json:"chat_type,omitempty"`                // 群类型, 可选值有: `private`: 私有群, `public`: 公开群, 示例值: "private"
	External               *bool               `json:"external,omitempty"`                 // 是否是外部群, 示例值: false
	JoinMessageVisibility  *MessageVisibility  `json:"join_message_visibility,omitempty"`  // 入群消息可见性, 可选值有: `only_owner`: 仅群主和管理员可见, `all_members`: 所有成员可见, `not_anyone`: 任何人均不可见, 示例值: "all_members"
	LeaveMessageVisibility *MessageVisibility  `json:"leave_message_visibility,omitempty"` // 退群消息可见性, 可选值有: `only_owner`: 仅群主和管理员可见, `all_members`: 所有成员可见, `not_anyone`: 任何人均不可见, 示例值: "all_members"
	MembershipApproval     *MembershipApproval `json:"membership_approval,omitempty"`      // 加群审批, 可选值有: `no_approval_required`: 无需审批, `approval_required`: 需要审批, 示例值: "no_approval_required"
}

// CreateChatResp ...
type CreateChatResp struct {
	ChatID                 string               `json:"chat_id,omitempty"`                  // 群 ID, 详情参见: [群ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description)
	Avatar                 string               `json:"avatar,omitempty"`                   // 群头像 URL
	Name                   string               `json:"name,omitempty"`                     // 群名称
	Description            string               `json:"description,omitempty"`              // 群描述
	I18nNames              *I18nNames           `json:"i18n_names,omitempty"`               // 群国际化名称
	OwnerID                string               `json:"owner_id,omitempty"`                 // 群主 ID, ID值与查询参数中的 user_id_type 对应, 不同 ID 的说明参见 [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction), 当群主是机器人时, 该字段不返回
	OwnerIDType            IDType               `json:"owner_id_type,omitempty"`            // 群主 ID 对应的ID类型, 与查询参数中的 user_id_type 相同。取值为: `open_id`、`user_id`、`union_id`其中之一, 当群主是机器人时, 该字段不返回
	AddMemberPermission    AddMemberPermission  `json:"add_member_permission,omitempty"`    // 拉 用户或机器人 入群权限, 可选值有: `only_owner`: 仅群主和管理员, `all_members`: 所有成员
	ShareCardPermission    ShareCardPermission  `json:"share_card_permission,omitempty"`    // 群分享权限, 可选值有: `allowed`: 允许, `not_allowed`: 不允许
	AtAllPermission        AtAllPermission      `json:"at_all_permission,omitempty"`        // at 所有人权限, 可选值有: `only_owner`: 仅群主和管理员, `all_members`: 所有成员
	EditPermission         EditPermission       `json:"edit_permission,omitempty"`          // 群编辑权限, 可选值有: `only_owner`: 仅群主和管理员, `all_members`: 所有成员
	ChatMode               ChatMode             `json:"chat_mode,omitempty"`                // 群模式, 可选值有: `group`: 群组
	ChatType               ChatType             `json:"chat_type,omitempty"`                // 群类型, 可选值有: `private`: 私有群, `public`: 公开群
	ChatTag                string               `json:"chat_tag,omitempty"`                 // 群标签, 如有多个, 则按照下列顺序返回第一个, 可选值有: `inner`: 内部群, `tenant`: 公司群, `department`: 部门群, `edu`: 教育群, `meeting`: 会议群, `customer_service`: 客服群
	External               bool                 `json:"external,omitempty"`                 // 是否是外部群
	TenantKey              string               `json:"tenant_key,omitempty"`               // 租户在飞书上的唯一标识, 用来换取对应的tenant_access_token, 也可以用作租户在应用里面的唯一标识
	JoinMessageVisibility  MessageVisibility    `json:"join_message_visibility,omitempty"`  // 入群消息可见性, 可选值有: `only_owner`: 仅群主和管理员可见, `all_members`: 所有成员可见, `not_anyone`: 任何人均不可见
	LeaveMessageVisibility MessageVisibility    `json:"leave_message_visibility,omitempty"` // 出群消息可见性, 可选值有: `only_owner`: 仅群主和管理员可见, `all_members`: 所有成员可见, `not_anyone`: 任何人均不可见
	MembershipApproval     MembershipApproval   `json:"membership_approval,omitempty"`      // 加群审批, 可选值有: `no_approval_required`: 无需审批, `approval_required`: 需要审批
	ModerationPermission   ModerationPermission `json:"moderation_permission,omitempty"`    // 发言权限, 可选值有: `only_owner`: 仅群主和管理员, `all_members`: 所有成员, `moderator_list`: 指定群成员
}

// createChatResp ...
type createChatResp struct {
	Code int64           `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 错误描述
	Data *CreateChatResp `json:"data,omitempty"`
}
