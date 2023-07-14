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

// GetTenantProductAssignInfo 获取租户下待分配的席位列表, 包含席位名称、席位ID、数量及对应有效期。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/tenant-v2/tenant-product_assign_info/query
// new doc: https://open.feishu.cn/document/server-docs/tenant-v2/tenant-product_assign_info/query
func (r *TenantService) GetTenantProductAssignInfo(ctx context.Context, request *GetTenantProductAssignInfoReq, options ...MethodOptionFunc) (*GetTenantProductAssignInfoResp, *Response, error) {
	if r.cli.mock.mockTenantGetTenantProductAssignInfo != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Tenant#GetTenantProductAssignInfo mock enable")
		return r.cli.mock.mockTenantGetTenantProductAssignInfo(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Tenant",
		API:                   "GetTenantProductAssignInfo",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/tenant/v2/tenant/assign_info_list/query",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getTenantProductAssignInfoResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockTenantGetTenantProductAssignInfo mock TenantGetTenantProductAssignInfo method
func (r *Mock) MockTenantGetTenantProductAssignInfo(f func(ctx context.Context, request *GetTenantProductAssignInfoReq, options ...MethodOptionFunc) (*GetTenantProductAssignInfoResp, *Response, error)) {
	r.mockTenantGetTenantProductAssignInfo = f
}

// UnMockTenantGetTenantProductAssignInfo un-mock TenantGetTenantProductAssignInfo method
func (r *Mock) UnMockTenantGetTenantProductAssignInfo() {
	r.mockTenantGetTenantProductAssignInfo = nil
}

// GetTenantProductAssignInfoReq ...
type GetTenantProductAssignInfoReq struct {
}

// GetTenantProductAssignInfoResp ...
type GetTenantProductAssignInfoResp struct {
	AssignInfoList []*GetTenantProductAssignInfoRespAssignInfo `json:"assign_info_list,omitempty"` // 租户待分配席位列表
}

// GetTenantProductAssignInfoRespAssignInfo ...
type GetTenantProductAssignInfoRespAssignInfo struct {
	SubscriptionID string                                            `json:"subscription_id,omitempty"`  // 席位id
	LicensePlanKey string                                            `json:"license_plan_key,omitempty"` // license_plan_key
	ProductName    string                                            `json:"product_name,omitempty"`     // 商业化产品名称
	I18nName       *GetTenantProductAssignInfoRespAssignInfoI18nName `json:"i18n_name,omitempty"`        // 国际化名称
	TotalSeats     int64                                             `json:"total_seats,omitempty"`      // 席位总数
	AssignedSeats  string                                            `json:"assigned_seats,omitempty"`   // 已分配席位数
	StartTime      string                                            `json:"start_time,omitempty"`       // 席位起始时间
	EndTime        string                                            `json:"end_time,omitempty"`         // 席位结束时间
}

// GetTenantProductAssignInfoRespAssignInfoI18nName ...
type GetTenantProductAssignInfoRespAssignInfoI18nName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 商业化产品的中文名
	JaJp string `json:"ja_jp,omitempty"` // 商业化产品的日文名
	EnUs string `json:"en_us,omitempty"` // 商业化产品的英文名
}

// getTenantProductAssignInfoResp ...
type getTenantProductAssignInfoResp struct {
	Code int64                           `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                          `json:"msg,omitempty"`  // 错误描述
	Data *GetTenantProductAssignInfoResp `json:"data,omitempty"`
}
