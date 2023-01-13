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

// UpdateHireEmployee 根据员工 ID 更新员工转正、离职状态。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/employee/patch
func (r *HireService) UpdateHireEmployee(ctx context.Context, request *UpdateHireEmployeeReq, options ...MethodOptionFunc) (*UpdateHireEmployeeResp, *Response, error) {
	if r.cli.mock.mockHireUpdateHireEmployee != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Hire#UpdateHireEmployee mock enable")
		return r.cli.mock.mockHireUpdateHireEmployee(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "UpdateHireEmployee",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/hire/v1/employees/:employee_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateHireEmployeeResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHireUpdateHireEmployee mock HireUpdateHireEmployee method
func (r *Mock) MockHireUpdateHireEmployee(f func(ctx context.Context, request *UpdateHireEmployeeReq, options ...MethodOptionFunc) (*UpdateHireEmployeeResp, *Response, error)) {
	r.mockHireUpdateHireEmployee = f
}

// UnMockHireUpdateHireEmployee un-mock HireUpdateHireEmployee method
func (r *Mock) UnMockHireUpdateHireEmployee() {
	r.mockHireUpdateHireEmployee = nil
}

// UpdateHireEmployeeReq ...
type UpdateHireEmployeeReq struct {
	EmployeeID     string                               `path:"employee_id" json:"-"`      // 员工ID, 示例值: "123"
	UserIDType     *IDType                              `query:"user_id_type" json:"-"`    // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	Operation      int64                                `json:"operation,omitempty"`       // 修改状态操作, 示例值: 1, 可选值有: 1: 转正, 2: 离职
	ConversionInfo *UpdateHireEmployeeReqConversionInfo `json:"conversion_info,omitempty"` // 转正信息
	OverboardInfo  *UpdateHireEmployeeReqOverboardInfo  `json:"overboard_info,omitempty"`  // 离职信息
}

// UpdateHireEmployeeReqConversionInfo ...
type UpdateHireEmployeeReqConversionInfo struct {
	ActualConversionTime *int64 `json:"actual_conversion_time,omitempty"` // 实际转正日期, 示例值: 1637596800000
}

// UpdateHireEmployeeReqOverboardInfo ...
type UpdateHireEmployeeReqOverboardInfo struct {
	ActualOverboardTime *int64  `json:"actual_overboard_time,omitempty"` // 实际离职日期, 示例值: 1637596800000
	OverboardNote       *string `json:"overboard_note,omitempty"`        // 离职原因, 示例值: "职业发展考虑"
}

// UpdateHireEmployeeResp ...
type UpdateHireEmployeeResp struct {
	Employee *UpdateHireEmployeeRespEmployee `json:"employee,omitempty"` // 员工信息
}

// UpdateHireEmployeeRespEmployee ...
type UpdateHireEmployeeRespEmployee struct {
	ID                     string       `json:"id,omitempty"`                       // 员工ID
	ApplicationID          string       `json:"application_id,omitempty"`           // 投递ID
	OnboardStatus          int64        `json:"onboard_status,omitempty"`           // 入职状态, 可选值有: 1: 已入职, 2: 已离职
	ConversionStatus       int64        `json:"conversion_status,omitempty"`        // 转正状态, 可选值有: 1: 未转正, 2: 已转正
	OnboardTime            int64        `json:"onboard_time,omitempty"`             // 实际入职时间
	ExpectedConversionTime int64        `json:"expected_conversion_time,omitempty"` // 预期转正时间
	ActualConversionTime   int64        `json:"actual_conversion_time,omitempty"`   // 实际转正时间
	OverboardTime          int64        `json:"overboard_time,omitempty"`           // 离职时间
	OverboardNote          string       `json:"overboard_note,omitempty"`           // 离职原因
	OnboardCityCode        string       `json:"onboard_city_code,omitempty"`        // 办公地点
	Department             string       `json:"department,omitempty"`               // 入职部门
	Leader                 string       `json:"leader,omitempty"`                   // 直属上级
	Sequence               string       `json:"sequence,omitempty"`                 // 序列
	Level                  string       `json:"level,omitempty"`                    // 职级
	EmployeeType           EmployeeType `json:"employee_type,omitempty"`            // 员工类型
}

// updateHireEmployeeResp ...
type updateHireEmployeeResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *UpdateHireEmployeeResp `json:"data,omitempty"`
}
