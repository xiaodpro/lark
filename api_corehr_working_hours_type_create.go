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

// CreateCoreHrWorkingHoursType 创建工时制度。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/working_hours_type/create
func (r *CoreHrService) CreateCoreHrWorkingHoursType(ctx context.Context, request *CreateCoreHrWorkingHoursTypeReq, options ...MethodOptionFunc) (*CreateCoreHrWorkingHoursTypeResp, *Response, error) {
	if r.cli.mock.mockCoreHrCreateCoreHrWorkingHoursType != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] CoreHr#CreateCoreHrWorkingHoursType mock enable")
		return r.cli.mock.mockCoreHrCreateCoreHrWorkingHoursType(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHr",
		API:                   "CreateCoreHrWorkingHoursType",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/working_hours_types",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createCoreHrWorkingHoursTypeResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHrCreateCoreHrWorkingHoursType mock CoreHrCreateCoreHrWorkingHoursType method
func (r *Mock) MockCoreHrCreateCoreHrWorkingHoursType(f func(ctx context.Context, request *CreateCoreHrWorkingHoursTypeReq, options ...MethodOptionFunc) (*CreateCoreHrWorkingHoursTypeResp, *Response, error)) {
	r.mockCoreHrCreateCoreHrWorkingHoursType = f
}

// UnMockCoreHrCreateCoreHrWorkingHoursType un-mock CoreHrCreateCoreHrWorkingHoursType method
func (r *Mock) UnMockCoreHrCreateCoreHrWorkingHoursType() {
	r.mockCoreHrCreateCoreHrWorkingHoursType = nil
}

// CreateCoreHrWorkingHoursTypeReq ...
type CreateCoreHrWorkingHoursTypeReq struct {
	ClientToken         *string                                       `query:"client_token" json:"-"`           // 根据client_token是否一致来判断是否为同一请求, 示例值: 12454646
	Code                *string                                       `json:"code,omitempty"`                   // 编码, 示例值: "1"
	Name                []*CreateCoreHrWorkingHoursTypeReqName        `json:"name,omitempty"`                   // 名称
	CountryRegionIDList []string                                      `json:"country_region_id_list,omitempty"` // 国家/地区 ID 列表, 示例值: ["4378784343"]
	DefaultForJob       bool                                          `json:"default_for_job,omitempty"`        // 职务默认值, 示例值: true
	Active              bool                                          `json:"active,omitempty"`                 // 是否启用, 示例值: true
	CustomFields        []*CreateCoreHrWorkingHoursTypeReqCustomField `json:"custom_fields,omitempty"`          // 自定义字段
}

// CreateCoreHrWorkingHoursTypeReqCustomField ...
type CreateCoreHrWorkingHoursTypeReqCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名, 示例值: "name"
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05"), 示例值: "\"Sandy\""
}

// CreateCoreHrWorkingHoursTypeReqName ...
type CreateCoreHrWorkingHoursTypeReqName struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言, 示例值: "zh-CN"
	Value string `json:"value,omitempty"` // 名称信息的内容, 示例值: "张三"
}

// CreateCoreHrWorkingHoursTypeResp ...
type CreateCoreHrWorkingHoursTypeResp struct {
	WorkingHoursType *CreateCoreHrWorkingHoursTypeRespWorkingHoursType `json:"working_hours_type,omitempty"` // 工时制度
}

// CreateCoreHrWorkingHoursTypeRespWorkingHoursType ...
type CreateCoreHrWorkingHoursTypeRespWorkingHoursType struct {
	ID                  string                                                         `json:"id,omitempty"`                     // 工时制度 ID
	Code                string                                                         `json:"code,omitempty"`                   // 编码
	Name                []*CreateCoreHrWorkingHoursTypeRespWorkingHoursTypeName        `json:"name,omitempty"`                   // 名称
	CountryRegionIDList []string                                                       `json:"country_region_id_list,omitempty"` // 国家/地区 ID 列表
	DefaultForJob       bool                                                           `json:"default_for_job,omitempty"`        // 职务默认值
	Active              bool                                                           `json:"active,omitempty"`                 // 是否启用
	CustomFields        []*CreateCoreHrWorkingHoursTypeRespWorkingHoursTypeCustomField `json:"custom_fields,omitempty"`          // 自定义字段
}

// CreateCoreHrWorkingHoursTypeRespWorkingHoursTypeCustomField ...
type CreateCoreHrWorkingHoursTypeRespWorkingHoursTypeCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05")
}

// CreateCoreHrWorkingHoursTypeRespWorkingHoursTypeName ...
type CreateCoreHrWorkingHoursTypeRespWorkingHoursTypeName struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// createCoreHrWorkingHoursTypeResp ...
type createCoreHrWorkingHoursTypeResp struct {
	Code int64                             `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                            `json:"msg,omitempty"`  // 错误描述
	Data *CreateCoreHrWorkingHoursTypeResp `json:"data,omitempty"`
}
