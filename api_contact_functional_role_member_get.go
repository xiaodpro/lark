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

// GetContactFunctionalRoleMember 通过本接口可以查询角色ID下的成员信息（含成员ID及其管理范围）
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/functional_role-member/list
// new doc: https://open.feishu.cn/document/server-docs/contact-v3/functional_role-member/list
func (r *ContactService) GetContactFunctionalRoleMember(ctx context.Context, request *GetContactFunctionalRoleMemberReq, options ...MethodOptionFunc) (*GetContactFunctionalRoleMemberResp, *Response, error) {
	if r.cli.mock.mockContactGetContactFunctionalRoleMember != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Contact#GetContactFunctionalRoleMember mock enable")
		return r.cli.mock.mockContactGetContactFunctionalRoleMember(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "GetContactFunctionalRoleMember",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/functional_roles/:role_id/members",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getContactFunctionalRoleMemberResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactGetContactFunctionalRoleMember mock ContactGetContactFunctionalRoleMember method
func (r *Mock) MockContactGetContactFunctionalRoleMember(f func(ctx context.Context, request *GetContactFunctionalRoleMemberReq, options ...MethodOptionFunc) (*GetContactFunctionalRoleMemberResp, *Response, error)) {
	r.mockContactGetContactFunctionalRoleMember = f
}

// UnMockContactGetContactFunctionalRoleMember un-mock ContactGetContactFunctionalRoleMember method
func (r *Mock) UnMockContactGetContactFunctionalRoleMember() {
	r.mockContactGetContactFunctionalRoleMember = nil
}

// GetContactFunctionalRoleMemberReq ...
type GetContactFunctionalRoleMemberReq struct {
	RoleID           string            `path:"role_id" json:"-"`             // 角色的唯一标识, 单租户下唯一, 示例值: "7vrj3vk70xk7v5r"
	PageSize         *int64            `query:"page_size" json:"-"`          // 分页大小, 示例值: 50, 默认值: `20`, 取值范围: `0` ～ `100`
	PageToken        *string           `query:"page_token" json:"-"`         // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: dawdewd
	UserIDType       *IDType           `query:"user_id_type" json:"-"`       // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	DepartmentIDType *DepartmentIDType `query:"department_id_type" json:"-"` // 此次调用中使用的部门ID的类型, 示例值: open_department_id, 可选值有: department_id: 以自定义department_id来标识部门, open_department_id: 以open_department_id来标识部门, 默认值: `open_department_id`
}

// GetContactFunctionalRoleMemberResp ...
type GetContactFunctionalRoleMemberResp struct {
	Members   []*GetContactFunctionalRoleMemberRespMember `json:"members,omitempty"`    // 角色成员列表
	PageToken string                                      `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore   bool                                        `json:"has_more,omitempty"`   // 是否还有更多项
}

// GetContactFunctionalRoleMemberRespMember ...
type GetContactFunctionalRoleMemberRespMember struct {
	UserID        string   `json:"user_id,omitempty"`        // 成员ID
	ScopeType     string   `json:"scope_type,omitempty"`     // 管理范围的类型, 可选值有: All: 管理范围是全部, Part: 管理范围是部分, None: 管理范围为空
	DepartmentIDs []string `json:"department_ids,omitempty"` // 表示该角色成员的管理范围, scope_type为“指定范围”时, 返回该值
}

// getContactFunctionalRoleMemberResp ...
type getContactFunctionalRoleMemberResp struct {
	Code  int64                               `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                              `json:"msg,omitempty"`  // 错误描述
	Data  *GetContactFunctionalRoleMemberResp `json:"data,omitempty"`
	Error *ErrorDetail                        `json:"error,omitempty"`
}
