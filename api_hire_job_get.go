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

// GetHireJob 根据职位 ID 获取职位信息
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/job/get
func (r *HireService) GetHireJob(ctx context.Context, request *GetHireJobReq, options ...MethodOptionFunc) (*GetHireJobResp, *Response, error) {
	if r.cli.mock.mockHireGetHireJob != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Hire#GetHireJob mock enable")
		return r.cli.mock.mockHireGetHireJob(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "GetHireJob",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/hire/v1/jobs/:job_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getHireJobResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHireGetHireJob mock HireGetHireJob method
func (r *Mock) MockHireGetHireJob(f func(ctx context.Context, request *GetHireJobReq, options ...MethodOptionFunc) (*GetHireJobResp, *Response, error)) {
	r.mockHireGetHireJob = f
}

// UnMockHireGetHireJob un-mock HireGetHireJob method
func (r *Mock) UnMockHireGetHireJob() {
	r.mockHireGetHireJob = nil
}

// GetHireJobReq ...
type GetHireJobReq struct {
	JobID      int64   `path:"job_id" json:"-"`        // 职位 ID, 请求Path中, 示例值: 6001
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: `open_id`: 用户的 open id, `union_id`: 用户的 union id, `user_id`: 用户的 user id, `people_admin_id`: 以people_admin_id来识别用户, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetHireJobResp ...
type GetHireJobResp struct {
	Job *GetHireJobRespJob `json:"job,omitempty"` // 职位数据
}

// GetHireJobRespJob ...
type GetHireJobRespJob struct {
	ID                 string                             `json:"id,omitempty"`                   // 职位 ID
	Title              string                             `json:"title,omitempty"`                // 职位名称
	Description        string                             `json:"description,omitempty"`          // 职位描述
	Code               string                             `json:"code,omitempty"`                 // 职位编号
	Requirement        string                             `json:"requirement,omitempty"`          // 职位要求
	RecruitmentType    *GetHireJobRespJobRecruitmentType  `json:"recruitment_type,omitempty"`     // 雇佣类型
	Department         *GetHireJobRespJobDepartment       `json:"department,omitempty"`           // 部门
	City               *GetHireJobRespJobCity             `json:"city,omitempty"`                 // 工作地点
	MinJobLevel        *GetHireJobRespJobMinJobLevel      `json:"min_job_level,omitempty"`        // 最低职级
	MaxJobLevel        *GetHireJobRespJobMaxJobLevel      `json:"max_job_level,omitempty"`        // 最高职级
	HighlightList      []*GetHireJobRespJobHighlight      `json:"highlight_list,omitempty"`       // 职位亮点
	JobCategory        *GetHireJobRespJobJobCategory      `json:"job_category,omitempty"`         // 职位序列
	JobType            *GetHireJobRespJobJobType          `json:"job_type,omitempty"`             // 职位类别
	ActiveStatus       int64                              `json:"active_status,omitempty"`        // 启用状态, 可选值有: `1`: 启用, `2`: 未启用
	CreateUserID       string                             `json:"create_user_id,omitempty"`       // 创建人ID, 若为空则为系统或其他对接系统创建
	CreateTime         int64                              `json:"create_time,omitempty"`          // 创建时间
	UpdateTime         int64                              `json:"update_time,omitempty"`          // 更新时间
	ProcessType        int64                              `json:"process_type,omitempty"`         // 招聘流程类型, 可选值有: `1`: 社招流程, `2`: 校招流程
	ProcessID          string                             `json:"process_id,omitempty"`           // 招聘流程 ID
	ProcessName        string                             `json:"process_name,omitempty"`         // 招聘流程中文名称
	ProcessEnName      string                             `json:"process_en_name,omitempty"`      // 招聘流程英文名称
	CustomizedDataList []*GetHireJobRespJobCustomizedData `json:"customized_data_list,omitempty"` // 自定义字段列表
	JobFunction        *GetHireJobRespJobJobFunction      `json:"job_function,omitempty"`         // 职能分类
	Subject            *GetHireJobRespJobSubject          `json:"subject,omitempty"`              // 职位项目
	HeadCount          int64                              `json:"head_count,omitempty"`           // 招聘数量
	Experience         int64                              `json:"experience,omitempty"`           // 工作年限, 可选值有: `1`: 不限, `2`: 应届毕业生, `3`: 1年以下, `4`: 1-3年, `5`: 3-5年, `6`: 5-7年, `7`: 7-10年, `8`: 10年以上
	ExpiryTime         int64                              `json:"expiry_time,omitempty"`          // 到期日期
	MinSalary          int64                              `json:"min_salary,omitempty"`           // 最低薪资, 单位:k
	MaxSalary          int64                              `json:"max_salary,omitempty"`           // 最高薪资, 单位:k
	RequiredDegree     int64                              `json:"required_degree,omitempty"`      // 学历要求, 可选值有: `1`: 小学及以上, `2`: 初中及以上, `3`: 专职及以上, `4`: 高中及以上, `5`: 大专及以上, `6`: 本科及以上, `7`: 硕士及以上, `8`: 博士及以上, `20`: 不限
	CityList           []*GetHireJobRespJobCity           `json:"city_list,omitempty"`            // 工作地点列表
}

// GetHireJobRespJobCity ...
type GetHireJobRespJobCity struct {
	Code string                     `json:"code,omitempty"` // 编码
	Name *GetHireJobRespJobCityName `json:"name,omitempty"` // 名称
}

// GetHireJobRespJobCityName ...
type GetHireJobRespJobCityName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireJobRespJobCustomizedData ...
type GetHireJobRespJobCustomizedData struct {
	ObjectID   string                                `json:"object_id,omitempty"`   // 自定义字段 ID
	Name       *GetHireJobRespJobCustomizedDataName  `json:"name,omitempty"`        // 字段名称
	ObjectType int64                                 `json:"object_type,omitempty"` // 字段类型, 可选值有: `1`: 单行文本, `2`: 多行文本, `3`: 单选, `4`: 多选, `5`: 日期, `6`: 月份选择, `7`: 年份选择, `8`: 时间段, `9`: 数字, `10`: 默认字段, `11`: 模块
	Value      *GetHireJobRespJobCustomizedDataValue `json:"value,omitempty"`       // 自定义字段值
}

// GetHireJobRespJobCustomizedDataName ...
type GetHireJobRespJobCustomizedDataName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireJobRespJobCustomizedDataValue ...
type GetHireJobRespJobCustomizedDataValue struct {
	Content    string                                         `json:"content,omitempty"`     // 当字段类型为单行文本、多行文本、模块、默认字段时, 从此字段取值
	Option     *GetHireJobRespJobCustomizedDataValueOption    `json:"option,omitempty"`      // 当字段类型为单选时, 从此字段取值
	OptionList []*GetHireJobRespJobCustomizedDataValueOption  `json:"option_list,omitempty"` // 当字段类型为多选时, 从此字段取值
	TimeRange  *GetHireJobRespJobCustomizedDataValueTimeRange `json:"time_range,omitempty"`  // 当字段类型为时间段时, 从此字段取值
	Time       string                                         `json:"time,omitempty"`        // 当字段类型为日期选择、月份选择、年份选择时, 从此字段取值, 该字段是毫秒级时间戳
	Number     string                                         `json:"number,omitempty"`      // 当字段类型为数字时, 从此字段取值
}

// GetHireJobRespJobCustomizedDataValueOption ...
type GetHireJobRespJobCustomizedDataValueOption struct {
	Key  string                                          `json:"key,omitempty"`  // 选项 ID
	Name *GetHireJobRespJobCustomizedDataValueOptionName `json:"name,omitempty"` // 选项名称
}

// GetHireJobRespJobCustomizedDataValueOptionName ...
type GetHireJobRespJobCustomizedDataValueOptionName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireJobRespJobCustomizedDataValueTimeRange ...
type GetHireJobRespJobCustomizedDataValueTimeRange struct {
	StartTime string `json:"start_time,omitempty"` // 开始时间
	EndTime   string `json:"end_time,omitempty"`   // 结束时间
}

// GetHireJobRespJobDepartment ...
type GetHireJobRespJobDepartment struct {
	ID     string `json:"id,omitempty"`      // 部门 ID
	ZhName string `json:"zh_name,omitempty"` // 部门中文名称
	EnName string `json:"en_name,omitempty"` // 部门英文名称
}

// GetHireJobRespJobHighlight ...
type GetHireJobRespJobHighlight struct {
	ID     string `json:"id,omitempty"`      // 职位亮点 ID
	ZhName string `json:"zh_name,omitempty"` // 职位亮点中文名称
	EnName string `json:"en_name,omitempty"` // 职位亮点英文名称
}

// GetHireJobRespJobJobCategory ...
type GetHireJobRespJobJobCategory struct {
	ID           string `json:"id,omitempty"`            // 职位序列 ID
	ZhName       string `json:"zh_name,omitempty"`       // 职位序列中文名称
	EnName       string `json:"en_name,omitempty"`       // 职位序列英文名称
	ActiveStatus int64  `json:"active_status,omitempty"` // 职位序列启用状态, 可选值有: `1`: 启用, `2`: 未启用
}

// GetHireJobRespJobJobFunction ...
type GetHireJobRespJobJobFunction struct {
	ID   string                            `json:"id,omitempty"`   // ID
	Name *GetHireJobRespJobJobFunctionName `json:"name,omitempty"` // 名称
}

// GetHireJobRespJobJobFunctionName ...
type GetHireJobRespJobJobFunctionName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireJobRespJobJobType ...
type GetHireJobRespJobJobType struct {
	ID     string `json:"id,omitempty"`      // 职位类别 ID
	ZhName string `json:"zh_name,omitempty"` // 职位类别中文名称
	EnName string `json:"en_name,omitempty"` // 职位类别英文名称
}

// GetHireJobRespJobMaxJobLevel ...
type GetHireJobRespJobMaxJobLevel struct {
	ID           string `json:"id,omitempty"`            // 职级 ID
	ZhName       string `json:"zh_name,omitempty"`       // 职级中文名称
	EnName       string `json:"en_name,omitempty"`       // 职级英文名称
	ActiveStatus int64  `json:"active_status,omitempty"` // 职级启用状态, 可选值有: `1`: 启用, `2`: 未启用
}

// GetHireJobRespJobMinJobLevel ...
type GetHireJobRespJobMinJobLevel struct {
	ID           string `json:"id,omitempty"`            // 职级 ID
	ZhName       string `json:"zh_name,omitempty"`       // 职级中文名称
	EnName       string `json:"en_name,omitempty"`       // 职级英文名称
	ActiveStatus int64  `json:"active_status,omitempty"` // 职级启用状态, 可选值有: `1`: 启用, `2`: 未启用
}

// GetHireJobRespJobRecruitmentType ...
type GetHireJobRespJobRecruitmentType struct {
	ID           string `json:"id,omitempty"`            // 雇佣类型 ID
	ZhName       string `json:"zh_name,omitempty"`       // 雇佣类型中文名称
	EnName       string `json:"en_name,omitempty"`       // 雇佣类型英文名称
	ActiveStatus int64  `json:"active_status,omitempty"` // 雇佣类型启用状态, 可选值有: `1`: 启用, `2`: 未启用
}

// GetHireJobRespJobSubject ...
type GetHireJobRespJobSubject struct {
	ID   string                        `json:"id,omitempty"`   // ID
	Name *GetHireJobRespJobSubjectName `json:"name,omitempty"` // 名称
}

// GetHireJobRespJobSubjectName ...
type GetHireJobRespJobSubjectName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// getHireJobResp ...
type getHireJobResp struct {
	Code int64           `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 错误描述
	Data *GetHireJobResp `json:"data,omitempty"`
}
