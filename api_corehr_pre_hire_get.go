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

// GetCoreHRPreHire 根据 ID 查询单个待入职人员。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/pre_hire/get
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/pre_hire/get
func (r *CoreHRService) GetCoreHRPreHire(ctx context.Context, request *GetCoreHRPreHireReq, options ...MethodOptionFunc) (*GetCoreHRPreHireResp, *Response, error) {
	if r.cli.mock.mockCoreHRGetCoreHRPreHire != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#GetCoreHRPreHire mock enable")
		return r.cli.mock.mockCoreHRGetCoreHRPreHire(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "GetCoreHRPreHire",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/pre_hires/:pre_hire_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getCoreHRPreHireResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRGetCoreHRPreHire mock CoreHRGetCoreHRPreHire method
func (r *Mock) MockCoreHRGetCoreHRPreHire(f func(ctx context.Context, request *GetCoreHRPreHireReq, options ...MethodOptionFunc) (*GetCoreHRPreHireResp, *Response, error)) {
	r.mockCoreHRGetCoreHRPreHire = f
}

// UnMockCoreHRGetCoreHRPreHire un-mock CoreHRGetCoreHRPreHire method
func (r *Mock) UnMockCoreHRGetCoreHRPreHire() {
	r.mockCoreHRGetCoreHRPreHire = nil
}

// GetCoreHRPreHireReq ...
type GetCoreHRPreHireReq struct {
	PreHireID string `path:"pre_hire_id" json:"-"` // 待入职ID, 示例值: "121215"
}

// GetCoreHRPreHireResp ...
type GetCoreHRPreHireResp struct {
	PreHire *GetCoreHRPreHireRespPreHire `json:"pre_hire,omitempty"` // 待入职信息
}

// GetCoreHRPreHireRespPreHire ...
type GetCoreHRPreHireRespPreHire struct {
	AtsApplicationID string                                       `json:"ats_application_id,omitempty"` // 招聘投递 ID, 详细信息可以通过招聘的[获取投递信息]接口查询获得
	ID               string                                       `json:"id,omitempty"`                 // 待入职ID
	HireDate         string                                       `json:"hire_date,omitempty"`          // 入职日期
	EmployeeType     *GetCoreHRPreHireRespPreHireEmployeeType     `json:"employee_type,omitempty"`      // 雇佣类型
	WorkerID         string                                       `json:"worker_id,omitempty"`          // 人员编号
	EmployeeTypeID   string                                       `json:"employee_type_id,omitempty"`   // 雇佣类型
	PersonID         string                                       `json:"person_id,omitempty"`          // 引用Person ID
	CustomFields     []*GetCoreHRPreHireRespPreHireCustomField    `json:"custom_fields,omitempty"`      // 自定义字段
	CostCenterRate   []*GetCoreHRPreHireRespPreHireCostCenterRate `json:"cost_center_rate,omitempty"`   // 成本中心分摊信息
	OnboardingStatus *GetCoreHRPreHireRespPreHireOnboardingStatus `json:"onboarding_status,omitempty"`  // 入职状态, `preboarding`: 待入职, `day_one`: 准备就绪, `completed`: 已完成, `withdrawn`: 已撤销, `deleted`: 已删除, 对应的系统操作是将待入职人员回退至 Offer 沟通阶段
}

// GetCoreHRPreHireRespPreHireCostCenterRate ...
type GetCoreHRPreHireRespPreHireCostCenterRate struct {
	CostCenterID string `json:"cost_center_id,omitempty"` // 支持的成本中心id
	Rate         int64  `json:"rate,omitempty"`           // 分摊比例
}

// GetCoreHRPreHireRespPreHireCustomField ...
type GetCoreHRPreHireRespPreHireCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05")
}

// GetCoreHRPreHireRespPreHireEmployeeType ...
type GetCoreHRPreHireRespPreHireEmployeeType struct {
	EnumName string                                            `json:"enum_name,omitempty"` // 枚举值
	Display  []*GetCoreHRPreHireRespPreHireEmployeeTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// GetCoreHRPreHireRespPreHireEmployeeTypeDisplay ...
type GetCoreHRPreHireRespPreHireEmployeeTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// GetCoreHRPreHireRespPreHireOnboardingStatus ...
type GetCoreHRPreHireRespPreHireOnboardingStatus struct {
	EnumName string                                                `json:"enum_name,omitempty"` // 枚举值
	Display  []*GetCoreHRPreHireRespPreHireOnboardingStatusDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// GetCoreHRPreHireRespPreHireOnboardingStatusDisplay ...
type GetCoreHRPreHireRespPreHireOnboardingStatusDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// getCoreHRPreHireResp ...
type getCoreHRPreHireResp struct {
	Code  int64                 `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                `json:"msg,omitempty"`  // 错误描述
	Data  *GetCoreHRPreHireResp `json:"data,omitempty"`
	Error *ErrorDetail          `json:"error,omitempty"`
}
