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

// UpdateCoreHRJobLevel 更新职级。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/job_level/patch
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/job-management/job_level/patch
func (r *CoreHRService) UpdateCoreHRJobLevel(ctx context.Context, request *UpdateCoreHRJobLevelReq, options ...MethodOptionFunc) (*UpdateCoreHRJobLevelResp, *Response, error) {
	if r.cli.mock.mockCoreHRUpdateCoreHRJobLevel != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#UpdateCoreHRJobLevel mock enable")
		return r.cli.mock.mockCoreHRUpdateCoreHRJobLevel(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "UpdateCoreHRJobLevel",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/job_levels/:job_level_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateCoreHRJobLevelResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRUpdateCoreHRJobLevel mock CoreHRUpdateCoreHRJobLevel method
func (r *Mock) MockCoreHRUpdateCoreHRJobLevel(f func(ctx context.Context, request *UpdateCoreHRJobLevelReq, options ...MethodOptionFunc) (*UpdateCoreHRJobLevelResp, *Response, error)) {
	r.mockCoreHRUpdateCoreHRJobLevel = f
}

// UnMockCoreHRUpdateCoreHRJobLevel un-mock CoreHRUpdateCoreHRJobLevel method
func (r *Mock) UnMockCoreHRUpdateCoreHRJobLevel() {
	r.mockCoreHRUpdateCoreHRJobLevel = nil
}

// UpdateCoreHRJobLevelReq ...
type UpdateCoreHRJobLevelReq struct {
	JobLevelID   string                                `path:"job_level_id" json:"-"`   // 级别ID, 示例值: "1616161616"
	ClientToken  *string                               `query:"client_token" json:"-"`  // 根据client_token是否一致来判断是否为同一请求, 示例值: 12454646
	LevelOrder   *int64                                `json:"level_order,omitempty"`   // 职级数值, 示例值: 9999
	Code         *string                               `json:"code,omitempty"`          // 编码, 示例值: "VQzo/BSonp8l6PmcZ+VlDhkd2595LMkhyBAGX6HAlCY="
	Name         []*UpdateCoreHRJobLevelReqName        `json:"name,omitempty"`          // 名称
	Description  []*UpdateCoreHRJobLevelReqDescription `json:"description,omitempty"`   // 描述
	Active       *bool                                 `json:"active,omitempty"`        // 是否启用, 示例值: true
	CustomFields []*UpdateCoreHRJobLevelReqCustomField `json:"custom_fields,omitempty"` // 自定义字段
}

// UpdateCoreHRJobLevelReqCustomField ...
type UpdateCoreHRJobLevelReqCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名, 示例值: "name"
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05"), 示例值: "Sandy"
}

// UpdateCoreHRJobLevelReqDescription ...
type UpdateCoreHRJobLevelReqDescription struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言, 示例值: "zh-CN"
	Value string `json:"value,omitempty"` // 名称信息的内容, 示例值: "张三"
}

// UpdateCoreHRJobLevelReqName ...
type UpdateCoreHRJobLevelReqName struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言, 示例值: "zh-CN"
	Value string `json:"value,omitempty"` // 名称信息的内容, 示例值: "张三"
}

// UpdateCoreHRJobLevelResp ...
type UpdateCoreHRJobLevelResp struct {
	JobLevel *UpdateCoreHRJobLevelRespJobLevel `json:"job_level,omitempty"` // 职级
}

// UpdateCoreHRJobLevelRespJobLevel ...
type UpdateCoreHRJobLevelRespJobLevel struct {
	ID           string                                         `json:"id,omitempty"`            // 职级 ID
	LevelOrder   int64                                          `json:"level_order,omitempty"`   // 职级数值
	Code         string                                         `json:"code,omitempty"`          // 编码
	Name         []*UpdateCoreHRJobLevelRespJobLevelName        `json:"name,omitempty"`          // 名称
	Description  []*UpdateCoreHRJobLevelRespJobLevelDescription `json:"description,omitempty"`   // 描述
	Active       bool                                           `json:"active,omitempty"`        // 是否启用
	CustomFields []*UpdateCoreHRJobLevelRespJobLevelCustomField `json:"custom_fields,omitempty"` // 自定义字段
}

// UpdateCoreHRJobLevelRespJobLevelCustomField ...
type UpdateCoreHRJobLevelRespJobLevelCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05")
}

// UpdateCoreHRJobLevelRespJobLevelDescription ...
type UpdateCoreHRJobLevelRespJobLevelDescription struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// UpdateCoreHRJobLevelRespJobLevelName ...
type UpdateCoreHRJobLevelRespJobLevelName struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// updateCoreHRJobLevelResp ...
type updateCoreHRJobLevelResp struct {
	Code  int64                     `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                    `json:"msg,omitempty"`  // 错误描述
	Data  *UpdateCoreHRJobLevelResp `json:"data,omitempty"`
	Error *ErrorDetail              `json:"error,omitempty"`
}
