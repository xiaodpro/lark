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

// CreateContactJobFamily 该接口用于创建租户内的序列信息。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/job_family/create
// new doc: https://open.feishu.cn/document/server-docs/contact-v3/job_family/create
func (r *ContactService) CreateContactJobFamily(ctx context.Context, request *CreateContactJobFamilyReq, options ...MethodOptionFunc) (*CreateContactJobFamilyResp, *Response, error) {
	if r.cli.mock.mockContactCreateContactJobFamily != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Contact#CreateContactJobFamily mock enable")
		return r.cli.mock.mockContactCreateContactJobFamily(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "CreateContactJobFamily",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/job_families",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createContactJobFamilyResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactCreateContactJobFamily mock ContactCreateContactJobFamily method
func (r *Mock) MockContactCreateContactJobFamily(f func(ctx context.Context, request *CreateContactJobFamilyReq, options ...MethodOptionFunc) (*CreateContactJobFamilyResp, *Response, error)) {
	r.mockContactCreateContactJobFamily = f
}

// UnMockContactCreateContactJobFamily un-mock ContactCreateContactJobFamily method
func (r *Mock) UnMockContactCreateContactJobFamily() {
	r.mockContactCreateContactJobFamily = nil
}

// CreateContactJobFamilyReq ...
type CreateContactJobFamilyReq struct {
	Name              string                                      `json:"name,omitempty"`                 // 序列名称。1-100字符, 支持中、英文及符号, 示例值: "产品", 长度范围: `1` ～ `100` 字符
	Description       *string                                     `json:"description,omitempty"`          // 序列描述, 描述序列详情信息, 示例值: "负责产品策略制定的相关工作"
	ParentJobFamilyID *string                                     `json:"parent_job_family_id,omitempty"` // 上级序列ID。需是该租户的序列ID列表中的值, 对应唯一的序列名称, 示例值: "mga5oa8ayjlpzjq"
	Status            bool                                        `json:"status,omitempty"`               // 是否启用, 示例值: true
	I18nName          []*CreateContactJobFamilyReqI18nName        `json:"i18n_name,omitempty"`            // 多语言序列名称
	I18nDescription   []*CreateContactJobFamilyReqI18nDescription `json:"i18n_description,omitempty"`     // 多语言描述
}

// CreateContactJobFamilyReqI18nDescription ...
type CreateContactJobFamilyReqI18nDescription struct {
	Locale *string `json:"locale,omitempty"` // 语言版本, 示例值: "zh_cn"
	Value  *string `json:"value,omitempty"`  // 字段名, 示例值: "多语言内容"
}

// CreateContactJobFamilyReqI18nName ...
type CreateContactJobFamilyReqI18nName struct {
	Locale *string `json:"locale,omitempty"` // 语言版本, 示例值: "zh_cn"
	Value  *string `json:"value,omitempty"`  // 字段名, 示例值: "多语言内容"
}

// CreateContactJobFamilyResp ...
type CreateContactJobFamilyResp struct {
	JobFamily *CreateContactJobFamilyRespJobFamily `json:"job_family,omitempty"` // 序列信息
}

// CreateContactJobFamilyRespJobFamily ...
type CreateContactJobFamilyRespJobFamily struct {
	Name              string                                                `json:"name,omitempty"`                 // 序列名称。1-100字符, 支持中、英文及符号
	Description       string                                                `json:"description,omitempty"`          // 序列描述, 描述序列详情信息
	ParentJobFamilyID string                                                `json:"parent_job_family_id,omitempty"` // 上级序列ID。需是该租户的序列ID列表中的值, 对应唯一的序列名称。
	Status            bool                                                  `json:"status,omitempty"`               // 是否启用
	I18nName          []*CreateContactJobFamilyRespJobFamilyI18nName        `json:"i18n_name,omitempty"`            // 多语言序列名称
	I18nDescription   []*CreateContactJobFamilyRespJobFamilyI18nDescription `json:"i18n_description,omitempty"`     // 多语言描述
	JobFamilyID       string                                                `json:"job_family_id,omitempty"`        // 职级序列ID
}

// CreateContactJobFamilyRespJobFamilyI18nDescription ...
type CreateContactJobFamilyRespJobFamilyI18nDescription struct {
	Locale string `json:"locale,omitempty"` // 语言版本
	Value  string `json:"value,omitempty"`  // 字段名
}

// CreateContactJobFamilyRespJobFamilyI18nName ...
type CreateContactJobFamilyRespJobFamilyI18nName struct {
	Locale string `json:"locale,omitempty"` // 语言版本
	Value  string `json:"value,omitempty"`  // 字段名
}

// createContactJobFamilyResp ...
type createContactJobFamilyResp struct {
	Code  int64                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                      `json:"msg,omitempty"`  // 错误描述
	Data  *CreateContactJobFamilyResp `json:"data,omitempty"`
	Error *ErrorDetail                `json:"error,omitempty"`
}
