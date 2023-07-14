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

// UpdateCoreHREmployment 更新雇佣信息。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/employment/patch
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/employee/employment/patch
func (r *CoreHRService) UpdateCoreHREmployment(ctx context.Context, request *UpdateCoreHREmploymentReq, options ...MethodOptionFunc) (*UpdateCoreHREmploymentResp, *Response, error) {
	if r.cli.mock.mockCoreHRUpdateCoreHREmployment != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] CoreHR#UpdateCoreHREmployment mock enable")
		return r.cli.mock.mockCoreHRUpdateCoreHREmployment(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "UpdateCoreHREmployment",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/employments/:employment_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateCoreHREmploymentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRUpdateCoreHREmployment mock CoreHRUpdateCoreHREmployment method
func (r *Mock) MockCoreHRUpdateCoreHREmployment(f func(ctx context.Context, request *UpdateCoreHREmploymentReq, options ...MethodOptionFunc) (*UpdateCoreHREmploymentResp, *Response, error)) {
	r.mockCoreHRUpdateCoreHREmployment = f
}

// UnMockCoreHRUpdateCoreHREmployment un-mock CoreHRUpdateCoreHREmployment method
func (r *Mock) UnMockCoreHRUpdateCoreHREmployment() {
	r.mockCoreHRUpdateCoreHREmployment = nil
}

// UpdateCoreHREmploymentReq ...
type UpdateCoreHREmploymentReq struct {
	EmploymentID         string                                         `path:"employment_id" json:"-"`           // 雇员ID, 示例值: "1616161616"
	ClientToken          *string                                        `query:"client_token" json:"-"`           // 根据client_token是否一致来判断是否为同一请求, 示例值: 12454646
	UserIDType           *IDType                                        `query:"user_id_type" json:"-"`           // 用户 ID 类型, 示例值: people_corehr_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), people_corehr_id: 以飞书人事的 ID 来识别用户, 默认值: `people_corehr_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	DepartmentIDType     *DepartmentIDType                              `query:"department_id_type" json:"-"`     // 此次调用中使用的部门 ID 类型, 示例值: department_id, 可选值有: open_department_id: 以 open_department_id 来标识部门, department_id: 以 department_id 来标识部门, people_corehr_department_id: 以 people_corehr_department_id 来标识部门, 默认值: `people_corehr_department_id`
	SeniorityDate        *string                                        `json:"seniority_date,omitempty"`         // 资历起算日期, 示例值: "2020-01-01"
	EmployeeNumber       *string                                        `json:"employee_number,omitempty"`        // 员工编号, 示例值: "1000000"
	EffectiveTime        *string                                        `json:"effective_time,omitempty"`         // 入职日期, 示例值: "2020-01-01 00:00:00"
	ExpirationTime       *string                                        `json:"expiration_time,omitempty"`        // 离职日期, 即员工的最后一个工作日, 最后一个工作日时员工的雇佣状态仍为“在职”, 次日凌晨将更改为“离职”, 示例值: "2021-01-01 00:00:00"
	EmploymentType       *UpdateCoreHREmploymentReqEmploymentType       `json:"employment_type,omitempty"`        // 雇佣类型, 枚举值可通过文档[【飞书人事枚举常量】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)雇佣类型（employment_type）枚举定义获得
	PersonID             *string                                        `json:"person_id,omitempty"`              // 个人信息 ID, 详细信息可通过【查询单个个人信息】接口获得, 示例值: "6919733936050406926"
	PrimaryEmployment    *bool                                          `json:"primary_employment,omitempty"`     // 是否是主雇佣信息, 示例值: true
	EmploymentStatus     *UpdateCoreHREmploymentReqEmploymentStatus     `json:"employment_status,omitempty"`      // 雇员状态, 枚举值可通过文档[【飞书人事枚举常量】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)雇员状态（employment_status）枚举定义获得
	CustomFields         []*UpdateCoreHREmploymentReqCustomField        `json:"custom_fields,omitempty"`          // 自定义字段
	WorkEmailList        []*UpdateCoreHREmploymentReqWorkEmail          `json:"work_email_list,omitempty"`        // 工作邮箱列表, 只有当邮箱下面所有条件时, 才在个人信息页面可见: is_primary = "true", is_public = "true", email_usage = "work"
	ReasonForOffboarding *UpdateCoreHREmploymentReqReasonForOffboarding `json:"reason_for_offboarding,omitempty"` // 离职原因, 枚举值可通过文档[【飞书人事枚举常量】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)离职原因（reason_for_offboarding）枚举定义部分获得
	AtsApplicationID     *string                                        `json:"ats_application_id,omitempty"`     // 招聘投递 ID, 详细信息可以通过[【获取投递信息】](https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/application/get)接口查询获得, 示例值: "6838119494196871234"
}

// UpdateCoreHREmploymentReqCustomField ...
type UpdateCoreHREmploymentReqCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名, 示例值: "name"
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05"), 示例值: "\"Sandy\""
}

// UpdateCoreHREmploymentReqEmploymentStatus ...
type UpdateCoreHREmploymentReqEmploymentStatus struct {
	EnumName string `json:"enum_name,omitempty"` // 枚举值, 示例值: "type_1"
}

// UpdateCoreHREmploymentReqEmploymentType ...
type UpdateCoreHREmploymentReqEmploymentType struct {
	EnumName string `json:"enum_name,omitempty"` // 枚举值, 示例值: "type_1"
}

// UpdateCoreHREmploymentReqReasonForOffboarding ...
type UpdateCoreHREmploymentReqReasonForOffboarding struct {
	EnumName string `json:"enum_name,omitempty"` // 枚举值, 示例值: "type_1"
}

// UpdateCoreHREmploymentReqWorkEmail ...
type UpdateCoreHREmploymentReqWorkEmail struct {
	Email        string                                           `json:"email,omitempty"`         // 邮箱号, 示例值: "12456@test.com"
	IsPrimary    *bool                                            `json:"is_primary,omitempty"`    // 是否为主要邮箱, 示例值: true
	IsPublic     *bool                                            `json:"is_public,omitempty"`     // 是否为公开邮箱, 示例值: true
	EmailUsage   *UpdateCoreHREmploymentReqWorkEmailEmailUsage    `json:"email_usage,omitempty"`   // 邮箱用途, 枚举值可通过文档[【飞书人事枚举常量】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)邮箱用途（email_usage）枚举定义获得
	CustomFields []*UpdateCoreHREmploymentReqWorkEmailCustomField `json:"custom_fields,omitempty"` // 自定义字段
}

// UpdateCoreHREmploymentReqWorkEmailCustomField ...
type UpdateCoreHREmploymentReqWorkEmailCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名, 示例值: "name"
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05"), 示例值: "\"Sandy\""
}

// UpdateCoreHREmploymentReqWorkEmailEmailUsage ...
type UpdateCoreHREmploymentReqWorkEmailEmailUsage struct {
	EnumName string `json:"enum_name,omitempty"` // 枚举值, 示例值: "type_1"
}

// UpdateCoreHREmploymentResp ...
type UpdateCoreHREmploymentResp struct {
	Employment *UpdateCoreHREmploymentRespEmployment `json:"employment,omitempty"` // 雇佣信息
}

// UpdateCoreHREmploymentRespEmployment ...
type UpdateCoreHREmploymentRespEmployment struct {
	PrehireID            string                                                    `json:"prehire_id,omitempty"`             // 待入职ID
	EmployeeTypeID       string                                                    `json:"employee_type_id,omitempty"`       // 人员类型
	Tenure               string                                                    `json:"tenure,omitempty"`                 // 司龄
	DepartmentID         string                                                    `json:"department_id,omitempty"`          // 部门 ID, 详细信息可通过【查询单个部门】接口获得
	JobLevelID           string                                                    `json:"job_level_id,omitempty"`           // 职级 ID, 详细信息可通过【查询单个职务级别】接口获得
	WorkLocationID       string                                                    `json:"work_location_id,omitempty"`       // 工作地点 ID, 详细信息可通过【查询单个地点】接口获得
	JobFamilyID          string                                                    `json:"job_family_id,omitempty"`          // 序列 ID, 详细信息可通过【查询单个职务序列】接口获得
	JobID                string                                                    `json:"job_id,omitempty"`                 // 职务 ID, 详细信息可通过【查询单个职务】接口获得
	CompanyID            string                                                    `json:"company_id,omitempty"`             // 法人主体 ID, 详细信息可通过【查询单个公司】接口获得
	WorkingHoursTypeID   string                                                    `json:"working_hours_type_id,omitempty"`  // 工时制度 ID, 详细信息可通过【查询单个工时制度】接口获得
	ID                   string                                                    `json:"id,omitempty"`                     // 雇员ID
	SeniorityDate        string                                                    `json:"seniority_date,omitempty"`         // 资历起算日期
	EmployeeNumber       string                                                    `json:"employee_number,omitempty"`        // 员工编号
	EffectiveTime        string                                                    `json:"effective_time,omitempty"`         // 入职日期
	ExpirationTime       string                                                    `json:"expiration_time,omitempty"`        // 离职日期, 即员工的最后一个工作日, 最后一个工作日时员工的雇佣状态仍为“在职”, 次日凌晨将更改为“离职”
	EmploymentType       *UpdateCoreHREmploymentRespEmploymentEmploymentType       `json:"employment_type,omitempty"`        // 雇佣类型, 枚举值可通过文档[【飞书人事枚举常量】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)雇佣类型（employment_type）枚举定义获得
	PersonID             string                                                    `json:"person_id,omitempty"`              // 个人信息 ID, 详细信息可通过【查询单个个人信息】接口获得
	ProbationPeriod      int64                                                     `json:"probation_period,omitempty"`       // 试用期时长（月份）
	OnProbation          string                                                    `json:"on_probation,omitempty"`           // 是否在试用期中, 满足以下任一条件时, 该字段值为`"true"`: 预计试用结束日期非空, 且实际结束日期为空, 预计试用结束日期非空, 实际结束日期非空, 且当日日期小于等于实际结束日期, 其余情况下, 该字段值为`"false"`；
	ProbationEndDate     string                                                    `json:"probation_end_date,omitempty"`     // 试用期结束日期（实际结束日期）
	PrimaryEmployment    bool                                                      `json:"primary_employment,omitempty"`     // 是否是主雇佣信息
	EmploymentStatus     *UpdateCoreHREmploymentRespEmploymentEmploymentStatus     `json:"employment_status,omitempty"`      // 雇员状态, 枚举值可通过文档[【飞书人事枚举常量】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)雇员状态（employment_status）枚举定义获得
	CustomFields         []*UpdateCoreHREmploymentRespEmploymentCustomField        `json:"custom_fields,omitempty"`          // 自定义字段
	WorkEmailList        []*UpdateCoreHREmploymentRespEmploymentWorkEmail          `json:"work_email_list,omitempty"`        // 工作邮箱列表, 只有当邮箱下面所有条件时, 才在个人信息页面可见: is_primary = "true", is_public = "true", email_usage = "work"
	EmailAddress         string                                                    `json:"email_address,omitempty"`          // 邮箱
	ReasonForOffboarding *UpdateCoreHREmploymentRespEmploymentReasonForOffboarding `json:"reason_for_offboarding,omitempty"` // 离职原因, 枚举值可通过文档[【飞书人事枚举常量】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)离职原因（reason_for_offboarding）枚举定义部分获得
	CostCenterList       []*UpdateCoreHREmploymentRespEmploymentCostCenter         `json:"cost_center_list,omitempty"`       // 成本中心id列表
	AtsApplicationID     string                                                    `json:"ats_application_id,omitempty"`     // 招聘投递 ID, 详细信息可以通过[【获取投递信息】](https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/application/get)接口查询获得
}

// UpdateCoreHREmploymentRespEmploymentCostCenter ...
type UpdateCoreHREmploymentRespEmploymentCostCenter struct {
	CostCenterID string                                                       `json:"cost_center_id,omitempty"` // 成本中心id, 可以通过【查询单个成本中心信息】接口获取对应的成本中心信息
	Rate         int64                                                        `json:"rate,omitempty"`           // 分摊比例
	CustomFields []*UpdateCoreHREmploymentRespEmploymentCostCenterCustomField `json:"custom_fields,omitempty"`  // 自定义字段
}

// UpdateCoreHREmploymentRespEmploymentCostCenterCustomField ...
type UpdateCoreHREmploymentRespEmploymentCostCenterCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05")
}

// UpdateCoreHREmploymentRespEmploymentCustomField ...
type UpdateCoreHREmploymentRespEmploymentCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05")
}

// UpdateCoreHREmploymentRespEmploymentEmploymentStatus ...
type UpdateCoreHREmploymentRespEmploymentEmploymentStatus struct {
	EnumName string                                                         `json:"enum_name,omitempty"` // 枚举值
	Display  []*UpdateCoreHREmploymentRespEmploymentEmploymentStatusDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// UpdateCoreHREmploymentRespEmploymentEmploymentStatusDisplay ...
type UpdateCoreHREmploymentRespEmploymentEmploymentStatusDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// UpdateCoreHREmploymentRespEmploymentEmploymentType ...
type UpdateCoreHREmploymentRespEmploymentEmploymentType struct {
	EnumName string                                                       `json:"enum_name,omitempty"` // 枚举值
	Display  []*UpdateCoreHREmploymentRespEmploymentEmploymentTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// UpdateCoreHREmploymentRespEmploymentEmploymentTypeDisplay ...
type UpdateCoreHREmploymentRespEmploymentEmploymentTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// UpdateCoreHREmploymentRespEmploymentReasonForOffboarding ...
type UpdateCoreHREmploymentRespEmploymentReasonForOffboarding struct {
	EnumName string                                                             `json:"enum_name,omitempty"` // 枚举值
	Display  []*UpdateCoreHREmploymentRespEmploymentReasonForOffboardingDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// UpdateCoreHREmploymentRespEmploymentReasonForOffboardingDisplay ...
type UpdateCoreHREmploymentRespEmploymentReasonForOffboardingDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// UpdateCoreHREmploymentRespEmploymentWorkEmail ...
type UpdateCoreHREmploymentRespEmploymentWorkEmail struct {
	Email        string                                                      `json:"email,omitempty"`         // 邮箱号
	IsPrimary    bool                                                        `json:"is_primary,omitempty"`    // 是否为主要邮箱
	IsPublic     bool                                                        `json:"is_public,omitempty"`     // 是否为公开邮箱
	EmailUsage   *UpdateCoreHREmploymentRespEmploymentWorkEmailEmailUsage    `json:"email_usage,omitempty"`   // 邮箱用途, 枚举值可通过文档[【飞书人事枚举常量】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)邮箱用途（email_usage）枚举定义获得
	CustomFields []*UpdateCoreHREmploymentRespEmploymentWorkEmailCustomField `json:"custom_fields,omitempty"` // 自定义字段
}

// UpdateCoreHREmploymentRespEmploymentWorkEmailCustomField ...
type UpdateCoreHREmploymentRespEmploymentWorkEmailCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05")
}

// UpdateCoreHREmploymentRespEmploymentWorkEmailEmailUsage ...
type UpdateCoreHREmploymentRespEmploymentWorkEmailEmailUsage struct {
	EnumName string                                                            `json:"enum_name,omitempty"` // 枚举值
	Display  []*UpdateCoreHREmploymentRespEmploymentWorkEmailEmailUsageDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// UpdateCoreHREmploymentRespEmploymentWorkEmailEmailUsageDisplay ...
type UpdateCoreHREmploymentRespEmploymentWorkEmailEmailUsageDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// updateCoreHREmploymentResp ...
type updateCoreHREmploymentResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *UpdateCoreHREmploymentResp `json:"data,omitempty"`
}
