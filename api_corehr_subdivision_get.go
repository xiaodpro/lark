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

// GetCoreHRSubdivision 查询单条省份/行政区信息。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/subdivision/get
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/basic-infomation/location_data/get-2
//
// Deprecated
func (r *CoreHRService) GetCoreHRSubdivision(ctx context.Context, request *GetCoreHRSubdivisionReq, options ...MethodOptionFunc) (*GetCoreHRSubdivisionResp, *Response, error) {
	if r.cli.mock.mockCoreHRGetCoreHRSubdivision != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#GetCoreHRSubdivision mock enable")
		return r.cli.mock.mockCoreHRGetCoreHRSubdivision(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "GetCoreHRSubdivision",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/subdivisions/:subdivision_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getCoreHRSubdivisionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRGetCoreHRSubdivision mock CoreHRGetCoreHRSubdivision method
func (r *Mock) MockCoreHRGetCoreHRSubdivision(f func(ctx context.Context, request *GetCoreHRSubdivisionReq, options ...MethodOptionFunc) (*GetCoreHRSubdivisionResp, *Response, error)) {
	r.mockCoreHRGetCoreHRSubdivision = f
}

// UnMockCoreHRGetCoreHRSubdivision un-mock CoreHRGetCoreHRSubdivision method
func (r *Mock) UnMockCoreHRGetCoreHRSubdivision() {
	r.mockCoreHRGetCoreHRSubdivision = nil
}

// GetCoreHRSubdivisionReq ...
type GetCoreHRSubdivisionReq struct {
	SubdivisionID string `path:"subdivision_id" json:"-"` // 省份/行政区 ID, 示例值: "67489937334909845"
}

// GetCoreHRSubdivisionResp ...
type GetCoreHRSubdivisionResp struct {
	Subdivision *GetCoreHRSubdivisionRespSubdivision `json:"subdivision,omitempty"` // 国家/地址信息
}

// GetCoreHRSubdivisionRespSubdivision ...
type GetCoreHRSubdivisionRespSubdivision struct {
	ID              string                                              `json:"id,omitempty"`                // 省份/行政区id
	Name            []*GetCoreHRSubdivisionRespSubdivisionName          `json:"name,omitempty"`              // 省份/行政区名称
	CountryRegionID string                                              `json:"country_region_id,omitempty"` // 所属国家/地区id, 详细信息可通过[查询国家/地区信息]接口查询获得
	SubdivisionType *GetCoreHRSubdivisionRespSubdivisionSubdivisionType `json:"subdivision_type,omitempty"`  // 行政区类型, 枚举值可通过文档[飞书人事枚举常量]行政区类型（subdivision_type）枚举定义部分获得
}

// GetCoreHRSubdivisionRespSubdivisionName ...
type GetCoreHRSubdivisionRespSubdivisionName struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// GetCoreHRSubdivisionRespSubdivisionSubdivisionType ...
type GetCoreHRSubdivisionRespSubdivisionSubdivisionType struct {
	EnumName string                                                       `json:"enum_name,omitempty"` // 枚举值
	Display  []*GetCoreHRSubdivisionRespSubdivisionSubdivisionTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// GetCoreHRSubdivisionRespSubdivisionSubdivisionTypeDisplay ...
type GetCoreHRSubdivisionRespSubdivisionSubdivisionTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// getCoreHRSubdivisionResp ...
type getCoreHRSubdivisionResp struct {
	Code  int64                     `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                    `json:"msg,omitempty"`  // 错误描述
	Data  *GetCoreHRSubdivisionResp `json:"data,omitempty"`
	Error *ErrorDetail              `json:"error,omitempty"`
}
