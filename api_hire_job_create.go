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

// CreateHireJob 新建职位, 字段的是否必填, 以系统中的「职位字段管理」中的设置为准。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/job/combined_create
// new doc: https://open.feishu.cn/document/server-docs/hire-v1/recruitment-related-configuration/job/combined_create
func (r *HireService) CreateHireJob(ctx context.Context, request *CreateHireJobReq, options ...MethodOptionFunc) (*CreateHireJobResp, *Response, error) {
	if r.cli.mock.mockHireCreateHireJob != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Hire#CreateHireJob mock enable")
		return r.cli.mock.mockHireCreateHireJob(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "CreateHireJob",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/hire/v1/jobs/combined_create",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createHireJobResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHireCreateHireJob mock HireCreateHireJob method
func (r *Mock) MockHireCreateHireJob(f func(ctx context.Context, request *CreateHireJobReq, options ...MethodOptionFunc) (*CreateHireJobResp, *Response, error)) {
	r.mockHireCreateHireJob = f
}

// UnMockHireCreateHireJob un-mock HireCreateHireJob method
func (r *Mock) UnMockHireCreateHireJob() {
	r.mockHireCreateHireJob = nil
}

// CreateHireJobReq ...
type CreateHireJobReq struct {
	UserIDType                    *IDType                           `query:"user_id_type" json:"-"`                     // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	DepartmentIDType              *DepartmentIDType                 `query:"department_id_type" json:"-"`               // 此次调用中使用的部门 ID 的类型, 示例值: 此次调用中使用的部门 ID 的类型, 可选值有: open_department_id: 以 open_department_id 来标识部门, department_id: 以 department_id 来标识部门, 默认值: `open_department_id`
	Code                          *string                           `json:"code,omitempty"`                             // 职位编号, 可传入职位的「职位编号」、「职位 ID」或者「职位序号」, 将以传入的参数作为职位编号, 以便双方系统的数据映射, 示例值: "R18"
	Experience                    *int64                            `json:"experience,omitempty"`                       // 工作年限, 示例值: 1, 可选值有: 1: 不限, 2: 应届毕业生, 3: 1年以下, 4: 1-3年, 5: 3-5年, 6: 5-7年, 7: 7-10年, 8: 10年以上
	ExpiryTime                    *int64                            `json:"expiry_time,omitempty"`                      // 到期日期 请使用expiry_timestamp, 示例值: 1622484739
	CustomizedDataList            []*CreateHireJobReqCustomizedData `json:"customized_data_list,omitempty"`             // 自定义字段
	MinLevelID                    *string                           `json:"min_level_id,omitempty"`                     // 最低职级, 枚举通过接口「获取职级列表」获取, 示例值: "6960663240925956547"
	MinSalary                     *int64                            `json:"min_salary,omitempty"`                       // 最低薪资, 单位: k, 示例值: 1000
	Title                         string                            `json:"title,omitempty"`                            // 职位名称, 示例值: "后端研发"
	JobManagers                   *CreateHireJobReqJobManagers      `json:"job_managers,omitempty"`                     // 职位的招聘团队
	JobProcessID                  string                            `json:"job_process_id,omitempty"`                   // 招聘流程, 枚举通过接口「获取招聘流程信息」获取, 示例值: "6960663240925956554"
	ProcessType                   int64                             `json:"process_type,omitempty"`                     // 职位流程类型, 示例值: 1, 可选值有: 1: 社招, 2: 校招
	SubjectID                     *string                           `json:"subject_id,omitempty"`                       // 项目, 枚举通过「获取项目列表」接口, 示例值: "6960663240925956555"
	JobFunctionID                 *string                           `json:"job_function_id,omitempty"`                  // 职能分类, 通过「获取职能分类」获取, 示例值: "6960663240925956555"
	DepartmentID                  string                            `json:"department_id,omitempty"`                    // 部门 ID, 须传入open_department_id, 格式为"od-xxxx"。可通过「获取部门信息列表」获取, 示例值: "od-b2fafdce6fc5800b574ba5b0e2798b36"
	HeadCount                     *int64                            `json:"head_count,omitempty"`                       // 招聘数量, 示例值: 100
	IsNeverExpired                bool                              `json:"is_never_expired,omitempty"`                 // 是否长期有效, 示例值: false
	MaxSalary                     *int64                            `json:"max_salary,omitempty"`                       // 最高薪资, 单位: k, 示例值: 2000
	Requirement                   *string                           `json:"requirement,omitempty"`                      // 职位要求, 示例值: "熟悉后端研发"
	Description                   *string                           `json:"description,omitempty"`                      // 职位描述, 示例值: "后端研发岗位描述"
	HighlightList                 []string                          `json:"highlight_list,omitempty"`                   // 职位亮点, 示例值: ["免费三餐"]
	JobTypeID                     string                            `json:"job_type_id,omitempty"`                      // 职位类别, 示例值: "6960663240925956551"
	MaxLevelID                    *string                           `json:"max_level_id,omitempty"`                     // 最高职级, 枚举通过接口「获取职级列表」获取, 示例值: "6960663240925956548"
	RecruitmentTypeID             string                            `json:"recruitment_type_id,omitempty"`              // 雇佣类型, 社招: 101-全职  102-外包 103-劳务 105-顾问 301-实习, 校招: 201-正式 202-实习, 示例值: "102"
	RequiredDegree                *int64                            `json:"required_degree,omitempty"`                  // 学历要求, 示例值: 20, 可选值有: 1: 小学及以上, 2: 初中及以上, 3: 专职及以上, 4: 高中及以上, 5: 大专及以上, 6: 本科及以上, 7: 硕士及以上, 8: 博士及以上, 20: 不限
	JobCategoryID                 *string                           `json:"job_category_id,omitempty"`                  // 序列ID, 示例值: "6960663240925956550"
	AddressIDList                 []string                          `json:"address_id_list,omitempty"`                  // 工作地点列表, 枚举通过接口「获取地址列表」获取, 选择地点用途为「职位地址」, 与「address_id」必填其一, 同时使用则以当前字段为准, 推荐使用当前字段, 示例值: ["7243965681799839748"]
	JobAttribute                  *int64                            `json:"job_attribute,omitempty"`                    // 职位属性, 1是实体职位, 2是虚拟职位, 示例值: 1, 可选值有: 1: 实体职位, 2: 虚拟职位
	ExpiryTimestamp               *string                           `json:"expiry_timestamp,omitempty"`                 // 到期日期的毫秒时间戳, 如果「是否长期有效」字段选择true, 则不会实际使用该字段的值, 职位为长期有效, 示例值: "1622484739955"
	InterviewRegistrationSchemaID *string                           `json:"interview_registration_schema_id,omitempty"` // 面试登记表ID, 当在飞书招聘「设置 - 面试登记表使用设置 - 面试登记表使用方式」中选择「全部职位应用同一登记表」时, 「职位设置」下的面试登记表不生效；选择「HR 按职位选择登记表」时, 该字段为必填, 示例值: "6930815272790114324"
}

// CreateHireJobReqCustomizedData ...
type CreateHireJobReqCustomizedData struct {
	ObjectID *string `json:"object_id,omitempty"` // 结构 ID, 示例值: "6960663240925956549"
	Value    *string `json:"value,omitempty"`     // 结构值, 示例值: "测试"
}

// CreateHireJobReqJobManagers ...
type CreateHireJobReqJobManagers struct {
	ID                  *string  `json:"id,omitempty"`                     // 职位 ID, 示例值: "1618209327096"
	RecruiterID         string   `json:"recruiter_id,omitempty"`           // 招聘负责人 ID, 仅一位, 可通过用户相关接口获取用户 ID, 示例值: "ou_efk39117c300506837def50545420c6a"
	HiringManagerIDList []string `json:"hiring_manager_id_list,omitempty"` // 用人经理 ID 列表, 示例值: ["on_94a1ee5551019f18cd73d9f111898cf2"]
	AssistantIDList     []string `json:"assistant_id_list,omitempty"`      // 协助人 ID 列表, 示例值: ["on_94a1ee5551019f18cd73d9f111898cf2"]
}

// CreateHireJobResp ...
type CreateHireJobResp struct {
	DefaultJobPost                  *CreateHireJobRespDefaultJobPost                  `json:"default_job_post,omitempty"`                   // 职位广告
	Job                             *CreateHireJobRespJob                             `json:"job,omitempty"`                                // 职位
	JobManager                      *CreateHireJobRespJobManager                      `json:"job_manager,omitempty"`                        // 职位负责人
	InterviewRegistrationSchemaInfo *CreateHireJobRespInterviewRegistrationSchemaInfo `json:"interview_registration_schema_info,omitempty"` // 面试登记表信息
}

// CreateHireJobRespDefaultJobPost ...
type CreateHireJobRespDefaultJobPost struct {
	ID string `json:"id,omitempty"` // 默认职位广告的 ID, 用以发布至招聘渠道的内容
}

// CreateHireJobRespInterviewRegistrationSchemaInfo ...
type CreateHireJobRespInterviewRegistrationSchemaInfo struct {
	SchemaID string `json:"schema_id,omitempty"` // 面试登记表ID
	Name     string `json:"name,omitempty"`      // 面试登记表名称
}

// CreateHireJobRespJob ...
type CreateHireJobRespJob struct {
	ID                 string                                `json:"id,omitempty"`                   // 职位 ID
	Title              string                                `json:"title,omitempty"`                // 职位名称
	Description        string                                `json:"description,omitempty"`          // 职位描述
	Code               string                                `json:"code,omitempty"`                 // 职位编号
	Requirement        string                                `json:"requirement,omitempty"`          // 职位要求
	RecruitmentType    *CreateHireJobRespJobRecruitmentType  `json:"recruitment_type,omitempty"`     // 雇佣类型
	Department         *CreateHireJobRespJobDepartment       `json:"department,omitempty"`           // 部门
	City               *CreateHireJobRespJobCity             `json:"city,omitempty"`                 // 工作地点
	MinJobLevel        *CreateHireJobRespJobMinJobLevel      `json:"min_job_level,omitempty"`        // 最低职级
	MaxJobLevel        *CreateHireJobRespJobMaxJobLevel      `json:"max_job_level,omitempty"`        // 最高职级
	HighlightList      []*CreateHireJobRespJobHighlight      `json:"highlight_list,omitempty"`       // 职位亮点
	JobCategory        *CreateHireJobRespJobJobCategory      `json:"job_category,omitempty"`         // 职位序列
	JobType            *CreateHireJobRespJobJobType          `json:"job_type,omitempty"`             // 职位类别
	ActiveStatus       int64                                 `json:"active_status,omitempty"`        // 启用状态, 可选值有: 1: 启用, 2: 未启用
	CreateUserID       string                                `json:"create_user_id,omitempty"`       // 创建人ID, 若为空则为系统或其他对接系统创建
	CreateTime         int64                                 `json:"create_time,omitempty"`          // 创建时间 请使用create_timestamp
	UpdateTime         int64                                 `json:"update_time,omitempty"`          // 更新时间 请使用update_timestamp
	ProcessType        int64                                 `json:"process_type,omitempty"`         // 招聘流程类型, 可选值有: 1: 社招流程, 2: 校招流程
	ProcessID          string                                `json:"process_id,omitempty"`           // 招聘流程 ID
	ProcessName        string                                `json:"process_name,omitempty"`         // 招聘流程中文名称
	ProcessEnName      string                                `json:"process_en_name,omitempty"`      // 招聘流程英文名称
	CustomizedDataList []*CreateHireJobRespJobCustomizedData `json:"customized_data_list,omitempty"` // 自定义字段列表
	JobFunction        *CreateHireJobRespJobJobFunction      `json:"job_function,omitempty"`         // 职能分类
	Subject            *CreateHireJobRespJobSubject          `json:"subject,omitempty"`              // 职位项目
	HeadCount          int64                                 `json:"head_count,omitempty"`           // 招聘数量
	Experience         int64                                 `json:"experience,omitempty"`           // 工作年限, 可选值有: 1: 不限, 2: 应届毕业生, 3: 1年以下, 4: 1-3年, 5: 3-5年, 6: 5-7年, 7: 7-10年, 8: 10年以上
	ExpiryTime         int64                                 `json:"expiry_time,omitempty"`          // 到期日期 请使用expiry_timestamp
	MinSalary          int64                                 `json:"min_salary,omitempty"`           // 最低薪资, 单位:k
	MaxSalary          int64                                 `json:"max_salary,omitempty"`           // 最高薪资, 单位:k
	RequiredDegree     int64                                 `json:"required_degree,omitempty"`      // 学历要求, 可选值有: 1: JuniorMiddleSchoolEducation, 2: 初中及以上, 3: 专职及以上, 4: 高中及以上, 5: 大专及以上, 6: 本科及以上, 7: 硕士及以上, 8: 博士及以上, 20: 不限
	CityList           []*CreateHireJobRespJobCity           `json:"city_list,omitempty"`            // 工作地点列表
	JobAttribute       int64                                 `json:"job_attribute,omitempty"`        // 职位属性, 1是实体职位, 2是虚拟职位, 可选值有: 1: 实体职位, 2: 虚拟职位
	CreateTimestamp    string                                `json:"create_timestamp,omitempty"`     // 创建时间戳
	UpdateTimestamp    string                                `json:"update_timestamp,omitempty"`     // 更新时间戳
	ExpiryTimestamp    string                                `json:"expiry_timestamp,omitempty"`     // 到期时间戳
}

// CreateHireJobRespJobCity ...
type CreateHireJobRespJobCity struct {
	CityCode string `json:"city_code,omitempty"` // 工作地点城市代码
	ZhName   string `json:"zh_name,omitempty"`   // 工作地点中文名称
	EnName   string `json:"en_name,omitempty"`   // 工作地点英文名称
}

// CreateHireJobRespJobCityName ...
type CreateHireJobRespJobCityName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// CreateHireJobRespJobCustomizedData ...
type CreateHireJobRespJobCustomizedData struct {
	ObjectID   string                                   `json:"object_id,omitempty"`   // 自定义字段 ID
	Name       *CreateHireJobRespJobCustomizedDataName  `json:"name,omitempty"`        // 字段名称
	ObjectType int64                                    `json:"object_type,omitempty"` // 字段类型, 可选值有: 1: 单行文本, 2: 多行文本, 3: 单选, 4: 多选, 5: 日期, 6: 月份选择, 7: 年份选择, 8: 时间段, 9: 数字, 10: 默认字段, 11: 模块
	Value      *CreateHireJobRespJobCustomizedDataValue `json:"value,omitempty"`       // 自定义字段值
}

// CreateHireJobRespJobCustomizedDataName ...
type CreateHireJobRespJobCustomizedDataName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// CreateHireJobRespJobCustomizedDataValue ...
type CreateHireJobRespJobCustomizedDataValue struct {
	Content    string                                            `json:"content,omitempty"`     // 当字段类型为单行文本、多行文本、模块、默认字段时, 从此字段取值
	Option     *CreateHireJobRespJobCustomizedDataValueOption    `json:"option,omitempty"`      // 当字段类型为单选时, 从此字段取值
	OptionList []*CreateHireJobRespJobCustomizedDataValueOption  `json:"option_list,omitempty"` // 当字段类型为多选时, 从此字段取值
	TimeRange  *CreateHireJobRespJobCustomizedDataValueTimeRange `json:"time_range,omitempty"`  // 当字段类型为时间段时, 从此字段取值
	Time       string                                            `json:"time,omitempty"`        // 当字段类型为日期选择、月份选择、年份选择时, 从此字段取值, 该字段是毫秒级时间戳
	Number     string                                            `json:"number,omitempty"`      // 当字段类型为数字时, 从此字段取值
}

// CreateHireJobRespJobCustomizedDataValueOption ...
type CreateHireJobRespJobCustomizedDataValueOption struct {
	Key  string                                             `json:"key,omitempty"`  // 选项 ID
	Name *CreateHireJobRespJobCustomizedDataValueOptionName `json:"name,omitempty"` // 选项名称
}

// CreateHireJobRespJobCustomizedDataValueOptionName ...
type CreateHireJobRespJobCustomizedDataValueOptionName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// CreateHireJobRespJobCustomizedDataValueTimeRange ...
type CreateHireJobRespJobCustomizedDataValueTimeRange struct {
	StartTime string `json:"start_time,omitempty"` // 开始时间
	EndTime   string `json:"end_time,omitempty"`   // 结束时间
}

// CreateHireJobRespJobDepartment ...
type CreateHireJobRespJobDepartment struct {
	ID     string `json:"id,omitempty"`      // 部门 ID
	ZhName string `json:"zh_name,omitempty"` // 部门中文名称
	EnName string `json:"en_name,omitempty"` // 部门英文名称
}

// CreateHireJobRespJobHighlight ...
type CreateHireJobRespJobHighlight struct {
	ID     string `json:"id,omitempty"`      // 职位亮点 ID
	ZhName string `json:"zh_name,omitempty"` // 职位亮点中文名称
	EnName string `json:"en_name,omitempty"` // 职位亮点英文名称
}

// CreateHireJobRespJobJobCategory ...
type CreateHireJobRespJobJobCategory struct {
	ID           string `json:"id,omitempty"`            // 职位序列 ID
	ZhName       string `json:"zh_name,omitempty"`       // 职位序列中文名称
	EnName       string `json:"en_name,omitempty"`       // 职位序列英文名称
	ActiveStatus int64  `json:"active_status,omitempty"` // 职位序列启用状态, 可选值有: 1: 启用, 2: 未启用
}

// CreateHireJobRespJobJobFunction ...
type CreateHireJobRespJobJobFunction struct {
	ID   string                               `json:"id,omitempty"`   // ID
	Name *CreateHireJobRespJobJobFunctionName `json:"name,omitempty"` // 名称
}

// CreateHireJobRespJobJobFunctionName ...
type CreateHireJobRespJobJobFunctionName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// CreateHireJobRespJobJobType ...
type CreateHireJobRespJobJobType struct {
	ID     string `json:"id,omitempty"`      // 职位类别 ID
	ZhName string `json:"zh_name,omitempty"` // 职位类别中文名称
	EnName string `json:"en_name,omitempty"` // 职位类别英文名称
}

// CreateHireJobRespJobManager ...
type CreateHireJobRespJobManager struct {
	ID                  string   `json:"id,omitempty"`                     // 职位 ID
	RecruiterID         string   `json:"recruiter_id,omitempty"`           // 招聘负责人 ID, 仅一位, 可通过用户相关接口获取用户 ID
	HiringManagerIDList []string `json:"hiring_manager_id_list,omitempty"` // 用人经理 ID 列表
	AssistantIDList     []string `json:"assistant_id_list,omitempty"`      // 协助人 ID 列表
}

// CreateHireJobRespJobMaxJobLevel ...
type CreateHireJobRespJobMaxJobLevel struct {
	ID           string `json:"id,omitempty"`            // 职级 ID
	ZhName       string `json:"zh_name,omitempty"`       // 职级中文名称
	EnName       string `json:"en_name,omitempty"`       // 职级英文名称
	ActiveStatus int64  `json:"active_status,omitempty"` // 职级启用状态, 可选值有: 1: 启用, 2: 未启用
}

// CreateHireJobRespJobMinJobLevel ...
type CreateHireJobRespJobMinJobLevel struct {
	ID           string `json:"id,omitempty"`            // 职级 ID
	ZhName       string `json:"zh_name,omitempty"`       // 职级中文名称
	EnName       string `json:"en_name,omitempty"`       // 职级英文名称
	ActiveStatus int64  `json:"active_status,omitempty"` // 职级启用状态, 可选值有: 1: 启用, 2: 未启用
}

// CreateHireJobRespJobRecruitmentType ...
type CreateHireJobRespJobRecruitmentType struct {
	ID           string `json:"id,omitempty"`            // 雇佣类型 ID
	ZhName       string `json:"zh_name,omitempty"`       // 雇佣类型中文名称
	EnName       string `json:"en_name,omitempty"`       // 雇佣类型英文名称
	ActiveStatus int64  `json:"active_status,omitempty"` // 雇佣类型启用状态, 可选值有: 1: 启用, 2: 未启用
}

// CreateHireJobRespJobSubject ...
type CreateHireJobRespJobSubject struct {
	ID   string                           `json:"id,omitempty"`   // ID
	Name *CreateHireJobRespJobSubjectName `json:"name,omitempty"` // 名称
}

// CreateHireJobRespJobSubjectName ...
type CreateHireJobRespJobSubjectName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// createHireJobResp ...
type createHireJobResp struct {
	Code int64              `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string             `json:"msg,omitempty"`  // 错误描述
	Data *CreateHireJobResp `json:"data,omitempty"`
}
