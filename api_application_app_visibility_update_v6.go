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

// UpdateApplicationAppVisibilityV6 该接口用于更新当前企业内自建应用或已安装的商店应用的可见范围, 包括可用人员与禁用人员。更新后对线上立即生效。
//
// 当通过接口新增用户或部门时, 提前判断对应用户或部门是否已在禁用名单中, 如果已在禁用名单中, 则即便将用户或部门添加到可用名单, 该用户或部门也无法看到该应用, 即禁用名单优先级高于可用名单。
// 同一个成员(user_id)在30s内不能重复添加到禁用名单, 否则会导致调用失败
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-visibility/patch
func (r *ApplicationService) UpdateApplicationAppVisibilityV6(ctx context.Context, request *UpdateApplicationAppVisibilityV6Req, options ...MethodOptionFunc) (*UpdateApplicationAppVisibilityV6Resp, *Response, error) {
	if r.cli.mock.mockApplicationUpdateApplicationAppVisibilityV6 != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Application#UpdateApplicationAppVisibilityV6 mock enable")
		return r.cli.mock.mockApplicationUpdateApplicationAppVisibilityV6(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Application",
		API:                   "UpdateApplicationAppVisibilityV6",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/application/v6/applications/:app_id/visibility",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateApplicationAppVisibilityV6Resp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApplicationUpdateApplicationAppVisibilityV6 mock ApplicationUpdateApplicationAppVisibilityV6 method
func (r *Mock) MockApplicationUpdateApplicationAppVisibilityV6(f func(ctx context.Context, request *UpdateApplicationAppVisibilityV6Req, options ...MethodOptionFunc) (*UpdateApplicationAppVisibilityV6Resp, *Response, error)) {
	r.mockApplicationUpdateApplicationAppVisibilityV6 = f
}

// UnMockApplicationUpdateApplicationAppVisibilityV6 un-mock ApplicationUpdateApplicationAppVisibilityV6 method
func (r *Mock) UnMockApplicationUpdateApplicationAppVisibilityV6() {
	r.mockApplicationUpdateApplicationAppVisibilityV6 = nil
}

// UpdateApplicationAppVisibilityV6Req ...
type UpdateApplicationAppVisibilityV6Req struct {
	AppID            string                                               `path:"app_id" json:"-"`              // 应用id, 示例值: "cli_9b445f5258795107"
	DepartmentIDType *DepartmentIDType                                    `query:"department_id_type" json:"-"` // 部门id 类型, 示例值: open_department_id, 可选值有: open_department_id: 以open_department_id标识部门, department_id: 以department_id标识部门, 默认值: `open_department_id`
	UserIDType       *IDType                                              `query:"user_id_type" json:"-"`       // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	AddVisibleList   *UpdateApplicationAppVisibilityV6ReqAddVisibleList   `json:"add_visible_list,omitempty"`   // 添加可用人员列表, 如果参数is_visible_to_all不设置且当前已经是全员可见, 或者参数is_visible_to_all设置为true, 则该参数不生效
	DelVisibleList   *UpdateApplicationAppVisibilityV6ReqDelVisibleList   `json:"del_visible_list,omitempty"`   // 删除可用人员列表, 如果参数is_visible_to_all不设置且当前已经是全员可见, 或者参数is_visible_to_all设置为true, 则该参数不生效
	AddInvisibleList *UpdateApplicationAppVisibilityV6ReqAddInvisibleList `json:"add_invisible_list,omitempty"` // 添加禁用人员列表
	DelInvisibleList *UpdateApplicationAppVisibilityV6ReqDelInvisibleList `json:"del_invisible_list,omitempty"` // 删除禁用人员列表
	IsVisibleToAll   *bool                                                `json:"is_visible_to_all,omitempty"`  // 是否全员可见, false: 否, true: 是, 不设置: 继续保持当前状态不改变, 如果参数不设置且当前已经是全员可见, 或者设置为true, 则add_visible_list/del_visible_list不生效, 示例值: false
}

// UpdateApplicationAppVisibilityV6ReqAddInvisibleList ...
type UpdateApplicationAppVisibilityV6ReqAddInvisibleList struct {
	UserIDs       []string `json:"user_ids,omitempty"`       // 成员id列表 id类型根据user_id_type参数指定, 相同的成员不能在30s内重复添加到禁用列表, 否则会导致调用失败, 示例值: ["ou_84aad35d084aa403a838cf73ee18467"], 最大长度: `100`
	DepartmentIDs []string `json:"department_ids,omitempty"` // 部门id列表 id类型根据department_id_type参数指定, 示例值: ["od-4e6ac4d14bcd5071a37a39de902c7141"], 最大长度: `100`
	GroupIDs      []string `json:"group_ids,omitempty"`      // 用户组id列表, 示例值: ["g193821"], 最大长度: `100`
}

// UpdateApplicationAppVisibilityV6ReqAddVisibleList ...
type UpdateApplicationAppVisibilityV6ReqAddVisibleList struct {
	UserIDs       []string `json:"user_ids,omitempty"`       // 成员id列表 id类型根据user_id_type参数指定, 示例值: ["ou_84aad35d084aa403a838cf73ee18467"], 最大长度: `100`
	DepartmentIDs []string `json:"department_ids,omitempty"` // 部门id列表 id类型根据department_id_type参数指定, 示例值: ["od-4e6ac4d14bcd5071a37a39de902c7141"], 最大长度: `100`
	GroupIDs      []string `json:"group_ids,omitempty"`      // 用户组id列表, 示例值: ["g193821"], 最大长度: `100`
}

// UpdateApplicationAppVisibilityV6ReqDelInvisibleList ...
type UpdateApplicationAppVisibilityV6ReqDelInvisibleList struct {
	UserIDs       []string `json:"user_ids,omitempty"`       // 成员id列表 id类型根据user_id_type参数指定, 示例值: ["ou_84aad35d084aa403a838cf73ee18467"], 最大长度: `100`
	DepartmentIDs []string `json:"department_ids,omitempty"` // 部门id列表 id类型根据department_id_type参数指定, 示例值: ["od-4e6ac4d14bcd5071a37a39de902c7141"], 最大长度: `100`
	GroupIDs      []string `json:"group_ids,omitempty"`      // 用户组id列表, 示例值: ["g193821"], 最大长度: `100`
}

// UpdateApplicationAppVisibilityV6ReqDelVisibleList ...
type UpdateApplicationAppVisibilityV6ReqDelVisibleList struct {
	UserIDs       []string `json:"user_ids,omitempty"`       // 成员id列表 id类型根据user_id_type参数指定, 示例值: ["ou_84aad35d084aa403a838cf73ee18467"], 最大长度: `100`
	DepartmentIDs []string `json:"department_ids,omitempty"` // 部门id列表 id类型根据department_id_type参数指定, 示例值: ["od-4e6ac4d14bcd5071a37a39de902c7141"], 最大长度: `100`
	GroupIDs      []string `json:"group_ids,omitempty"`      // 用户组id, 示例值: ["g193821"], 最大长度: `100`
}

// UpdateApplicationAppVisibilityV6Resp ...
type UpdateApplicationAppVisibilityV6Resp struct {
}

// updateApplicationAppVisibilityV6Resp ...
type updateApplicationAppVisibilityV6Resp struct {
	Code  int64                                 `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                                `json:"msg,omitempty"`  // 错误描述
	Data  *UpdateApplicationAppVisibilityV6Resp `json:"data,omitempty"`
	Error *ErrorDetail                          `json:"error,omitempty"`
}
