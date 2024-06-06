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

// CancelHireEcoBackgroundCheck 调用此接口将会将系统订单状态变成已终止, 已终止订单将不会响应[「更新背调进度」](https://open.feishu.cn/document/server-docs/hire-v1/ecological-docking/eco_background_check/update_progress)和[「回传背调订单的最终结果」](https://open.feishu.cn/document/server-docs/hire-v1/ecological-docking/eco_background_check/update_result)。 建议在回调终止背调订单之前, 推送账号进度为「已终止」
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/eco_background_check/cancel
// new doc: https://open.feishu.cn/document/server-docs/hire-v1/ecological-docking/eco_background_check/cancel
func (r *HireService) CancelHireEcoBackgroundCheck(ctx context.Context, request *CancelHireEcoBackgroundCheckReq, options ...MethodOptionFunc) (*CancelHireEcoBackgroundCheckResp, *Response, error) {
	if r.cli.mock.mockHireCancelHireEcoBackgroundCheck != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Hire#CancelHireEcoBackgroundCheck mock enable")
		return r.cli.mock.mockHireCancelHireEcoBackgroundCheck(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "CancelHireEcoBackgroundCheck",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/hire/v1/eco_background_checks/cancel",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(cancelHireEcoBackgroundCheckResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHireCancelHireEcoBackgroundCheck mock HireCancelHireEcoBackgroundCheck method
func (r *Mock) MockHireCancelHireEcoBackgroundCheck(f func(ctx context.Context, request *CancelHireEcoBackgroundCheckReq, options ...MethodOptionFunc) (*CancelHireEcoBackgroundCheckResp, *Response, error)) {
	r.mockHireCancelHireEcoBackgroundCheck = f
}

// UnMockHireCancelHireEcoBackgroundCheck un-mock HireCancelHireEcoBackgroundCheck method
func (r *Mock) UnMockHireCancelHireEcoBackgroundCheck() {
	r.mockHireCancelHireEcoBackgroundCheck = nil
}

// CancelHireEcoBackgroundCheckReq ...
type CancelHireEcoBackgroundCheckReq struct {
	BackgroundCheckID string `json:"background_check_id,omitempty"` // 背调 ID, 示例值: "6931286400470354183"
}

// CancelHireEcoBackgroundCheckResp ...
type CancelHireEcoBackgroundCheckResp struct {
}

// cancelHireEcoBackgroundCheckResp ...
type cancelHireEcoBackgroundCheckResp struct {
	Code  int64                             `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                            `json:"msg,omitempty"`  // 错误描述
	Data  *CancelHireEcoBackgroundCheckResp `json:"data,omitempty"`
	Error *ErrorDetail                      `json:"error,omitempty"`
}
