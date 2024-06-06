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

// UpdateCoreHRContract 更新合同。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/contract/patch
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/contract/patch
func (r *CoreHRService) UpdateCoreHRContract(ctx context.Context, request *UpdateCoreHRContractReq, options ...MethodOptionFunc) (*UpdateCoreHRContractResp, *Response, error) {
	if r.cli.mock.mockCoreHRUpdateCoreHRContract != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#UpdateCoreHRContract mock enable")
		return r.cli.mock.mockCoreHRUpdateCoreHRContract(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "UpdateCoreHRContract",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/contracts/:contract_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateCoreHRContractResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRUpdateCoreHRContract mock CoreHRUpdateCoreHRContract method
func (r *Mock) MockCoreHRUpdateCoreHRContract(f func(ctx context.Context, request *UpdateCoreHRContractReq, options ...MethodOptionFunc) (*UpdateCoreHRContractResp, *Response, error)) {
	r.mockCoreHRUpdateCoreHRContract = f
}

// UnMockCoreHRUpdateCoreHRContract un-mock CoreHRUpdateCoreHRContract method
func (r *Mock) UnMockCoreHRUpdateCoreHRContract() {
	r.mockCoreHRUpdateCoreHRContract = nil
}

// UpdateCoreHRContractReq ...
type UpdateCoreHRContractReq struct {
	ContractID          string                                `path:"contract_id" json:"-"`             // 合同ID, 示例值: "1616161616"
	ClientToken         *string                               `query:"client_token" json:"-"`           // 根据client_token是否一致来判断是否为同一请求, 示例值: 12454646
	EffectiveTime       *string                               `json:"effective_time,omitempty"`         // 合同开始日期, 示例值: "2050-01-01 00:00:00"
	ExpirationTime      *string                               `json:"expiration_time,omitempty"`        // 实际结束日期, 示例值: "9999-12-31 23:59:59"
	EmploymentID        *string                               `json:"employment_id,omitempty"`          // 雇员 ID, 枚举值及详细信息可通过[批量查询雇佣信息]接口查询获得, 示例值: "6893013238632416776"
	ContractType        *UpdateCoreHRContractReqContractType  `json:"contract_type,omitempty"`          // 合同类型, 枚举值可通过文档[飞书人事枚举常量](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)合同类型（contract_type）枚举定义部分获得
	FirstPartyCompanyID *string                               `json:"first_party_company_id,omitempty"` // 甲方, 引用Company的ID, 枚举值及详细信息可通过[批量查询公司]接口查询获得, 示例值: "6892686614112241165"
	PersonID            *string                               `json:"person_id,omitempty"`              // Person ID, 枚举值及详细信息可通过[批量查询个人信息]接口查询获得, 示例值: "151515151"
	CustomFields        []*UpdateCoreHRContractReqCustomField `json:"custom_fields,omitempty"`          // 自定义字段
	DurationType        *UpdateCoreHRContractReqDurationType  `json:"duration_type,omitempty"`          // 期限类型, 枚举值可通过文档[飞书人事枚举常量](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)合同期限类型（duration_type）枚举定义部分获得, 示例值: fixed_term
	ContractEndDate     *string                               `json:"contract_end_date,omitempty"`      // 合同结束日期, 示例值: "2006-01-02"
	ContractNumber      *string                               `json:"contract_number,omitempty"`        // 合同编号, 示例值: "6919737965274990093"
	SigningType         *UpdateCoreHRContractReqSigningType   `json:"signing_type,omitempty"`           // 签订类型, 枚举值可通过文档[飞书人事枚举常量](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)签订类型（signing_type）枚举定义部分获得
}

// UpdateCoreHRContractReqContractType ...
type UpdateCoreHRContractReqContractType struct {
	EnumName string `json:"enum_name,omitempty"` // 枚举值, 示例值: "type_1"
}

// UpdateCoreHRContractReqCustomField ...
type UpdateCoreHRContractReqCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名, 示例值: "name"
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05"), 示例值: "Sandy"
}

// UpdateCoreHRContractReqDurationType ...
type UpdateCoreHRContractReqDurationType struct {
	EnumName string `json:"enum_name,omitempty"` // 枚举值, 示例值: "type_1"
}

// UpdateCoreHRContractReqSigningType ...
type UpdateCoreHRContractReqSigningType struct {
	EnumName string `json:"enum_name,omitempty"` // 枚举值, 示例值: "type_1"
}

// UpdateCoreHRContractResp ...
type UpdateCoreHRContractResp struct {
	Contract *UpdateCoreHRContractRespContract `json:"contract,omitempty"` // 合同
}

// UpdateCoreHRContractRespContract ...
type UpdateCoreHRContractRespContract struct {
	ID                  string                                         `json:"id,omitempty"`                     // 合同ID
	EffectiveTime       string                                         `json:"effective_time,omitempty"`         // 合同开始日期
	ExpirationTime      string                                         `json:"expiration_time,omitempty"`        // 实际结束日期
	EmploymentID        string                                         `json:"employment_id,omitempty"`          // 雇员 ID, 枚举值及详细信息可通过[批量查询雇佣信息]接口查询获得
	ContractType        *UpdateCoreHRContractRespContractContractType  `json:"contract_type,omitempty"`          // 合同类型, 枚举值可通过文档[飞书人事枚举常量](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)合同类型（contract_type）枚举定义部分获得
	FirstPartyCompanyID string                                         `json:"first_party_company_id,omitempty"` // 甲方, 引用Company的ID, 枚举值及详细信息可通过[批量查询公司]接口查询获得
	PersonID            string                                         `json:"person_id,omitempty"`              // Person ID, 枚举值及详细信息可通过[批量查询个人信息]接口查询获得
	CustomFields        []*UpdateCoreHRContractRespContractCustomField `json:"custom_fields,omitempty"`          // 自定义字段
	DurationType        *UpdateCoreHRContractRespContractDurationType  `json:"duration_type,omitempty"`          // 期限类型, 枚举值可通过文档[飞书人事枚举常量](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)合同期限类型（duration_type）枚举定义部分获得
	ContractEndDate     string                                         `json:"contract_end_date,omitempty"`      // 合同结束日期
	ContractNumber      string                                         `json:"contract_number,omitempty"`        // 合同编号
	SigningType         *UpdateCoreHRContractRespContractSigningType   `json:"signing_type,omitempty"`           // 签订类型, 枚举值可通过文档[飞书人事枚举常量](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)签订类型（signing_type）枚举定义部分获得
}

// UpdateCoreHRContractRespContractContractType ...
type UpdateCoreHRContractRespContractContractType struct {
	EnumName string                                                 `json:"enum_name,omitempty"` // 枚举值
	Display  []*UpdateCoreHRContractRespContractContractTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// UpdateCoreHRContractRespContractContractTypeDisplay ...
type UpdateCoreHRContractRespContractContractTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// UpdateCoreHRContractRespContractCustomField ...
type UpdateCoreHRContractRespContractCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05")
}

// UpdateCoreHRContractRespContractDurationType ...
type UpdateCoreHRContractRespContractDurationType struct {
	EnumName string                                                 `json:"enum_name,omitempty"` // 枚举值
	Display  []*UpdateCoreHRContractRespContractDurationTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// UpdateCoreHRContractRespContractDurationTypeDisplay ...
type UpdateCoreHRContractRespContractDurationTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// UpdateCoreHRContractRespContractSigningType ...
type UpdateCoreHRContractRespContractSigningType struct {
	EnumName string                                                `json:"enum_name,omitempty"` // 枚举值
	Display  []*UpdateCoreHRContractRespContractSigningTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// UpdateCoreHRContractRespContractSigningTypeDisplay ...
type UpdateCoreHRContractRespContractSigningTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// updateCoreHRContractResp ...
type updateCoreHRContractResp struct {
	Code  int64                     `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                    `json:"msg,omitempty"`  // 错误描述
	Data  *UpdateCoreHRContractResp `json:"data,omitempty"`
	Error *ErrorDetail              `json:"error,omitempty"`
}
