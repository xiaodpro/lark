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

// AddChatMember 将用户或机器人拉入群聊。
//
// 注意事项:
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
// - 如需拉用户进群, 需要机器人对用户有[可用性](https://open.feishu.cn/document/home/introduction-to-scope-and-authorization/availability)
// - 如需拉机器人进群, 需要被拉的机器人开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
// - 机器人或授权用户必须在群组中
// - 外部租户不能被加入到内部群中
// - 操作内部群时, 操作者须与群组在同一租户下
// - 在开启 [仅群主和群管理员可添加群成员] 的设置时, 仅有群主/管理员 或 创建群组且具备 [更新应用所创建群的群信息] 权限的机器人, 可以拉用户或者机器人进群
// - 在未开启 [仅群主和群管理员可添加群成员] 的设置时, 所有群成员都可以拉用户或机器人进群
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/create
func (r *ChatService) AddChatMember(ctx context.Context, request *AddChatMemberReq, options ...MethodOptionFunc) (*AddChatMemberResp, *Response, error) {
	if r.cli.mock.mockChatAddChatMember != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Chat#AddChatMember mock enable")
		return r.cli.mock.mockChatAddChatMember(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Chat",
		API:                   "AddChatMember",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/chats/:chat_id/members",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(addChatMemberResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockChatAddChatMember mock ChatAddChatMember method
func (r *Mock) MockChatAddChatMember(f func(ctx context.Context, request *AddChatMemberReq, options ...MethodOptionFunc) (*AddChatMemberResp, *Response, error)) {
	r.mockChatAddChatMember = f
}

// UnMockChatAddChatMember un-mock ChatAddChatMember method
func (r *Mock) UnMockChatAddChatMember() {
	r.mockChatAddChatMember = nil
}

// AddChatMemberReq ...
type AddChatMemberReq struct {
	ChatID       string   `path:"chat_id" json:"-"`         // 群 ID, 详情参见[群ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description), 注意: 仅支持群模式为`group`、`topic`的群组ID, 示例值: "oc_a0553eda9014c201e6969b478895c230"
	MemberIDType *IDType  `query:"member_id_type" json:"-"` // 进群成员 ID 类型 open_id/user_id/union_id/app_id, 注意: 拉机器人入群请使用 [app_id], 示例值: "open_id", 可选值有: user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), app_id: 飞书开放平台应用的唯一标识。在创建应用时, 由系统自动生成, 用户不能自行修改。[了解更多: 如何获取应用的 App ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-app-id), 默认值: `open_id`
	SucceedType  *int64   `query:"succeed_type" json:"-"`   // 出现不可用ID后的处理方式 0/1/2, 示例值: 0, 可选值有: 0: 兼容之前的策略, 不存在/不可见的 ID 会拉群失败, 并返回错误响应。存在已离职 ID 时, 会将其他可用 ID 拉入群聊, 返回拉群成功的响应。, 1: 将参数中可用的 ID 全部拉入群聊, 返回拉群成功的响应, 并展示剩余不可用的 ID 及原因。, 2: 参数中只要存在任一不可用的 ID, 就会拉群失败, 返回错误响应, 并展示出不可用的 ID。
	IDList       []string `json:"id_list,omitempty"`        // 成员ID列表。邀请用户进群时推荐使用 OpenID, 获取方式可参考文档[如何获取 Open ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid)；邀请机器人进群时需填写应用的App ID, 请参考[如何获取应用的 App ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-app-id), 注意: 成员列表不可为空, 每次请求最多拉50个用户或者5个机器人, 并且群组最多容纳15个机器人, 列表中填写的成员ID类型应与 [member_id_type] 参数中选择的类型相对应, 对于已认证企业的飞书的群人数默认上限: 普通群5000人, 会议群3000人, 话题群5000人。若租户管理员配置了群人数上限, 则群人数上限为该人数上限, 示例值: ["ou_9204a37300b3700d61effaa439f34295"]
}

// AddChatMemberResp ...
type AddChatMemberResp struct {
	InvalidIDList    []string `json:"invalid_id_list,omitempty"`     // 无效成员列表, 注意: 当`success_type=0`时, `invalid_id_list`只包含已离职的用户ID, 当`success_type=1`时, `invalid_id_list`中包含已离职的、不可见的、应用未激活的ID
	NotExistedIDList []string `json:"not_existed_id_list,omitempty"` // ID不存在的成员列表
}

// addChatMemberResp ...
type addChatMemberResp struct {
	Code int64              `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string             `json:"msg,omitempty"`  // 错误描述
	Data *AddChatMemberResp `json:"data,omitempty"`
}
