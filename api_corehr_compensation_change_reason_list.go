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

// GetCoreHRCompensationChangeReasonList 批量查询定调薪原因
//
// doc: https://open.larkoffice.com/document/uAjLw4CM/ukTMukTMukTM/compensation-v1/change_reason/list
func (r *CoreHRService) GetCoreHRCompensationChangeReasonList(ctx context.Context, request *GetCoreHRCompensationChangeReasonListReq, options ...MethodOptionFunc) (*GetCoreHRCompensationChangeReasonListResp, *Response, error) {
	if r.cli.mock.mockCoreHRGetCoreHRCompensationChangeReasonList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#GetCoreHRCompensationChangeReasonList mock enable")
		return r.cli.mock.mockCoreHRGetCoreHRCompensationChangeReasonList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "GetCoreHRCompensationChangeReasonList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/compensation/v1/change_reasons",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getCoreHRCompensationChangeReasonListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRGetCoreHRCompensationChangeReasonList mock CoreHRGetCoreHRCompensationChangeReasonList method
func (r *Mock) MockCoreHRGetCoreHRCompensationChangeReasonList(f func(ctx context.Context, request *GetCoreHRCompensationChangeReasonListReq, options ...MethodOptionFunc) (*GetCoreHRCompensationChangeReasonListResp, *Response, error)) {
	r.mockCoreHRGetCoreHRCompensationChangeReasonList = f
}

// UnMockCoreHRGetCoreHRCompensationChangeReasonList un-mock CoreHRGetCoreHRCompensationChangeReasonList method
func (r *Mock) UnMockCoreHRGetCoreHRCompensationChangeReasonList() {
	r.mockCoreHRGetCoreHRCompensationChangeReasonList = nil
}

// GetCoreHRCompensationChangeReasonListReq ...
type GetCoreHRCompensationChangeReasonListReq struct {
	PageSize  int64   `query:"page_size" json:"-"`  // 分页大小, 示例值: 100, 默认值: `100`, 取值范围: `1` ～ `500`
	PageToken *string `query:"page_token" json:"-"` // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: 12314342
}

// GetCoreHRCompensationChangeReasonListResp ...
type GetCoreHRCompensationChangeReasonListResp struct {
	Items     []*GetCoreHRCompensationChangeReasonListRespItem `json:"items,omitempty"`      // 调薪原因信息列表
	PageToken string                                           `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore   bool                                             `json:"has_more,omitempty"`   // 是否还有更多项
}

// GetCoreHRCompensationChangeReasonListRespItem ...
type GetCoreHRCompensationChangeReasonListRespItem struct {
	ID           string                                                   `json:"id,omitempty"`            // 调薪原因ID
	Name         string                                                   `json:"name,omitempty"`          // 调薪原因名称
	Note         string                                                   `json:"note,omitempty"`          // 调薪原因备注
	ActiveStatus int64                                                    `json:"active_status,omitempty"` // 启用状态, 可选值有: 1: 启用, 0: 禁用
	I18nNames    []*GetCoreHRCompensationChangeReasonListRespItemI18nName `json:"i18n_names,omitempty"`    // 多语言名称
	I18nNotes    []*GetCoreHRCompensationChangeReasonListRespItemI18nNote `json:"i18n_notes,omitempty"`    // 多语言描述
}

// GetCoreHRCompensationChangeReasonListRespItemI18nName ...
type GetCoreHRCompensationChangeReasonListRespItemI18nName struct {
	Locale string `json:"locale,omitempty"` // 语言版本, 例如: “zh-CN”、“en-US”
	Value  string `json:"value,omitempty"`  // 语言名称
}

// GetCoreHRCompensationChangeReasonListRespItemI18nNote ...
type GetCoreHRCompensationChangeReasonListRespItemI18nNote struct {
	Locale string `json:"locale,omitempty"` // 语言版本, 例如: “zh-CN”、“en-US”
	Value  string `json:"value,omitempty"`  // 语言名称
}

// getCoreHRCompensationChangeReasonListResp ...
type getCoreHRCompensationChangeReasonListResp struct {
	Code  int64                                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                                     `json:"msg,omitempty"`  // 错误描述
	Data  *GetCoreHRCompensationChangeReasonListResp `json:"data,omitempty"`
	Error *ErrorDetail                               `json:"error,omitempty"`
}
