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

// UpdateApplicationVersion 通过接口来更新应用版本的审核结果: 通过后应用可以直接上架；拒绝后则开发者可以看到拒绝理由, 并在修改后再次申请发布。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-app_version/patch
func (r *ApplicationService) UpdateApplicationVersion(ctx context.Context, request *UpdateApplicationVersionReq, options ...MethodOptionFunc) (*UpdateApplicationVersionResp, *Response, error) {
	if r.cli.mock.mockApplicationUpdateApplicationVersion != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Application#UpdateApplicationVersion mock enable")
		return r.cli.mock.mockApplicationUpdateApplicationVersion(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Application",
		API:                   "UpdateApplicationVersion",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/application/v6/applications/:app_id/app_versions/:version_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateApplicationVersionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApplicationUpdateApplicationVersion mock ApplicationUpdateApplicationVersion method
func (r *Mock) MockApplicationUpdateApplicationVersion(f func(ctx context.Context, request *UpdateApplicationVersionReq, options ...MethodOptionFunc) (*UpdateApplicationVersionResp, *Response, error)) {
	r.mockApplicationUpdateApplicationVersion = f
}

// UnMockApplicationUpdateApplicationVersion un-mock ApplicationUpdateApplicationVersion method
func (r *Mock) UnMockApplicationUpdateApplicationVersion() {
	r.mockApplicationUpdateApplicationVersion = nil
}

// UpdateApplicationVersionReq ...
type UpdateApplicationVersionReq struct {
	AppID        string  `path:"app_id" json:"-"`         // 应用 id, 示例值: "cli_9f3ca975326b501b"
	VersionID    string  `path:"version_id" json:"-"`     // 唯一标识应用版本的 ID, 示例值: "oav_d317f090b7258ad0372aa53963cda70d"
	UserIDType   IDType  `query:"user_id_type" json:"-"`  // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	OperatorID   string  `query:"operator_id" json:"-"`   // 操作者的 open_id, 示例值: "ou_4065981088f8ef67a504ba8bd6b24d85"
	RejectReason *string `query:"reject_reason" json:"-"` // 当修改版本状态为被驳回时, 这一项必填, 示例值: "拒绝理由"
	Status       *int64  `json:"status,omitempty"`        // 版本状态, 示例值: 1, 可选值有: 0: 未知状态, 1: 审核通过, 2: 审核拒绝, 3: 审核中, 4: 未提交审核
}

// UpdateApplicationVersionResp ...
type UpdateApplicationVersionResp struct {
}

// updateApplicationVersionResp ...
type updateApplicationVersionResp struct {
	Code int64                         `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                        `json:"msg,omitempty"`  // 错误描述
	Data *UpdateApplicationVersionResp `json:"data,omitempty"`
}
