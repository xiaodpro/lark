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

// GetApplicationVersionContactsRangeSuggest 该接口用于根据应用的 App ID 和版本 ID 获取企业自建应用某个版本的通讯录权限范围。
//
// 由于通讯录权限范围需要提交发布新的应用版本, 并且企业管理员审核通过后才会生效, 因此该权限范围可能与实际生效的权限范围有差别, 如需获取线上实际生效的通讯录权限范围, 可通过[获取应用通讯录权限范围配置](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application/contacts_range_configuration) 获取。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-app_version/contacts_range_suggest
// new doc: https://open.feishu.cn/document/server-docs/application-v6/application/contacts_range_suggest
func (r *ApplicationService) GetApplicationVersionContactsRangeSuggest(ctx context.Context, request *GetApplicationVersionContactsRangeSuggestReq, options ...MethodOptionFunc) (*GetApplicationVersionContactsRangeSuggestResp, *Response, error) {
	if r.cli.mock.mockApplicationGetApplicationVersionContactsRangeSuggest != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Application#GetApplicationVersionContactsRangeSuggest mock enable")
		return r.cli.mock.mockApplicationGetApplicationVersionContactsRangeSuggest(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Application",
		API:                   "GetApplicationVersionContactsRangeSuggest",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/application/v6/applications/:app_id/app_versions/:version_id/contacts_range_suggest",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getApplicationVersionContactsRangeSuggestResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApplicationGetApplicationVersionContactsRangeSuggest mock ApplicationGetApplicationVersionContactsRangeSuggest method
func (r *Mock) MockApplicationGetApplicationVersionContactsRangeSuggest(f func(ctx context.Context, request *GetApplicationVersionContactsRangeSuggestReq, options ...MethodOptionFunc) (*GetApplicationVersionContactsRangeSuggestResp, *Response, error)) {
	r.mockApplicationGetApplicationVersionContactsRangeSuggest = f
}

// UnMockApplicationGetApplicationVersionContactsRangeSuggest un-mock ApplicationGetApplicationVersionContactsRangeSuggest method
func (r *Mock) UnMockApplicationGetApplicationVersionContactsRangeSuggest() {
	r.mockApplicationGetApplicationVersionContactsRangeSuggest = nil
}

// GetApplicationVersionContactsRangeSuggestReq ...
type GetApplicationVersionContactsRangeSuggestReq struct {
	AppID            string            `path:"app_id" json:"-"`              // 应用的 AppID, 可以在[开发者后台](https://open.feishu.cn/app) > 凭证与基础信息页查看, * 仅查询本应用信息时, 可填应用自身App ID 或 `me`, * 当值为其他应用的App ID时, 必须申请以下权限: 获取应用信息, 示例值: "cli_9f3ca975326b501b"
	VersionID        string            `path:"version_id" json:"-"`          // 唯一标识应用版本的 ID, 可以调用[获取应用版本列表](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-app_version/list)接口获取, 示例值: "oav_d317f090b7258ad0372aa53963cda70d"
	DepartmentIDType *DepartmentIDType `query:"department_id_type" json:"-"` // 返回值的部门ID的类型, 示例值: department_id, 可选值有: department_id: 以自定义department_id来标识部门, open_department_id: 以open_department_id来标识部门, 默认值: `open_department_id`
	UserIDType       *IDType           `query:"user_id_type" json:"-"`       // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetApplicationVersionContactsRangeSuggestResp ...
type GetApplicationVersionContactsRangeSuggestResp struct {
	ContactsRange *GetApplicationVersionContactsRangeSuggestRespContactsRange `json:"contacts_range,omitempty"` // 应用版本通讯录权限范围建议信息。开发者在提交该版本时如果修改了通讯录权限范围则返回申请的通讯录权限范围。不代表最终应用生效的通讯录权限范围。如果没有修改, 则为空。[如果通讯录权限范围与应用可用范围保持一致, 上次的配置也是如此, 则认为没变化。]
}

// GetApplicationVersionContactsRangeSuggestRespContactsRange ...
type GetApplicationVersionContactsRangeSuggestRespContactsRange struct {
	ContactsScopeType string                                                                 `json:"contacts_scope_type,omitempty"` // 通讯录可见性类型, 可选值有: equal_to_availability: 与应用可用范围一致, 可通过[获取应用在企业内的可用范围](https://open.feishu.cn/document/ukTMukTMukTM/uIjM3UjLyIzN14iMycTN)接口查询具体人员, some: 部分成员, 具体人员参见visible_list, all: 全部成员
	VisibleList       *GetApplicationVersionContactsRangeSuggestRespContactsRangeVisibleList `json:"visible_list,omitempty"`        // 通讯录权限范围的可用名单
}

// GetApplicationVersionContactsRangeSuggestRespContactsRangeVisibleList ...
type GetApplicationVersionContactsRangeSuggestRespContactsRangeVisibleList struct {
	OpenIDs       []string `json:"open_ids,omitempty"`       // 通讯录权限范围的可用成员id列表, 按照user_id_type返回
	DepartmentIDs []string `json:"department_ids,omitempty"` // 通讯录权限范围的可用部门的 id 列表
	GroupIDs      []string `json:"group_ids,omitempty"`      // 通讯录权限范围的可用用户组 id 列表
}

// getApplicationVersionContactsRangeSuggestResp ...
type getApplicationVersionContactsRangeSuggestResp struct {
	Code  int64                                          `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                                         `json:"msg,omitempty"`  // 错误描述
	Data  *GetApplicationVersionContactsRangeSuggestResp `json:"data,omitempty"`
	Error *ErrorDetail                                   `json:"error,omitempty"`
}
