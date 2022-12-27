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

// GetMailGroupList 分页批量获取邮件组
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup/list
func (r *MailService) GetMailGroupList(ctx context.Context, request *GetMailGroupListReq, options ...MethodOptionFunc) (*GetMailGroupListResp, *Response, error) {
	if r.cli.mock.mockMailGetMailGroupList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Mail#GetMailGroupList mock enable")
		return r.cli.mock.mockMailGetMailGroupList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Mail",
		API:                   "GetMailGroupList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/mail/v1/mailgroups",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getMailGroupListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMailGetMailGroupList mock MailGetMailGroupList method
func (r *Mock) MockMailGetMailGroupList(f func(ctx context.Context, request *GetMailGroupListReq, options ...MethodOptionFunc) (*GetMailGroupListResp, *Response, error)) {
	r.mockMailGetMailGroupList = f
}

// UnMockMailGetMailGroupList un-mock MailGetMailGroupList method
func (r *Mock) UnMockMailGetMailGroupList() {
	r.mockMailGetMailGroupList = nil
}

// GetMailGroupListReq ...
type GetMailGroupListReq struct {
	ManagerUserID *string `query:"manager_user_id" json:"-"` // 邮件组管理员用户ID, 用于获取该用户有管理权限的邮件组, 示例值: "ou_xxxxxx"
	UserIDType    *IDType `query:"user_id_type" json:"-"`    // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	PageToken     *string `query:"page_token" json:"-"`      // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "xxx"
	PageSize      *int64  `query:"page_size" json:"-"`       // 分页大小, 示例值: 10, 默认值: `20`, 最大值: `200`
}

// GetMailGroupListResp ...
type GetMailGroupListResp struct {
	HasMore   bool                        `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                      `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	Items     []*GetMailGroupListRespItem `json:"items,omitempty"`      // 邮件组列表
}

// GetMailGroupListRespItem ...
type GetMailGroupListRespItem struct {
	MailGroupID             string `json:"mailgroup_id,omitempty"`               // 邮件组ID
	Email                   string `json:"email,omitempty"`                      // 邮件组地址
	Name                    string `json:"name,omitempty"`                       // 邮件组名称
	Description             string `json:"description,omitempty"`                // 邮件组描述
	DirectMembersCount      string `json:"direct_members_count,omitempty"`       // 邮件组成员数量
	IncludeExternalMember   bool   `json:"include_external_member,omitempty"`    // 是否包含外部成员
	IncludeAllCompanyMember bool   `json:"include_all_company_member,omitempty"` // 是否是全员邮件组
	WhoCanSendMail          string `json:"who_can_send_mail,omitempty"`          // 谁可发送邮件到此邮件组, 可选值有: ANYONE: 任何人, ALL_INTERNAL_USERS: 仅组织内部成员, ALL_GROUP_MEMBERS: 仅邮件组成员, CUSTOM_MEMBERS: 自定义成员
}

// getMailGroupListResp ...
type getMailGroupListResp struct {
	Code int64                 `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                `json:"msg,omitempty"`  // 错误描述
	Data *GetMailGroupListResp `json:"data,omitempty"`
}
