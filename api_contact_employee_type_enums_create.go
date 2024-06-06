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

// CreateEmployeeTypeEnum 新增自定义人员类型。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/employee_type_enum/create
// new doc: https://open.feishu.cn/document/server-docs/contact-v3/employee_type_enum/create
func (r *ContactService) CreateEmployeeTypeEnum(ctx context.Context, request *CreateEmployeeTypeEnumReq, options ...MethodOptionFunc) (*CreateEmployeeTypeEnumResp, *Response, error) {
	if r.cli.mock.mockContactCreateEmployeeTypeEnum != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Contact#CreateEmployeeTypeEnum mock enable")
		return r.cli.mock.mockContactCreateEmployeeTypeEnum(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "CreateEmployeeTypeEnum",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/employee_type_enums",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createEmployeeTypeEnumResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactCreateEmployeeTypeEnum mock ContactCreateEmployeeTypeEnum method
func (r *Mock) MockContactCreateEmployeeTypeEnum(f func(ctx context.Context, request *CreateEmployeeTypeEnumReq, options ...MethodOptionFunc) (*CreateEmployeeTypeEnumResp, *Response, error)) {
	r.mockContactCreateEmployeeTypeEnum = f
}

// UnMockContactCreateEmployeeTypeEnum un-mock ContactCreateEmployeeTypeEnum method
func (r *Mock) UnMockContactCreateEmployeeTypeEnum() {
	r.mockContactCreateEmployeeTypeEnum = nil
}

// CreateEmployeeTypeEnumReq ...
type CreateEmployeeTypeEnumReq struct {
	Content     string                                  `json:"content,omitempty"`      // 枚举内容, 示例值: "专家", 长度范围: `1` ～ `100` 字符
	EnumType    int64                                   `json:"enum_type,omitempty"`    // 类型, 示例值: 2, 可选值有: 1: 内置类型, 2: 自定义
	EnumStatus  int64                                   `json:"enum_status,omitempty"`  // 使用状态, 示例值: 1, 可选值有: 1: 激活, 2: 未激活
	I18nContent []*CreateEmployeeTypeEnumReqI18nContent `json:"i18n_content,omitempty"` // i18n定义
}

// CreateEmployeeTypeEnumReqI18nContent ...
type CreateEmployeeTypeEnumReqI18nContent struct {
	Locale *string `json:"locale,omitempty"` // 语言版本, 示例值: "zh_cn"
	Value  *string `json:"value,omitempty"`  // 字段名, 示例值: "专家（中文）"
}

// CreateEmployeeTypeEnumResp ...
type CreateEmployeeTypeEnumResp struct {
	EmployeeTypeEnum *CreateEmployeeTypeEnumRespEmployeeTypeEnum `json:"employee_type_enum,omitempty"` // 新建的人员类型信息
}

// CreateEmployeeTypeEnumRespEmployeeTypeEnum ...
type CreateEmployeeTypeEnumRespEmployeeTypeEnum struct {
	EnumID      string                                                   `json:"enum_id,omitempty"`      // 枚举值id
	EnumValue   string                                                   `json:"enum_value,omitempty"`   // 枚举的编号值, 创建新的人员类型后, 系统生成对应编号。对应[创建用户接口](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/create)中用户信息的employee_type字段值
	Content     string                                                   `json:"content,omitempty"`      // 枚举内容
	EnumType    int64                                                    `json:"enum_type,omitempty"`    // 类型, 可选值有: 1: 内置类型, 2: 自定义
	EnumStatus  int64                                                    `json:"enum_status,omitempty"`  // 使用状态, 可选值有: 1: 激活, 2: 未激活
	I18nContent []*CreateEmployeeTypeEnumRespEmployeeTypeEnumI18nContent `json:"i18n_content,omitempty"` // i18n定义
}

// CreateEmployeeTypeEnumRespEmployeeTypeEnumI18nContent ...
type CreateEmployeeTypeEnumRespEmployeeTypeEnumI18nContent struct {
	Locale string `json:"locale,omitempty"` // 语言版本
	Value  string `json:"value,omitempty"`  // 字段名
}

// createEmployeeTypeEnumResp ...
type createEmployeeTypeEnumResp struct {
	Code  int64                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                      `json:"msg,omitempty"`  // 错误描述
	Data  *CreateEmployeeTypeEnumResp `json:"data,omitempty"`
	Error *ErrorDetail                `json:"error,omitempty"`
}
