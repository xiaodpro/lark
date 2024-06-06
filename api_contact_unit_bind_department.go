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

// BindContactUnitDepartment 通过该接口建立部门与单位的绑定关系。由于单位是旗舰版付费功能, 企业需开通相关版本, 否则会绑定失败, 不同版本请参考[飞书版本对比](https://www.feishu.cn/service)。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/unit/bind_department
// new doc: https://open.feishu.cn/document/server-docs/contact-v3/unit/bind_department
func (r *ContactService) BindContactUnitDepartment(ctx context.Context, request *BindContactUnitDepartmentReq, options ...MethodOptionFunc) (*BindContactUnitDepartmentResp, *Response, error) {
	if r.cli.mock.mockContactBindContactUnitDepartment != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Contact#BindContactUnitDepartment mock enable")
		return r.cli.mock.mockContactBindContactUnitDepartment(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "BindContactUnitDepartment",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/unit/bind_department",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(bindContactUnitDepartmentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactBindContactUnitDepartment mock ContactBindContactUnitDepartment method
func (r *Mock) MockContactBindContactUnitDepartment(f func(ctx context.Context, request *BindContactUnitDepartmentReq, options ...MethodOptionFunc) (*BindContactUnitDepartmentResp, *Response, error)) {
	r.mockContactBindContactUnitDepartment = f
}

// UnMockContactBindContactUnitDepartment un-mock ContactBindContactUnitDepartment method
func (r *Mock) UnMockContactBindContactUnitDepartment() {
	r.mockContactBindContactUnitDepartment = nil
}

// BindContactUnitDepartmentReq ...
type BindContactUnitDepartmentReq struct {
	UnitID           string            `json:"unit_id,omitempty"`            // 单位ID, 示例值: "BU121"
	DepartmentID     string            `json:"department_id,omitempty"`      // 单位关联的部门ID, 示例值: "od-4e6ac4d14bcd5071a37a39de902c7141"
	DepartmentIDType *DepartmentIDType `json:"department_id_type,omitempty"` // 此次调用中使用的部门ID的类型, 示例值: "open_department_id", 可选值有: department_id: 以自定义department_id来标识部门, open_department_id: 以open_department_id来标识部门, 默认值: `open_department_id`
}

// BindContactUnitDepartmentResp ...
type BindContactUnitDepartmentResp struct {
}

// bindContactUnitDepartmentResp ...
type bindContactUnitDepartmentResp struct {
	Code  int64                          `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                         `json:"msg,omitempty"`  // 错误描述
	Data  *BindContactUnitDepartmentResp `json:"data,omitempty"`
	Error *ErrorDetail                   `json:"error,omitempty"`
}
