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

// GetCoreHROffboardingList 查询「飞书人事」-「离职设置」中的离职原因。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/offboarding/query
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/offboarding/query
func (r *CoreHRService) GetCoreHROffboardingList(ctx context.Context, request *GetCoreHROffboardingListReq, options ...MethodOptionFunc) (*GetCoreHROffboardingListResp, *Response, error) {
	if r.cli.mock.mockCoreHRGetCoreHROffboardingList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] CoreHR#GetCoreHROffboardingList mock enable")
		return r.cli.mock.mockCoreHRGetCoreHROffboardingList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "GetCoreHROffboardingList",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/offboardings/query",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getCoreHROffboardingListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRGetCoreHROffboardingList mock CoreHRGetCoreHROffboardingList method
func (r *Mock) MockCoreHRGetCoreHROffboardingList(f func(ctx context.Context, request *GetCoreHROffboardingListReq, options ...MethodOptionFunc) (*GetCoreHROffboardingListResp, *Response, error)) {
	r.mockCoreHRGetCoreHROffboardingList = f
}

// UnMockCoreHRGetCoreHROffboardingList un-mock CoreHRGetCoreHROffboardingList method
func (r *Mock) UnMockCoreHRGetCoreHROffboardingList() {
	r.mockCoreHRGetCoreHROffboardingList = nil
}

// GetCoreHROffboardingListReq ...
type GetCoreHROffboardingListReq struct {
	Active                            *bool    `json:"active,omitempty"`                               // 是否启用, 示例值: true
	OffboardingReasonUniqueIdentifier []string `json:"offboarding_reason_unique_identifier,omitempty"` // 离职原因唯一标识列表, 用于过滤, 最大20个, 示例值: ["offboarding_reason_unique_identifier"]
}

// GetCoreHROffboardingListResp ...
type GetCoreHROffboardingListResp struct {
	Items []*GetCoreHROffboardingListRespItem `json:"items,omitempty"` // 离职原因列表
}

// GetCoreHROffboardingListRespItem ...
type GetCoreHROffboardingListRespItem struct {
	OffboardingReasonUniqueIdentifier       string                                  `json:"offboarding_reason_unique_identifier,omitempty"`        // 离职原因唯一标识
	Name                                    []*GetCoreHROffboardingListRespItemName `json:"name,omitempty"`                                        // 名称
	Active                                  bool                                    `json:"active,omitempty"`                                      // 是否启用, true为启用
	ParentOffboardingReasonUniqueIdentifier string                                  `json:"parent_offboarding_reason_unique_identifier,omitempty"` // 当前离职原因的父级原因唯一标识
	CreatedTime                             string                                  `json:"created_time,omitempty"`                                // 创建时间
	UpdatedTime                             string                                  `json:"updated_time,omitempty"`                                // 更新时间
}

// GetCoreHROffboardingListRespItemName ...
type GetCoreHROffboardingListRespItemName struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// getCoreHROffboardingListResp ...
type getCoreHROffboardingListResp struct {
	Code int64                         `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                        `json:"msg,omitempty"`  // 错误描述
	Data *GetCoreHROffboardingListResp `json:"data,omitempty"`
}
