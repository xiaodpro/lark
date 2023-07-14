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

// GetHireReferralByApplication 根据投递 ID 获取内推信息。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/referral/get_by_application
// new doc: https://open.feishu.cn/document/server-docs/hire-v1/get-candidates/referral/get_by_application
func (r *HireService) GetHireReferralByApplication(ctx context.Context, request *GetHireReferralByApplicationReq, options ...MethodOptionFunc) (*GetHireReferralByApplicationResp, *Response, error) {
	if r.cli.mock.mockHireGetHireReferralByApplication != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Hire#GetHireReferralByApplication mock enable")
		return r.cli.mock.mockHireGetHireReferralByApplication(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "GetHireReferralByApplication",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/hire/v1/referrals/get_by_application",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getHireReferralByApplicationResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHireGetHireReferralByApplication mock HireGetHireReferralByApplication method
func (r *Mock) MockHireGetHireReferralByApplication(f func(ctx context.Context, request *GetHireReferralByApplicationReq, options ...MethodOptionFunc) (*GetHireReferralByApplicationResp, *Response, error)) {
	r.mockHireGetHireReferralByApplication = f
}

// UnMockHireGetHireReferralByApplication un-mock HireGetHireReferralByApplication method
func (r *Mock) UnMockHireGetHireReferralByApplication() {
	r.mockHireGetHireReferralByApplication = nil
}

// GetHireReferralByApplicationReq ...
type GetHireReferralByApplicationReq struct {
	ApplicationID string  `query:"application_id" json:"-"` // 投递的 ID, 示例值: 6134134355464633
	UserIDType    *IDType `query:"user_id_type" json:"-"`   // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), people_admin_id: 以people_admin_id来识别用户, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetHireReferralByApplicationResp ...
type GetHireReferralByApplicationResp struct {
	Referral *GetHireReferralByApplicationRespReferral `json:"referral,omitempty"` // 内推信息
}

// GetHireReferralByApplicationRespReferral ...
type GetHireReferralByApplicationRespReferral struct {
	ID             string                                                `json:"id,omitempty"`               // 内推的 ID
	ApplicationID  string                                                `json:"application_id,omitempty"`   // 投递 ID
	CreateTime     int64                                                 `json:"create_time,omitempty"`      // 创建时间（ms）
	ReferralUserID string                                                `json:"referral_user_id,omitempty"` // 内推人的 ID
	ReferralUser   *GetHireReferralByApplicationRespReferralReferralUser `json:"referral_user,omitempty"`    // 内推人信息
}

// GetHireReferralByApplicationRespReferralReferralUser ...
type GetHireReferralByApplicationRespReferralReferralUser struct {
	ID   string                                                    `json:"id,omitempty"`   // ID
	Name *GetHireReferralByApplicationRespReferralReferralUserName `json:"name,omitempty"` // 名称
}

// GetHireReferralByApplicationRespReferralReferralUserName ...
type GetHireReferralByApplicationRespReferralReferralUserName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// getHireReferralByApplicationResp ...
type getHireReferralByApplicationResp struct {
	Code int64                             `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                            `json:"msg,omitempty"`  // 错误描述
	Data *GetHireReferralByApplicationResp `json:"data,omitempty"`
}
