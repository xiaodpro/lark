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

// UpdateOKRPeriod 修改某个 OKR 周期的状态为「正常」、「失效」或「隐藏」, 对租户所有人生效, 请谨慎操作
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/okr-v1/period/patch
// new doc: https://open.feishu.cn/document/server-docs/okr-v1/period/patch
func (r *OKRService) UpdateOKRPeriod(ctx context.Context, request *UpdateOKRPeriodReq, options ...MethodOptionFunc) (*UpdateOKRPeriodResp, *Response, error) {
	if r.cli.mock.mockOKRUpdateOKRPeriod != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] OKR#UpdateOKRPeriod mock enable")
		return r.cli.mock.mockOKRUpdateOKRPeriod(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "OKR",
		API:                   "UpdateOKRPeriod",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/okr/v1/periods/:period_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateOKRPeriodResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockOKRUpdateOKRPeriod mock OKRUpdateOKRPeriod method
func (r *Mock) MockOKRUpdateOKRPeriod(f func(ctx context.Context, request *UpdateOKRPeriodReq, options ...MethodOptionFunc) (*UpdateOKRPeriodResp, *Response, error)) {
	r.mockOKRUpdateOKRPeriod = f
}

// UnMockOKRUpdateOKRPeriod un-mock OKRUpdateOKRPeriod method
func (r *Mock) UnMockOKRUpdateOKRPeriod() {
	r.mockOKRUpdateOKRPeriod = nil
}

// UpdateOKRPeriodReq ...
type UpdateOKRPeriodReq struct {
	PeriodID string `path:"period_id" json:"-"` // 周期id, 示例值: "6969864184272078374"
	Status   int64  `json:"status,omitempty"`   // 周期显示状态, 示例值: 1, 可选值有: 1: 正常状态, 2: 标记失效, 3: 隐藏周期
}

// UpdateOKRPeriodResp ...
type UpdateOKRPeriodResp struct {
	PeriodID string `json:"period_id,omitempty"` // 周期规则id
	Status   int64  `json:"status,omitempty"`    // 周期显示状态, 可选值有: 1: 正常状态, 2: 标记失效, 3: 隐藏周期
}

// updateOKRPeriodResp ...
type updateOKRPeriodResp struct {
	Code int64                `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string               `json:"msg,omitempty"`  // 错误描述
	Data *UpdateOKRPeriodResp `json:"data,omitempty"`
}
