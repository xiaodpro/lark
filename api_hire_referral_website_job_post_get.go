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

// GetHireReferralWebsiteJobPost 根据广告 ID 获取内推官网下的职位广告详情。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/referral_website-job_post/get
// new doc: https://open.feishu.cn/document/server-docs/hire-v1/get-candidates/referral/get
func (r *HireService) GetHireReferralWebsiteJobPost(ctx context.Context, request *GetHireReferralWebsiteJobPostReq, options ...MethodOptionFunc) (*GetHireReferralWebsiteJobPostResp, *Response, error) {
	if r.cli.mock.mockHireGetHireReferralWebsiteJobPost != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Hire#GetHireReferralWebsiteJobPost mock enable")
		return r.cli.mock.mockHireGetHireReferralWebsiteJobPost(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "GetHireReferralWebsiteJobPost",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/hire/v1/referral_websites/job_posts/:job_post_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getHireReferralWebsiteJobPostResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHireGetHireReferralWebsiteJobPost mock HireGetHireReferralWebsiteJobPost method
func (r *Mock) MockHireGetHireReferralWebsiteJobPost(f func(ctx context.Context, request *GetHireReferralWebsiteJobPostReq, options ...MethodOptionFunc) (*GetHireReferralWebsiteJobPostResp, *Response, error)) {
	r.mockHireGetHireReferralWebsiteJobPost = f
}

// UnMockHireGetHireReferralWebsiteJobPost un-mock HireGetHireReferralWebsiteJobPost method
func (r *Mock) UnMockHireGetHireReferralWebsiteJobPost() {
	r.mockHireGetHireReferralWebsiteJobPost = nil
}

// GetHireReferralWebsiteJobPostReq ...
type GetHireReferralWebsiteJobPostReq struct {
	JobPostID        string            `path:"job_post_id" json:"-"`         // 职位广告 ID, 示例值: "6701528341100366094"
	UserIDType       *IDType           `query:"user_id_type" json:"-"`       // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	DepartmentIDType *DepartmentIDType `query:"department_id_type" json:"-"` // 此次调用中使用的部门 ID 的类型, 示例值: open_department_id, 可选值有: open_department_id: 以 open_department_id 来标识部门, department_id: 以 department_id 来标识部门, 默认值: `open_department_id`
	JobLevelIDType   *IDType           `query:"job_level_id_type" json:"-"`  // 此次调用中使用的「职级 ID」的类型, 示例值: 6942778198054125570, 可选值有: people_admin_job_level_id: 「人力系统管理后台」适用的职级 ID。人力系统管理后台逐步下线中, 建议不继续使用此 ID。, job_level_id: 「飞书管理后台」适用的职级 ID, 通过[「获取租户职级列表」](https://open.feishu.cn/document/server-docs/contact-v3/job_level/list)接口获取, 默认值: `people_admin_job_level_id`
}

// GetHireReferralWebsiteJobPostResp ...
type GetHireReferralWebsiteJobPostResp struct {
	JobPost *GetHireReferralWebsiteJobPostRespJobPost `json:"job_post,omitempty"` // 职位广告信息
}

// GetHireReferralWebsiteJobPostRespJobPost ...
type GetHireReferralWebsiteJobPostRespJobPost struct {
	ID                 string                                                      `json:"id,omitempty"`                   // 职位广告 ID
	Title              string                                                      `json:"title,omitempty"`                // 标题
	JobID              string                                                      `json:"job_id,omitempty"`               // 职位 ID
	JobCode            string                                                      `json:"job_code,omitempty"`             // 职位编码
	JobExpireTime      string                                                      `json:"job_expire_time,omitempty"`      // 职位过期时间, 「null」代表「长期有效」
	JobActiveStatus    int64                                                       `json:"job_active_status,omitempty"`    // 职位状态, 可选值有: 1: 启用态, 2: 禁用态
	JobProcessType     int64                                                       `json:"job_process_type,omitempty"`     // 职位流程类型, 可选值有: 1: 社招, 2: 校招
	JobRecruitmentType *GetHireReferralWebsiteJobPostRespJobPostJobRecruitmentType `json:"job_recruitment_type,omitempty"` // 职位雇佣类型
	JobDepartment      *GetHireReferralWebsiteJobPostRespJobPostJobDepartment      `json:"job_department,omitempty"`       // 职位部门
	JobType            *GetHireReferralWebsiteJobPostRespJobPostJobType            `json:"job_type,omitempty"`             // 职位类型
	MinJobLevel        *GetHireReferralWebsiteJobPostRespJobPostMinJobLevel        `json:"min_job_level,omitempty"`        // 最低职级
	MaxJobLevel        *GetHireReferralWebsiteJobPostRespJobPostMaxJobLevel        `json:"max_job_level,omitempty"`        // 最高职级
	Address            *GetHireReferralWebsiteJobPostRespJobPostAddress            `json:"address,omitempty"`              // 职位地址
	MinSalary          string                                                      `json:"min_salary,omitempty"`           // 月薪范围-最低薪资
	MaxSalary          string                                                      `json:"max_salary,omitempty"`           // 月薪范围-最高薪资
	RequiredDegree     int64                                                       `json:"required_degree,omitempty"`      // 学历要求, 可选值有: 1: 小学及以上, 2: 初中及以上, 3: 专职及以上, 4: 高中及以上, 5: 大专及以上, 6: 本科及以上, 7: 硕士及以上, 8: 博士及以上, 20: 不限
	Experience         int64                                                       `json:"experience,omitempty"`           // 经验, 可选值有: 1: 不限, 2: 应届毕业生, 3: 1年以下, 4: 1-3年, 5: 3-5年, 6: 5-7年, 7: 7-10年, 8: 10年以上
	Headcount          int64                                                       `json:"headcount,omitempty"`            // 数量
	HighLightList      []*GetHireReferralWebsiteJobPostRespJobPostHighLight        `json:"high_light_list,omitempty"`      // 职位亮点
	Description        string                                                      `json:"description,omitempty"`          // 职位描述
	Requirement        string                                                      `json:"requirement,omitempty"`          // 职位要求
	Creator            *GetHireReferralWebsiteJobPostRespJobPostCreator            `json:"creator,omitempty"`              // 创建人
	CreateTime         string                                                      `json:"create_time,omitempty"`          // 创建时间
	ModifyTime         string                                                      `json:"modify_time,omitempty"`          // 修改时间
	CustomizedDataList []*GetHireReferralWebsiteJobPostRespJobPostCustomizedData   `json:"customized_data_list,omitempty"` // 自定义字段
	JobFunction        *GetHireReferralWebsiteJobPostRespJobPostJobFunction        `json:"job_function,omitempty"`         // 职能分类
	Subject            *GetHireReferralWebsiteJobPostRespJobPostSubject            `json:"subject,omitempty"`              // 职位项目
	AddressList        []*GetHireReferralWebsiteJobPostRespJobPostAddress          `json:"address_list,omitempty"`         // 职位广告地址列表
}

// GetHireReferralWebsiteJobPostRespJobPostAddress ...
type GetHireReferralWebsiteJobPostRespJobPostAddress struct {
	ID       string                                                   `json:"id,omitempty"`       // ID
	Name     *GetHireReferralWebsiteJobPostRespJobPostAddressName     `json:"name,omitempty"`     // 名称
	District *GetHireReferralWebsiteJobPostRespJobPostAddressDistrict `json:"district,omitempty"` // 区域信息
	City     *GetHireReferralWebsiteJobPostRespJobPostAddressCity     `json:"city,omitempty"`     // 城市信息
	State    *GetHireReferralWebsiteJobPostRespJobPostAddressState    `json:"state,omitempty"`    // 省信息
	Country  *GetHireReferralWebsiteJobPostRespJobPostAddressCountry  `json:"country,omitempty"`  // 国家信息
}

// GetHireReferralWebsiteJobPostRespJobPostAddressCity ...
type GetHireReferralWebsiteJobPostRespJobPostAddressCity struct {
	Code string                                                   `json:"code,omitempty"` // 编码
	Name *GetHireReferralWebsiteJobPostRespJobPostAddressCityName `json:"name,omitempty"` // 名称
}

// GetHireReferralWebsiteJobPostRespJobPostAddressCityName ...
type GetHireReferralWebsiteJobPostRespJobPostAddressCityName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireReferralWebsiteJobPostRespJobPostAddressCountry ...
type GetHireReferralWebsiteJobPostRespJobPostAddressCountry struct {
	Code string                                                      `json:"code,omitempty"` // 编码
	Name *GetHireReferralWebsiteJobPostRespJobPostAddressCountryName `json:"name,omitempty"` // 名称
}

// GetHireReferralWebsiteJobPostRespJobPostAddressCountryName ...
type GetHireReferralWebsiteJobPostRespJobPostAddressCountryName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireReferralWebsiteJobPostRespJobPostAddressDistrict ...
type GetHireReferralWebsiteJobPostRespJobPostAddressDistrict struct {
	Code string                                                       `json:"code,omitempty"` // 编码
	Name *GetHireReferralWebsiteJobPostRespJobPostAddressDistrictName `json:"name,omitempty"` // 名称
}

// GetHireReferralWebsiteJobPostRespJobPostAddressDistrictName ...
type GetHireReferralWebsiteJobPostRespJobPostAddressDistrictName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireReferralWebsiteJobPostRespJobPostAddressName ...
type GetHireReferralWebsiteJobPostRespJobPostAddressName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireReferralWebsiteJobPostRespJobPostAddressState ...
type GetHireReferralWebsiteJobPostRespJobPostAddressState struct {
	Code string                                                    `json:"code,omitempty"` // 编码
	Name *GetHireReferralWebsiteJobPostRespJobPostAddressStateName `json:"name,omitempty"` // 名称
}

// GetHireReferralWebsiteJobPostRespJobPostAddressStateName ...
type GetHireReferralWebsiteJobPostRespJobPostAddressStateName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireReferralWebsiteJobPostRespJobPostCreator ...
type GetHireReferralWebsiteJobPostRespJobPostCreator struct {
	ID   string                                               `json:"id,omitempty"`   // ID
	Name *GetHireReferralWebsiteJobPostRespJobPostCreatorName `json:"name,omitempty"` // 名称
}

// GetHireReferralWebsiteJobPostRespJobPostCreatorName ...
type GetHireReferralWebsiteJobPostRespJobPostCreatorName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireReferralWebsiteJobPostRespJobPostCustomizedData ...
type GetHireReferralWebsiteJobPostRespJobPostCustomizedData struct {
	ObjectID   string                                                       `json:"object_id,omitempty"`   // 自定义字段 ID
	Name       *GetHireReferralWebsiteJobPostRespJobPostCustomizedDataName  `json:"name,omitempty"`        // 字段名称
	ObjectType int64                                                        `json:"object_type,omitempty"` // 字段类型, 可选值有: 1: 单行文本, 2: 多行文本, 3: 单选, 4: 多选, 5: 日期, 6: 月份选择, 7: 年份选择, 8: 时间段, 9: 数字, 10: 默认字段
	Value      *GetHireReferralWebsiteJobPostRespJobPostCustomizedDataValue `json:"value,omitempty"`       // 自定义字段值
}

// GetHireReferralWebsiteJobPostRespJobPostCustomizedDataName ...
type GetHireReferralWebsiteJobPostRespJobPostCustomizedDataName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireReferralWebsiteJobPostRespJobPostCustomizedDataValue ...
type GetHireReferralWebsiteJobPostRespJobPostCustomizedDataValue struct {
	Content    string                                                                `json:"content,omitempty"`     // 当字段类型为单行文本、多行文本、模块、默认字段时, 从此字段取值
	Option     *GetHireReferralWebsiteJobPostRespJobPostCustomizedDataValueOption    `json:"option,omitempty"`      // 当字段类型为单选时, 从此字段取值
	OptionList []*GetHireReferralWebsiteJobPostRespJobPostCustomizedDataValueOption  `json:"option_list,omitempty"` // 当字段类型为多选时, 从此字段取值
	TimeRange  *GetHireReferralWebsiteJobPostRespJobPostCustomizedDataValueTimeRange `json:"time_range,omitempty"`  // 当字段类型为时间段时, 从此字段取值
	Time       string                                                                `json:"time,omitempty"`        // 当字段类型为日期选择、月份选择、年份选择时, 从此字段取值, 该字段是毫秒级时间戳
	Number     string                                                                `json:"number,omitempty"`      // 当字段类型为数字时, 从此字段取值
}

// GetHireReferralWebsiteJobPostRespJobPostCustomizedDataValueOption ...
type GetHireReferralWebsiteJobPostRespJobPostCustomizedDataValueOption struct {
	Key  string                                                                 `json:"key,omitempty"`  // 选项 ID
	Name *GetHireReferralWebsiteJobPostRespJobPostCustomizedDataValueOptionName `json:"name,omitempty"` // 选项名称
}

// GetHireReferralWebsiteJobPostRespJobPostCustomizedDataValueOptionName ...
type GetHireReferralWebsiteJobPostRespJobPostCustomizedDataValueOptionName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireReferralWebsiteJobPostRespJobPostCustomizedDataValueTimeRange ...
type GetHireReferralWebsiteJobPostRespJobPostCustomizedDataValueTimeRange struct {
	StartTime string `json:"start_time,omitempty"` // 开始时间
	EndTime   string `json:"end_time,omitempty"`   // 结束时间
}

// GetHireReferralWebsiteJobPostRespJobPostHighLight ...
type GetHireReferralWebsiteJobPostRespJobPostHighLight struct {
	ID   string                                                 `json:"id,omitempty"`   // ID
	Name *GetHireReferralWebsiteJobPostRespJobPostHighLightName `json:"name,omitempty"` // 名称
}

// GetHireReferralWebsiteJobPostRespJobPostHighLightName ...
type GetHireReferralWebsiteJobPostRespJobPostHighLightName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireReferralWebsiteJobPostRespJobPostJobDepartment ...
type GetHireReferralWebsiteJobPostRespJobPostJobDepartment struct {
	ID   string                                                     `json:"id,omitempty"`   // ID
	Name *GetHireReferralWebsiteJobPostRespJobPostJobDepartmentName `json:"name,omitempty"` // 名称
}

// GetHireReferralWebsiteJobPostRespJobPostJobDepartmentName ...
type GetHireReferralWebsiteJobPostRespJobPostJobDepartmentName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireReferralWebsiteJobPostRespJobPostJobFunction ...
type GetHireReferralWebsiteJobPostRespJobPostJobFunction struct {
	ID   string                                                   `json:"id,omitempty"`   // ID
	Name *GetHireReferralWebsiteJobPostRespJobPostJobFunctionName `json:"name,omitempty"` // 名称
}

// GetHireReferralWebsiteJobPostRespJobPostJobFunctionName ...
type GetHireReferralWebsiteJobPostRespJobPostJobFunctionName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireReferralWebsiteJobPostRespJobPostJobRecruitmentType ...
type GetHireReferralWebsiteJobPostRespJobPostJobRecruitmentType struct {
	ID   string                                                          `json:"id,omitempty"`   // ID
	Name *GetHireReferralWebsiteJobPostRespJobPostJobRecruitmentTypeName `json:"name,omitempty"` // 名称
}

// GetHireReferralWebsiteJobPostRespJobPostJobRecruitmentTypeName ...
type GetHireReferralWebsiteJobPostRespJobPostJobRecruitmentTypeName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireReferralWebsiteJobPostRespJobPostJobType ...
type GetHireReferralWebsiteJobPostRespJobPostJobType struct {
	ID   string                                               `json:"id,omitempty"`   // ID
	Name *GetHireReferralWebsiteJobPostRespJobPostJobTypeName `json:"name,omitempty"` // 名称
}

// GetHireReferralWebsiteJobPostRespJobPostJobTypeName ...
type GetHireReferralWebsiteJobPostRespJobPostJobTypeName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetHireReferralWebsiteJobPostRespJobPostMaxJobLevel ...
type GetHireReferralWebsiteJobPostRespJobPostMaxJobLevel struct {
	ID   string                                                   `json:"id,omitempty"`   // 职级ID
	Name *GetHireReferralWebsiteJobPostRespJobPostMaxJobLevelName `json:"name,omitempty"` // 职级名称
}

// GetHireReferralWebsiteJobPostRespJobPostMaxJobLevelName ...
type GetHireReferralWebsiteJobPostRespJobPostMaxJobLevelName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 职级中文名称
	EnUs string `json:"en_us,omitempty"` // 职级英文名称
}

// GetHireReferralWebsiteJobPostRespJobPostMinJobLevel ...
type GetHireReferralWebsiteJobPostRespJobPostMinJobLevel struct {
	ID   string                                                   `json:"id,omitempty"`   // 职级ID
	Name *GetHireReferralWebsiteJobPostRespJobPostMinJobLevelName `json:"name,omitempty"` // 职级名称
}

// GetHireReferralWebsiteJobPostRespJobPostMinJobLevelName ...
type GetHireReferralWebsiteJobPostRespJobPostMinJobLevelName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 职级中文名称
	EnUs string `json:"en_us,omitempty"` // 职级英文名称
}

// GetHireReferralWebsiteJobPostRespJobPostSubject ...
type GetHireReferralWebsiteJobPostRespJobPostSubject struct {
	ID   string                                               `json:"id,omitempty"`   // ID
	Name *GetHireReferralWebsiteJobPostRespJobPostSubjectName `json:"name,omitempty"` // 名称
}

// GetHireReferralWebsiteJobPostRespJobPostSubjectName ...
type GetHireReferralWebsiteJobPostRespJobPostSubjectName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// getHireReferralWebsiteJobPostResp ...
type getHireReferralWebsiteJobPostResp struct {
	Code  int64                              `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                             `json:"msg,omitempty"`  // 错误描述
	Data  *GetHireReferralWebsiteJobPostResp `json:"data,omitempty"`
	Error *ErrorDetail                       `json:"error,omitempty"`
}