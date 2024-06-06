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

// BatchGetCoreHRJobData 通过员工雇佣 ID 批量查询任职信息。
//
// 该接口会按照应用拥有的「员工资源」的权限范围返回数据, 请确定在「开发者后台 - 权限管理 - 数据权限」中已申请「员工资源」权限范围。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/employees-job_data/batch_get
func (r *CoreHRService) BatchGetCoreHRJobData(ctx context.Context, request *BatchGetCoreHRJobDataReq, options ...MethodOptionFunc) (*BatchGetCoreHRJobDataResp, *Response, error) {
	if r.cli.mock.mockCoreHRBatchGetCoreHRJobData != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#BatchGetCoreHRJobData mock enable")
		return r.cli.mock.mockCoreHRBatchGetCoreHRJobData(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "BatchGetCoreHRJobData",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v2/employees/job_datas/batch_get",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(batchGetCoreHRJobDataResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRBatchGetCoreHRJobData mock CoreHRBatchGetCoreHRJobData method
func (r *Mock) MockCoreHRBatchGetCoreHRJobData(f func(ctx context.Context, request *BatchGetCoreHRJobDataReq, options ...MethodOptionFunc) (*BatchGetCoreHRJobDataResp, *Response, error)) {
	r.mockCoreHRBatchGetCoreHRJobData = f
}

// UnMockCoreHRBatchGetCoreHRJobData un-mock CoreHRBatchGetCoreHRJobData method
func (r *Mock) UnMockCoreHRBatchGetCoreHRJobData() {
	r.mockCoreHRBatchGetCoreHRJobData = nil
}

// BatchGetCoreHRJobDataReq ...
type BatchGetCoreHRJobDataReq struct {
	UserIDType         *IDType           `query:"user_id_type" json:"-"`         // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), people_corehr_id: 以飞书人事的 ID 来识别用户, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	DepartmentIDType   *DepartmentIDType `query:"department_id_type" json:"-"`   // 此次调用中使用的部门 ID 类型, 示例值: people_corehr_department_id, 可选值有: open_department_id: 以 open_department_id 来标识部门, department_id: 以 department_id 来标识部门, people_corehr_department_id: 以 people_corehr_department_id 来标识部门, 默认值: `people_corehr_department_id`
	EmploymentIDs      []string          `json:"employment_ids,omitempty"`       // 员工雇佣 ID 列表, 示例值: ["7140964208476371111"], 长度范围: `1` ～ `100`
	GetAllVersion      *bool             `json:"get_all_version,omitempty"`      // 是否获取所有任职记录, true 为获取员工所有版本的任职记录, false 为仅获取当前生效的任职记录, 默认为 false, 示例值: false
	EffectiveDateStart *string           `json:"effective_date_start,omitempty"` // 生效日期 - 搜索范围开始, 示例值: "2020-01-01"
	EffectiveDateEnd   *string           `json:"effective_date_end,omitempty"`   // 生效日期 - 搜索范围结束, 示例值: "2020-01-01"
	DataDate           *string           `json:"data_date,omitempty"`            // 查看数据日期, 默认为今天, 示例值: "2020-01-01"
}

// BatchGetCoreHRJobDataResp ...
type BatchGetCoreHRJobDataResp struct {
	Items []*BatchGetCoreHRJobDataRespItem `json:"items,omitempty"` // 查询的雇佣信息
}

// BatchGetCoreHRJobDataRespItem ...
type BatchGetCoreHRJobDataRespItem struct {
	EmploymentID string                                  `json:"employment_id,omitempty"` // Employment ID
	JobDatas     []*BatchGetCoreHRJobDataRespItemJobData `json:"job_datas,omitempty"`     // 实体在 CoreHR 内部的唯一键
}

// BatchGetCoreHRJobDataRespItemJobData ...
type BatchGetCoreHRJobDataRespItemJobData struct {
	JobDataID                string                                                     `json:"job_data_id,omitempty"`                 // 任职信息 ID
	VersionID                string                                                     `json:"version_id,omitempty"`                  // 任职记录版本 ID
	EmployeeTypeID           string                                                     `json:"employee_type_id,omitempty"`            // 人员类型 ID, 枚举值及详细信息可通过[查询单个人员类型](https://open.feishu.cn/document/server-docs/corehr-v1/basic-infomation/employee_type/get)接口查询获得
	WorkingHoursTypeID       string                                                     `json:"working_hours_type_id,omitempty"`       // 工时制度 ID, 枚举值及详细信息可通过[查询单个工时制度](https://open.feishu.cn/document/server-docs/corehr-v1/basic-infomation/working_hours_type/get)接口查询获得
	WorkLocationID           string                                                     `json:"work_location_id,omitempty"`            // 工作地点 ID, 枚举值及详细信息可通过[查询单个地点](https://open.feishu.cn/document/server-docs/corehr-v1/organization-management/location/get)接口查询获得
	DepartmentID             string                                                     `json:"department_id,omitempty"`               // 部门 ID, 枚举值及详细信息可通过[查询单个部门](https://open.feishu.cn/document/server-docs/corehr-v1/organization-management/department/get)接口查询获得
	JobID                    string                                                     `json:"job_id,omitempty"`                      // 职务 ID, 枚举值及详细信息可通过[查询单个职务](https://open.feishu.cn/document/server-docs/corehr-v1/job-management/job/get)接口查询获得, 字段权限要求（满足任一）: 获取员工的职务信息, 获取职务级别信息, 读写员工的职务级别信息
	JobLevelID               string                                                     `json:"job_level_id,omitempty"`                // 职级 ID, 枚举值及详细信息可通过[查询单个职级](https://open.feishu.cn/document/server-docs/corehr-v1/job-management/job_level/get)接口查询获得, 字段权限要求（满足任一）: 获取职务级别信息, 读写员工的职务级别信息
	JobGradeID               string                                                     `json:"job_grade_id,omitempty"`                // 职等 ID, 字段权限要求（满足任一）: 获取职等信息, 读写职等信息
	JobFamilyID              string                                                     `json:"job_family_id,omitempty"`               // 序列 ID, 枚举值及详细信息可通过[查询单个序列](https://open.feishu.cn/document/server-docs/corehr-v1/job-management/job_family/get)接口查询获得
	ProbationStartDate       string                                                     `json:"probation_start_date,omitempty"`        // 试用期开始日期
	ProbationEndDate         string                                                     `json:"probation_end_date,omitempty"`          // 试用期结束日期（实际结束日期）
	PrimaryJobData           bool                                                       `json:"primary_job_data,omitempty"`            // 是否为主任职
	EmploymentID             string                                                     `json:"employment_id,omitempty"`               // 雇佣 ID
	EffectiveTime            string                                                     `json:"effective_time,omitempty"`              // 生效时间
	ExpirationTime           string                                                     `json:"expiration_time,omitempty"`             // 失效时间
	AssignmentStartReason    *BatchGetCoreHRJobDataRespItemJobDataAssignmentStartReason `json:"assignment_start_reason,omitempty"`     // 任职原因, 枚举值可通过文档[枚举常量介绍](https://open.feishu.cn/document/server-docs/corehr-v1/feishu-people-enum-constant)任职原因（assignment_start_reason）枚举定义部分获得, 字段权限要求: 查看任职记录的原因字段
	ProbationExpectedEndDate string                                                     `json:"probation_expected_end_date,omitempty"` // 预计试用期结束日期
	ProbationOutcome         *BatchGetCoreHRJobDataRespItemJobDataProbationOutcome      `json:"probation_outcome,omitempty"`           // 试用期结果, 枚举值可通过文档[枚举常量介绍](https://open.feishu.cn/document/server-docs/corehr-v1/feishu-people-enum-constant)试用期结果（probation_outcome）枚举定义部分获得
	DirectManager            *BatchGetCoreHRJobDataRespItemJobDataDirectManager         `json:"direct_manager,omitempty"`              // 直属上级
	DottedLineManagers       []*BatchGetCoreHRJobDataRespItemJobDataDottedLineManager   `json:"dotted_line_managers,omitempty"`        // 虚线上级
	SecondDirectManager      *BatchGetCoreHRJobDataRespItemJobDataSecondDirectManager   `json:"second_direct_manager,omitempty"`       // 第二实线主管
	CostCenterRates          []*BatchGetCoreHRJobDataRespItemJobDataCostCenterRate      `json:"cost_center_rates,omitempty"`           // 成本中心分摊信息
	WorkShift                *BatchGetCoreHRJobDataRespItemJobDataWorkShift             `json:"work_shift,omitempty"`                  // 排班类型, 字段权限要求: 获取排班信息
	CompensationType         *BatchGetCoreHRJobDataRespItemJobDataCompensationType      `json:"compensation_type,omitempty"`           // 薪资类型, 字段权限要求: 获取薪资类型
	ServiceCompany           string                                                     `json:"service_company,omitempty"`             // 任职公司, 字段权限要求: 获取任职公司
}

// BatchGetCoreHRJobDataRespItemJobDataAssignmentStartReason ...
type BatchGetCoreHRJobDataRespItemJobDataAssignmentStartReason struct {
	EnumName string                                                              `json:"enum_name,omitempty"` // 枚举值
	Display  []*BatchGetCoreHRJobDataRespItemJobDataAssignmentStartReasonDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// BatchGetCoreHRJobDataRespItemJobDataAssignmentStartReasonDisplay ...
type BatchGetCoreHRJobDataRespItemJobDataAssignmentStartReasonDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// BatchGetCoreHRJobDataRespItemJobDataCompensationType ...
type BatchGetCoreHRJobDataRespItemJobDataCompensationType struct {
	EnumName string                                                         `json:"enum_name,omitempty"` // 枚举值
	Display  []*BatchGetCoreHRJobDataRespItemJobDataCompensationTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// BatchGetCoreHRJobDataRespItemJobDataCompensationTypeDisplay ...
type BatchGetCoreHRJobDataRespItemJobDataCompensationTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// BatchGetCoreHRJobDataRespItemJobDataCostCenterRate ...
type BatchGetCoreHRJobDataRespItemJobDataCostCenterRate struct {
	CostCenterID string `json:"cost_center_id,omitempty"` // 成本中心 ID, 可以通过[搜索成本中心信息](https://open.feishu.cn/document/server-docs/corehr-v1/organization-management/cost_center/search)接口获取对应的成本中心信息
	Rate         int64  `json:"rate,omitempty"`           // 分摊比例
}

// BatchGetCoreHRJobDataRespItemJobDataDirectManager ...
type BatchGetCoreHRJobDataRespItemJobDataDirectManager struct {
	JobDataID    string `json:"job_data_id,omitempty"`   // 任职信息 ID
	EmploymentID string `json:"employment_id,omitempty"` // 雇佣 ID
}

// BatchGetCoreHRJobDataRespItemJobDataDottedLineManager ...
type BatchGetCoreHRJobDataRespItemJobDataDottedLineManager struct {
	JobDataID    string `json:"job_data_id,omitempty"`   // 任职信息 ID
	EmploymentID string `json:"employment_id,omitempty"` // 雇佣 ID
}

// BatchGetCoreHRJobDataRespItemJobDataProbationOutcome ...
type BatchGetCoreHRJobDataRespItemJobDataProbationOutcome struct {
	EnumName string                                                         `json:"enum_name,omitempty"` // 枚举值
	Display  []*BatchGetCoreHRJobDataRespItemJobDataProbationOutcomeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// BatchGetCoreHRJobDataRespItemJobDataProbationOutcomeDisplay ...
type BatchGetCoreHRJobDataRespItemJobDataProbationOutcomeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// BatchGetCoreHRJobDataRespItemJobDataSecondDirectManager ...
type BatchGetCoreHRJobDataRespItemJobDataSecondDirectManager struct {
	JobDataID    string `json:"job_data_id,omitempty"`   // 任职信息 ID
	EmploymentID string `json:"employment_id,omitempty"` // 雇佣 ID
}

// BatchGetCoreHRJobDataRespItemJobDataWorkShift ...
type BatchGetCoreHRJobDataRespItemJobDataWorkShift struct {
	EnumName string                                                  `json:"enum_name,omitempty"` // 枚举值
	Display  []*BatchGetCoreHRJobDataRespItemJobDataWorkShiftDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// BatchGetCoreHRJobDataRespItemJobDataWorkShiftDisplay ...
type BatchGetCoreHRJobDataRespItemJobDataWorkShiftDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// batchGetCoreHRJobDataResp ...
type batchGetCoreHRJobDataResp struct {
	Code  int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                     `json:"msg,omitempty"`  // 错误描述
	Data  *BatchGetCoreHRJobDataResp `json:"data,omitempty"`
	Error *ErrorDetail               `json:"error,omitempty"`
}
