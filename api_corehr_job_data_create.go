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

// CreateCoreHRJobData 在系统中第一次创建员工任职数据, 通常在员工入职或者做数据批量导入的时候使用, [任职原因]只支持填写“onboarding”。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/job_data/create
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/employee/job_data/create
func (r *CoreHRService) CreateCoreHRJobData(ctx context.Context, request *CreateCoreHRJobDataReq, options ...MethodOptionFunc) (*CreateCoreHRJobDataResp, *Response, error) {
	if r.cli.mock.mockCoreHRCreateCoreHRJobData != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#CreateCoreHRJobData mock enable")
		return r.cli.mock.mockCoreHRCreateCoreHRJobData(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "CreateCoreHRJobData",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/job_datas",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createCoreHRJobDataResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRCreateCoreHRJobData mock CoreHRCreateCoreHRJobData method
func (r *Mock) MockCoreHRCreateCoreHRJobData(f func(ctx context.Context, request *CreateCoreHRJobDataReq, options ...MethodOptionFunc) (*CreateCoreHRJobDataResp, *Response, error)) {
	r.mockCoreHRCreateCoreHRJobData = f
}

// UnMockCoreHRCreateCoreHRJobData un-mock CoreHRCreateCoreHRJobData method
func (r *Mock) UnMockCoreHRCreateCoreHRJobData() {
	r.mockCoreHRCreateCoreHRJobData = nil
}

// CreateCoreHRJobDataReq ...
type CreateCoreHRJobDataReq struct {
	ClientToken              *string                                      `query:"client_token" json:"-"`                // 操作的唯一标识, 用于幂等的进行更新操作, 格式为标准的 UUIDV4。此值为空表示将发起一次新的请求, 此值非空表示幂等的进行更新操作, 示例值: "fe599b60-450f-46ff-b2ef-9f6675625b97"
	UserIDType               *IDType                                      `query:"user_id_type" json:"-"`                // 用户 ID 类型, 示例值: people_corehr_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), people_corehr_id: 以飞书人事的 ID 来识别用户, 默认值: `people_corehr_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	DepartmentIDType         *DepartmentIDType                            `query:"department_id_type" json:"-"`          // 此次调用中使用的部门 ID 类型, 示例值: open_department_id, 可选值有: open_department_id: [飞书]用来在具体某个应用中标识一个部门, 同一个 department_id 在不同应用中的 open_department_id 相同。, department_id: [飞书]用来标识租户内一个唯一的部门。, people_corehr_department_id: [飞书人事]用来标识「飞书人事」中的部门。, 默认值: `people_corehr_department_id`
	JobLevelID               *string                                      `json:"job_level_id,omitempty"`                // 职务级别 ID, 枚举值及详细信息可通过[查询单个职级](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/job_level/get)接口查询获得, 示例值: "6890452208593372679"
	JobGradeID               *string                                      `json:"job_grade_id,omitempty"`                // 职等 ID, 枚举值及详细信息可通过[查询职等](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/job_grade/query)接口查询获得, 示例值: "6890452208593372679"
	EmployeeTypeID           string                                       `json:"employee_type_id,omitempty"`            // 人员类型 ID, 枚举值及详细信息可通过[查询单个人员类型](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/employee_type/get)接口查询获得, 示例值: "6890452208593372679"
	WorkingHoursTypeID       *string                                      `json:"working_hours_type_id,omitempty"`       // 工时制度 ID, 枚举值及详细信息可通过[查询单个工时制度](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/working_hours_type/get)接口查询获得, 示例值: "6890452208593372679"
	WorkLocationID           *string                                      `json:"work_location_id,omitempty"`            // 工作地点 ID, 枚举值及详细信息可通过[查询单个地点](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/location/get)接口查询获得, 示例值: "6890452208593372679"
	DepartmentID             *string                                      `json:"department_id,omitempty"`               // 部门 ID, 枚举值及详细信息可通过[查询单个部门](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/department/get)接口查询获得, 示例值: "6890452208593372679"
	JobID                    *string                                      `json:"job_id,omitempty"`                      // 职务 ID, 枚举值及详细信息可通过[查询单个职务](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/job/get)接口查询获得, 示例值: "6890452208593372679"
	ProbationStartDate       *string                                      `json:"probation_start_date,omitempty"`        // 试用期开始日期, 示例值: "2018-03-16"
	ProbationEndDate         *string                                      `json:"probation_end_date,omitempty"`          // 试用期结束日期（实际结束日期）, 示例值: "2019-05-24"
	PrimaryJobData           bool                                         `json:"primary_job_data,omitempty"`            // 是否为主任职, 示例值: true
	EmploymentID             string                                       `json:"employment_id,omitempty"`               // 雇佣 ID, 详细信息可以通过[搜索员工信息](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/employee/search)接口查询获得, 示例值: "6893014062142064135"
	EffectiveTime            string                                       `json:"effective_time,omitempty"`              // 生效时间, 示例值: "2020-05-01 00:00:00"
	ExpirationTime           *string                                      `json:"expiration_time,omitempty"`             // 失效时间, 示例值: "2020-05-02 00:00:00"
	JobFamilyID              *string                                      `json:"job_family_id,omitempty"`               // 职务序列 ID, 枚举值及详细信息可通过[查询单个序列](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/job_family/get)接口查询获得, 示例值: "1245678"
	AssignmentStartReason    *CreateCoreHRJobDataReqAssignmentStartReason `json:"assignment_start_reason,omitempty"`     // 任职原因, 枚举值可通过文档[飞书人事枚举常量](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)任职原因（assignment_start_reason）枚举定义部分获得, 这里只支持填写"onboarding", 示例值: onboarding
	ProbationExpectedEndDate *string                                      `json:"probation_expected_end_date,omitempty"` // 预计试用期结束日期, 示例值: "2006-01-02"
	DirectManagerID          *string                                      `json:"direct_manager_id,omitempty"`           // 直属上级的任职记录 ID, 详细信息可通过[批量查询员工任职信息](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/employees-job_data/batch_get)接口查询获得, 示例值: "6890452208593372679"
	DottedLineManagerIDList  []string                                     `json:"dotted_line_manager_id_list,omitempty"` // 虚线上级的任职记录 ID, 详细信息可通过[批量查询员工任职信息](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/employees-job_data/batch_get)接口查询获得, 示例值: ["6890452208593372681"]
	SecondDirectManagerID    *string                                      `json:"second_direct_manager_id,omitempty"`    // 第二直属上级的任职记录 ID, 详细信息可通过[批量查询员工任职信息](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/employees-job_data/batch_get)接口查询获得, 示例值: "6890452208593372679"
	CostCenterRate           []*CreateCoreHRJobDataReqCostCenterRate      `json:"cost_center_rate,omitempty"`            // 成本中心分摊信息
	WorkShift                *CreateCoreHRJobDataReqWorkShift             `json:"work_shift,omitempty"`                  // 排班类型
	CompensationType         *CreateCoreHRJobDataReqCompensationType      `json:"compensation_type,omitempty"`           // 薪资类型
	ServiceCompany           *string                                      `json:"service_company,omitempty"`             // 任职公司 ID, 详细信息可通过[查询单个公司](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/company/get)接口查询获得, 示例值: "6890452208593372680"
}

// CreateCoreHRJobDataReqAssignmentStartReason ...
type CreateCoreHRJobDataReqAssignmentStartReason struct {
	EnumName string `json:"enum_name,omitempty"` // 枚举值, 示例值: "onboarding"
}

// CreateCoreHRJobDataReqCompensationType ...
type CreateCoreHRJobDataReqCompensationType struct {
	EnumName string `json:"enum_name,omitempty"` // 枚举值。需自定义, 示例值: "phone_type"
}

// CreateCoreHRJobDataReqCostCenterRate ...
type CreateCoreHRJobDataReqCostCenterRate struct {
	CostCenterID *string `json:"cost_center_id,omitempty"` // 支持的成本中心 ID, 详细信息可通过[搜索成本中心信息](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/cost_center/search)接口查询获得, 示例值: "6950635856373745165"
	Rate         *int64  `json:"rate,omitempty"`           // 分摊比例, 示例值: 100
}

// CreateCoreHRJobDataReqWorkShift ...
type CreateCoreHRJobDataReqWorkShift struct {
	EnumName string `json:"enum_name,omitempty"` // 枚举值。需自定义, 示例值: "phone_type"
}

// CreateCoreHRJobDataResp ...
type CreateCoreHRJobDataResp struct {
	JobData *CreateCoreHRJobDataRespJobData `json:"job_data,omitempty"` // 创建成功返回job data信息
}

// CreateCoreHRJobDataRespJobData ...
type CreateCoreHRJobDataRespJobData struct {
	ID                       string                                               `json:"id,omitempty"`                          // 任职信息 ID
	VersionID                string                                               `json:"version_id,omitempty"`                  // 任职记录版本 ID
	JobLevelID               string                                               `json:"job_level_id,omitempty"`                // 职务级别 ID, 枚举值及详细信息可通过[查询单个职级](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/job_level/get)接口查询获得
	JobGradeID               string                                               `json:"job_grade_id,omitempty"`                // 职等 ID, 枚举值及详细信息可通过[查询职等](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/job_grade/query)接口查询获得
	EmployeeTypeID           string                                               `json:"employee_type_id,omitempty"`            // 人员类型 ID, 枚举值及详细信息可通过[查询单个人员类型](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/employee_type/get)接口查询获得
	WorkingHoursTypeID       string                                               `json:"working_hours_type_id,omitempty"`       // 工时制度 ID, 枚举值及详细信息可通过[查询单个工时制度](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/working_hours_type/get)接口查询获得
	WorkLocationID           string                                               `json:"work_location_id,omitempty"`            // 工作地点 ID, 枚举值及详细信息可通过[查询单个地点](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/location/get)接口查询获得
	DepartmentID             string                                               `json:"department_id,omitempty"`               // 部门 ID, 枚举值及详细信息可通过[查询单个部门](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/department/get)接口查询获得
	JobID                    string                                               `json:"job_id,omitempty"`                      // 职务 ID, 枚举值及详细信息可通过[查询单个职务](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/job/get)接口查询获得
	ProbationStartDate       string                                               `json:"probation_start_date,omitempty"`        // 试用期开始日期
	ProbationEndDate         string                                               `json:"probation_end_date,omitempty"`          // 试用期结束日期（实际结束日期）
	PrimaryJobData           bool                                                 `json:"primary_job_data,omitempty"`            // 是否为主任职
	EmploymentID             string                                               `json:"employment_id,omitempty"`               // 雇佣 ID, 详细信息可以通过[搜索员工信息](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/employee/search)接口查询获得
	EffectiveTime            string                                               `json:"effective_time,omitempty"`              // 生效时间
	ExpirationTime           string                                               `json:"expiration_time,omitempty"`             // 失效时间
	JobFamilyID              string                                               `json:"job_family_id,omitempty"`               // 职务序列 ID, 枚举值及详细信息可通过[查询单个序列](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/job_family/get)接口查询获得
	AssignmentStartReason    *CreateCoreHRJobDataRespJobDataAssignmentStartReason `json:"assignment_start_reason,omitempty"`     // 任职原因, 枚举值可通过文档[飞书人事枚举常量](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)任职原因（assignment_start_reason）枚举定义部分获得
	ProbationExpectedEndDate string                                               `json:"probation_expected_end_date,omitempty"` // 预计试用期结束日期
	WeeklyWorkingHours       int64                                                `json:"weekly_working_hours,omitempty"`        // 周工作时长
	DirectManagerID          string                                               `json:"direct_manager_id,omitempty"`           // 直属上级的任职记录 ID, 详细信息可通过[批量查询员工任职信息](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/employees-job_data/batch_get)接口查询获得
	DottedLineManagerIDList  []string                                             `json:"dotted_line_manager_id_list,omitempty"` // 虚线上级的任职记录 ID, 详细信息可通过[批量查询员工任职信息](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/employees-job_data/batch_get)接口查询获得
	SecondDirectManagerID    string                                               `json:"second_direct_manager_id,omitempty"`    // 第二直属上级的任职记录 ID, 详细信息可通过[批量查询员工任职信息](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/employees-job_data/batch_get)接口查询获得
	CostCenterRate           []*CreateCoreHRJobDataRespJobDataCostCenterRate      `json:"cost_center_rate,omitempty"`            // 成本中心分摊信息
	WeeklyWorkingHoursV2     float64                                              `json:"weekly_working_hours_v2,omitempty"`     // 周工作时长
	WorkShift                *CreateCoreHRJobDataRespJobDataWorkShift             `json:"work_shift,omitempty"`                  // 排班类型, 字段权限要求: 获取排班信息
	CompensationType         *CreateCoreHRJobDataRespJobDataCompensationType      `json:"compensation_type,omitempty"`           // 薪资类型, 字段权限要求: 获取薪资类型
	ServiceCompany           string                                               `json:"service_company,omitempty"`             // 任职公司, 字段权限要求: 获取任职公司
}

// CreateCoreHRJobDataRespJobDataAssignmentStartReason ...
type CreateCoreHRJobDataRespJobDataAssignmentStartReason struct {
	EnumName string                                                        `json:"enum_name,omitempty"` // 枚举值
	Display  []*CreateCoreHRJobDataRespJobDataAssignmentStartReasonDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// CreateCoreHRJobDataRespJobDataAssignmentStartReasonDisplay ...
type CreateCoreHRJobDataRespJobDataAssignmentStartReasonDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// CreateCoreHRJobDataRespJobDataCompensationType ...
type CreateCoreHRJobDataRespJobDataCompensationType struct {
	EnumName string                                                   `json:"enum_name,omitempty"` // 枚举值
	Display  []*CreateCoreHRJobDataRespJobDataCompensationTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// CreateCoreHRJobDataRespJobDataCompensationTypeDisplay ...
type CreateCoreHRJobDataRespJobDataCompensationTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// CreateCoreHRJobDataRespJobDataCostCenterRate ...
type CreateCoreHRJobDataRespJobDataCostCenterRate struct {
	CostCenterID string `json:"cost_center_id,omitempty"` // 支持的成本中心 ID, 详细信息可通过[搜索成本中心信息](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/cost_center/search)接口查询获得
	Rate         int64  `json:"rate,omitempty"`           // 分摊比例
}

// CreateCoreHRJobDataRespJobDataWorkShift ...
type CreateCoreHRJobDataRespJobDataWorkShift struct {
	EnumName string                                            `json:"enum_name,omitempty"` // 枚举值
	Display  []*CreateCoreHRJobDataRespJobDataWorkShiftDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// CreateCoreHRJobDataRespJobDataWorkShiftDisplay ...
type CreateCoreHRJobDataRespJobDataWorkShiftDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// createCoreHRJobDataResp ...
type createCoreHRJobDataResp struct {
	Code  int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                   `json:"msg,omitempty"`  // 错误描述
	Data  *CreateCoreHRJobDataResp `json:"data,omitempty"`
	Error *ErrorDetail             `json:"error,omitempty"`
}
