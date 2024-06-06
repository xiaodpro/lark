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

// GetEmployeeTypeEnumList 该接口用于获取员工的人员类型。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/employee_type_enum/list
// new doc: https://open.feishu.cn/document/server-docs/contact-v3/employee_type_enum/list
func (r *ContactService) GetEmployeeTypeEnumList(ctx context.Context, request *GetEmployeeTypeEnumListReq, options ...MethodOptionFunc) (*GetEmployeeTypeEnumListResp, *Response, error) {
	if r.cli.mock.mockContactGetEmployeeTypeEnumList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Contact#GetEmployeeTypeEnumList mock enable")
		return r.cli.mock.mockContactGetEmployeeTypeEnumList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "GetEmployeeTypeEnumList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/employee_type_enums",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getEmployeeTypeEnumListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactGetEmployeeTypeEnumList mock ContactGetEmployeeTypeEnumList method
func (r *Mock) MockContactGetEmployeeTypeEnumList(f func(ctx context.Context, request *GetEmployeeTypeEnumListReq, options ...MethodOptionFunc) (*GetEmployeeTypeEnumListResp, *Response, error)) {
	r.mockContactGetEmployeeTypeEnumList = f
}

// UnMockContactGetEmployeeTypeEnumList un-mock ContactGetEmployeeTypeEnumList method
func (r *Mock) UnMockContactGetEmployeeTypeEnumList() {
	r.mockContactGetEmployeeTypeEnumList = nil
}

// GetEmployeeTypeEnumListReq ...
type GetEmployeeTypeEnumListReq struct {
	PageToken *string `query:"page_token" json:"-"` // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: 3
	PageSize  *int64  `query:"page_size" json:"-"`  // 分页大小, 示例值: 10, 默认值: `20`, 最大值: `100`
}

// GetEmployeeTypeEnumListResp ...
type GetEmployeeTypeEnumListResp struct {
	Items     []*GetEmployeeTypeEnumListRespItem `json:"items,omitempty"`      // 枚举数据
	HasMore   bool                               `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                             `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
}

// GetEmployeeTypeEnumListRespItem ...
type GetEmployeeTypeEnumListRespItem struct {
	EnumID      string                                        `json:"enum_id,omitempty"`      // 枚举值id
	EnumValue   string                                        `json:"enum_value,omitempty"`   // 枚举的编号值, 创建新的人员类型后, 系统生成对应编号。对应[创建用户接口](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/create)中用户信息的employee_type字段值
	Content     string                                        `json:"content,omitempty"`      // 枚举内容
	EnumType    int64                                         `json:"enum_type,omitempty"`    // 类型, 可选值有: 1: 内置类型, 2: 自定义
	EnumStatus  int64                                         `json:"enum_status,omitempty"`  // 使用状态, 可选值有: 1: 激活, 2: 未激活
	I18nContent []*GetEmployeeTypeEnumListRespItemI18nContent `json:"i18n_content,omitempty"` // i18n定义
}

// GetEmployeeTypeEnumListRespItemI18nContent ...
type GetEmployeeTypeEnumListRespItemI18nContent struct {
	Locale string `json:"locale,omitempty"` // 语言版本
	Value  string `json:"value,omitempty"`  // 字段名
}

// getEmployeeTypeEnumListResp ...
type getEmployeeTypeEnumListResp struct {
	Code  int64                        `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                       `json:"msg,omitempty"`  // 错误描述
	Data  *GetEmployeeTypeEnumListResp `json:"data,omitempty"`
	Error *ErrorDetail                 `json:"error,omitempty"`
}
