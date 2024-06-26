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

// CreateCoreHREmployment 创建人员的雇佣信息。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/employment/create
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/employee/employment/create
func (r *CoreHRService) CreateCoreHREmployment(ctx context.Context, request *CreateCoreHREmploymentReq, options ...MethodOptionFunc) (*CreateCoreHREmploymentResp, *Response, error) {
	if r.cli.mock.mockCoreHRCreateCoreHREmployment != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#CreateCoreHREmployment mock enable")
		return r.cli.mock.mockCoreHRCreateCoreHREmployment(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "CreateCoreHREmployment",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/employments",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createCoreHREmploymentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRCreateCoreHREmployment mock CoreHRCreateCoreHREmployment method
func (r *Mock) MockCoreHRCreateCoreHREmployment(f func(ctx context.Context, request *CreateCoreHREmploymentReq, options ...MethodOptionFunc) (*CreateCoreHREmploymentResp, *Response, error)) {
	r.mockCoreHRCreateCoreHREmployment = f
}

// UnMockCoreHRCreateCoreHREmployment un-mock CoreHRCreateCoreHREmployment method
func (r *Mock) UnMockCoreHRCreateCoreHREmployment() {
	r.mockCoreHRCreateCoreHREmployment = nil
}

// CreateCoreHREmploymentReq ...
type CreateCoreHREmploymentReq struct {
	ClientToken          *string                                        `query:"client_token" json:"-"`           // 根据client_token是否一致来判断是否为同一请求, 示例值: 12454646
	SeniorityDate        *string                                        `json:"seniority_date,omitempty"`         // 资历起算日期, 示例值: "2020-01-01"
	EmployeeNumber       *string                                        `json:"employee_number,omitempty"`        // 员工编号, 示例值: "1000000"
	EffectiveTime        string                                         `json:"effective_time,omitempty"`         // 入职日期, 示例值: "2020-01-01 00:00:00"
	ExpirationTime       *string                                        `json:"expiration_time,omitempty"`        // 离职日期, 示例值: "2021-01-01"
	EmploymentType       *CreateCoreHREmploymentReqEmploymentType       `json:"employment_type,omitempty"`        // 雇佣类型
	PersonID             string                                         `json:"person_id,omitempty"`              // 人员信息, 引用Person的ID, 示例值: "6919733936050406926"
	PrimaryEmployment    bool                                           `json:"primary_employment,omitempty"`     // 是否是主雇佣信息, 示例值: true
	EmploymentStatus     *CreateCoreHREmploymentReqEmploymentStatus     `json:"employment_status,omitempty"`      // 雇员状态
	CustomFields         []*CreateCoreHREmploymentReqCustomField        `json:"custom_fields,omitempty"`          // 自定义字段
	WorkEmailList        []*CreateCoreHREmploymentReqWorkEmail          `json:"work_email_list,omitempty"`        // 工作邮箱列表, 只有当邮箱下面所有条件时, 才在个人信息页面可见: is_primary = "true", is_public = "true", email_usage = "work"
	ReasonForOffboarding *CreateCoreHREmploymentReqReasonForOffboarding `json:"reason_for_offboarding,omitempty"` // 离职原因
	AtsApplicationID     *string                                        `json:"ats_application_id,omitempty"`     // 招聘投递 ID, 详细信息可以通过[获取投递信息](https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/application/get)接口查询获得, 示例值: "6838119494196871234"
	Rehire               *CreateCoreHREmploymentReqRehire               `json:"rehire,omitempty"`                 // 是否离职重聘, `to_be_confirmed`: 待确认, 系统会判断该员工是否存在历史雇佣记录, 如果存在且需要二次确认时会调用失败, 并返回历史雇佣记录, `no`: 否, 系统直接标为非离职重聘人员, 不再做重复判断, `yes`: 是, 要求历史雇佣信息 ID 必填, 示例值: `no`, 默认值: `to_be_confirmed`
	RehireEmploymentID   *string                                        `json:"rehire_employment_id,omitempty"`   // 历史雇佣信息 ID, 示例值: "7051837122449425964"
}

// CreateCoreHREmploymentReqCustomField ...
type CreateCoreHREmploymentReqCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名, 示例值: "name"
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05"), 示例值: "Sandy"
}

// CreateCoreHREmploymentReqEmploymentStatus ...
type CreateCoreHREmploymentReqEmploymentStatus struct {
	EnumName string `json:"enum_name,omitempty"` // 枚举值, 示例值: "type_1"
}

// CreateCoreHREmploymentReqEmploymentType ...
type CreateCoreHREmploymentReqEmploymentType struct {
	EnumName string `json:"enum_name,omitempty"` // 枚举值, 示例值: "type_1"
}

// CreateCoreHREmploymentReqReasonForOffboarding ...
type CreateCoreHREmploymentReqReasonForOffboarding struct {
	EnumName string `json:"enum_name,omitempty"` // 枚举值, 示例值: "type_1"
}

// CreateCoreHREmploymentReqRehire ...
type CreateCoreHREmploymentReqRehire struct {
	EnumName string `json:"enum_name,omitempty"` // 枚举值, 示例值: "type_1"
}

// CreateCoreHREmploymentReqWorkEmail ...
type CreateCoreHREmploymentReqWorkEmail struct {
	Email        string                                           `json:"email,omitempty"`         // 邮箱号, 示例值: "12456@test.com"
	IsPrimary    *bool                                            `json:"is_primary,omitempty"`    // 是否为主要邮箱, 示例值: true
	IsPublic     *bool                                            `json:"is_public,omitempty"`     // 是否为公开邮箱, 示例值: true
	EmailUsage   *CreateCoreHREmploymentReqWorkEmailEmailUsage    `json:"email_usage,omitempty"`   // 邮箱用途, 枚举值可通过文档[飞书人事枚举常量](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)邮箱用途（email_usage）枚举定义获得
	CustomFields []*CreateCoreHREmploymentReqWorkEmailCustomField `json:"custom_fields,omitempty"` // 自定义字段
}

// CreateCoreHREmploymentReqWorkEmailCustomField ...
type CreateCoreHREmploymentReqWorkEmailCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名, 示例值: "name"
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05"), 示例值: "Sandy"
}

// CreateCoreHREmploymentReqWorkEmailEmailUsage ...
type CreateCoreHREmploymentReqWorkEmailEmailUsage struct {
	EnumName string `json:"enum_name,omitempty"` // 枚举值, 示例值: "type_1"
}

// CreateCoreHREmploymentResp ...
type CreateCoreHREmploymentResp struct {
	Employment *CreateCoreHREmploymentRespEmployment `json:"employment,omitempty"` // 创建人员的雇佣信息成功返回信息
}

// CreateCoreHREmploymentRespEmployment ...
type CreateCoreHREmploymentRespEmployment struct {
	PrehireID            string                                                    `json:"prehire_id,omitempty"`             // 待入职ID
	EmployeeTypeID       string                                                    `json:"employee_type_id,omitempty"`       // 人员类型
	Tenure               string                                                    `json:"tenure,omitempty"`                 // 司龄
	DepartmentID         string                                                    `json:"department_id,omitempty"`          // 部门 ID, 枚举值及详细信息可通过[批量查询部门]接口查询获得
	JobLevelID           string                                                    `json:"job_level_id,omitempty"`           // 职级 ID, 枚举值及详细信息可通过[批量查询职务级别]接口查询获得
	WorkLocationID       string                                                    `json:"work_location_id,omitempty"`       // 工作地点 ID, 枚举值及详细信息可通过[批量查询地点]接口查询获得
	JobFamilyID          string                                                    `json:"job_family_id,omitempty"`          // 职务序列 ID, 枚举值及详细信息可通过[批量查询职务序列]接口查询获得
	JobID                string                                                    `json:"job_id,omitempty"`                 // 职务 ID, 枚举值及详细信息可通过[批量查询职务]接口查询获得
	CompanyID            string                                                    `json:"company_id,omitempty"`             // 法人主体 ID, 枚举值及详细信息可通过[批量查询公司]接口查询获得
	WorkingHoursTypeID   string                                                    `json:"working_hours_type_id,omitempty"`  // 工时制度 ID, 枚举值及详细信息可通过[批量查询工时制度]接口查询获得
	ID                   string                                                    `json:"id,omitempty"`                     // 实体在CoreHR内部的唯一键
	SeniorityDate        string                                                    `json:"seniority_date,omitempty"`         // 资历起算日期
	EmployeeNumber       string                                                    `json:"employee_number,omitempty"`        // 员工编号
	EffectiveTime        string                                                    `json:"effective_time,omitempty"`         // 入职日期
	ExpirationTime       string                                                    `json:"expiration_time,omitempty"`        // 离职日期
	EmploymentType       *CreateCoreHREmploymentRespEmploymentEmploymentType       `json:"employment_type,omitempty"`        // 雇佣类型
	PersonID             string                                                    `json:"person_id,omitempty"`              // 人员信息, 引用Person的ID
	ProbationPeriod      int64                                                     `json:"probation_period,omitempty"`       // 试用期时长
	OnProbation          string                                                    `json:"on_probation,omitempty"`           // 是否在试用期中
	ProbationEndDate     string                                                    `json:"probation_end_date,omitempty"`     // 试用期结束日期
	PrimaryEmployment    bool                                                      `json:"primary_employment,omitempty"`     // 是否是主雇佣信息
	EmploymentStatus     *CreateCoreHREmploymentRespEmploymentEmploymentStatus     `json:"employment_status,omitempty"`      // 雇员状态
	CustomFields         []*CreateCoreHREmploymentRespEmploymentCustomField        `json:"custom_fields,omitempty"`          // 自定义字段
	WorkEmailList        []*CreateCoreHREmploymentRespEmploymentWorkEmail          `json:"work_email_list,omitempty"`        // 工作邮箱列表, 只有当邮箱下面所有条件时, 才在个人信息页面可见: is_primary = "true", is_public = "true", email_usage = "work"
	EmailAddress         string                                                    `json:"email_address,omitempty"`          // 邮箱
	ReasonForOffboarding *CreateCoreHREmploymentRespEmploymentReasonForOffboarding `json:"reason_for_offboarding,omitempty"` // 离职原因
	CostCenterList       []*CreateCoreHREmploymentRespEmploymentCostCenter         `json:"cost_center_list,omitempty"`       // 成本中心id列表
	AtsApplicationID     string                                                    `json:"ats_application_id,omitempty"`     // 招聘投递 ID, 详细信息可以通过[获取投递信息](https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/application/get)接口查询获得
	Rehire               *CreateCoreHREmploymentRespEmploymentRehire               `json:"rehire,omitempty"`                 // 是否离职重聘
	RehireEmploymentID   string                                                    `json:"rehire_employment_id,omitempty"`   // 历史雇佣信息 ID
}

// CreateCoreHREmploymentRespEmploymentCostCenter ...
type CreateCoreHREmploymentRespEmploymentCostCenter struct {
	CostCenterID string                                                       `json:"cost_center_id,omitempty"` // 成本中心id, 可以通过[查询单个成本中心信息]接口获取对应的成本中心信息
	Rate         int64                                                        `json:"rate,omitempty"`           // 分摊比例
	CustomFields []*CreateCoreHREmploymentRespEmploymentCostCenterCustomField `json:"custom_fields,omitempty"`  // 自定义字段
}

// CreateCoreHREmploymentRespEmploymentCostCenterCustomField ...
type CreateCoreHREmploymentRespEmploymentCostCenterCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05")
}

// CreateCoreHREmploymentRespEmploymentCustomField ...
type CreateCoreHREmploymentRespEmploymentCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05")
}

// CreateCoreHREmploymentRespEmploymentEmploymentStatus ...
type CreateCoreHREmploymentRespEmploymentEmploymentStatus struct {
	EnumName string                                                         `json:"enum_name,omitempty"` // 枚举值
	Display  []*CreateCoreHREmploymentRespEmploymentEmploymentStatusDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// CreateCoreHREmploymentRespEmploymentEmploymentStatusDisplay ...
type CreateCoreHREmploymentRespEmploymentEmploymentStatusDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// CreateCoreHREmploymentRespEmploymentEmploymentType ...
type CreateCoreHREmploymentRespEmploymentEmploymentType struct {
	EnumName string                                                       `json:"enum_name,omitempty"` // 枚举值
	Display  []*CreateCoreHREmploymentRespEmploymentEmploymentTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// CreateCoreHREmploymentRespEmploymentEmploymentTypeDisplay ...
type CreateCoreHREmploymentRespEmploymentEmploymentTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// CreateCoreHREmploymentRespEmploymentReasonForOffboarding ...
type CreateCoreHREmploymentRespEmploymentReasonForOffboarding struct {
	EnumName string                                                             `json:"enum_name,omitempty"` // 枚举值
	Display  []*CreateCoreHREmploymentRespEmploymentReasonForOffboardingDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// CreateCoreHREmploymentRespEmploymentReasonForOffboardingDisplay ...
type CreateCoreHREmploymentRespEmploymentReasonForOffboardingDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// CreateCoreHREmploymentRespEmploymentRehire ...
type CreateCoreHREmploymentRespEmploymentRehire struct {
	EnumName string                                               `json:"enum_name,omitempty"` // 枚举值
	Display  []*CreateCoreHREmploymentRespEmploymentRehireDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// CreateCoreHREmploymentRespEmploymentRehireDisplay ...
type CreateCoreHREmploymentRespEmploymentRehireDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// CreateCoreHREmploymentRespEmploymentWorkEmail ...
type CreateCoreHREmploymentRespEmploymentWorkEmail struct {
	Email        string                                                      `json:"email,omitempty"`         // 邮箱号
	IsPrimary    bool                                                        `json:"is_primary,omitempty"`    // 是否为主要邮箱
	IsPublic     bool                                                        `json:"is_public,omitempty"`     // 是否为公开邮箱
	EmailUsage   *CreateCoreHREmploymentRespEmploymentWorkEmailEmailUsage    `json:"email_usage,omitempty"`   // 邮箱用途, 枚举值可通过文档[飞书人事枚举常量](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)邮箱用途（email_usage）枚举定义获得
	CustomFields []*CreateCoreHREmploymentRespEmploymentWorkEmailCustomField `json:"custom_fields,omitempty"` // 自定义字段
}

// CreateCoreHREmploymentRespEmploymentWorkEmailCustomField ...
type CreateCoreHREmploymentRespEmploymentWorkEmailCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05")
}

// CreateCoreHREmploymentRespEmploymentWorkEmailEmailUsage ...
type CreateCoreHREmploymentRespEmploymentWorkEmailEmailUsage struct {
	EnumName string                                                            `json:"enum_name,omitempty"` // 枚举值
	Display  []*CreateCoreHREmploymentRespEmploymentWorkEmailEmailUsageDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// CreateCoreHREmploymentRespEmploymentWorkEmailEmailUsageDisplay ...
type CreateCoreHREmploymentRespEmploymentWorkEmailEmailUsageDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// createCoreHREmploymentResp ...
type createCoreHREmploymentResp struct {
	Code  int64                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                      `json:"msg,omitempty"`  // 错误描述
	Data  *CreateCoreHREmploymentResp `json:"data,omitempty"`
	Error *ErrorDetail                `json:"error,omitempty"`
}
