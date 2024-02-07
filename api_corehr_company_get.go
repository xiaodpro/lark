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

// GetCoreHRCompany 根据 ID 查询单个公司。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/company/get
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/organization-management/company/get
func (r *CoreHRService) GetCoreHRCompany(ctx context.Context, request *GetCoreHRCompanyReq, options ...MethodOptionFunc) (*GetCoreHRCompanyResp, *Response, error) {
	if r.cli.mock.mockCoreHRGetCoreHRCompany != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#GetCoreHRCompany mock enable")
		return r.cli.mock.mockCoreHRGetCoreHRCompany(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "GetCoreHRCompany",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/companies/:company_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getCoreHRCompanyResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRGetCoreHRCompany mock CoreHRGetCoreHRCompany method
func (r *Mock) MockCoreHRGetCoreHRCompany(f func(ctx context.Context, request *GetCoreHRCompanyReq, options ...MethodOptionFunc) (*GetCoreHRCompanyResp, *Response, error)) {
	r.mockCoreHRGetCoreHRCompany = f
}

// UnMockCoreHRGetCoreHRCompany un-mock CoreHRGetCoreHRCompany method
func (r *Mock) UnMockCoreHRGetCoreHRCompany() {
	r.mockCoreHRGetCoreHRCompany = nil
}

// GetCoreHRCompanyReq ...
type GetCoreHRCompanyReq struct {
	CompanyID string `path:"company_id" json:"-"` // 公司 ID, 示例值: "151515"
}

// GetCoreHRCompanyResp ...
type GetCoreHRCompanyResp struct {
	Company *GetCoreHRCompanyRespCompany `json:"company,omitempty"` // 公司信息
}

// GetCoreHRCompanyRespCompany ...
type GetCoreHRCompanyRespCompany struct {
	ID                      string                                               `json:"id,omitempty"`                        // 公司 ID
	HiberarchyCommon        *GetCoreHRCompanyRespCompanyHiberarchyCommon         `json:"hiberarchy_common,omitempty"`         // 层级关系, 内层字段见实体
	Type                    *GetCoreHRCompanyRespCompanyType                     `json:"type,omitempty"`                      // 性质, 枚举值可通过文档[【飞书人事枚举常量】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)公司类型（company_type）枚举定义部分获得
	IndustryList            []*GetCoreHRCompanyRespCompanyIndustry               `json:"industry_list,omitempty"`             // 行业, 枚举值可通过文档[【飞书人事枚举常量】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)行业（industry）枚举定义部分获得
	LegalRepresentative     []*GetCoreHRCompanyRespCompanyLegalRepresentative    `json:"legal_representative,omitempty"`      // 法定代表人
	PostCode                string                                               `json:"post_code,omitempty"`                 // 邮编
	TaxPayerID              string                                               `json:"tax_payer_id,omitempty"`              // 纳税人识别号
	Confidential            bool                                                 `json:"confidential,omitempty"`              // 是否保密
	SubTypeList             []*GetCoreHRCompanyRespCompanySubType                `json:"sub_type_list,omitempty"`             // 主体类型, 枚举值可通过文档[【飞书人事枚举常量】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)主体类型（company_sub_type）枚举定义部分获得
	BranchCompany           bool                                                 `json:"branch_company,omitempty"`            // 是否为分公司
	PrimaryManager          []*GetCoreHRCompanyRespCompanyPrimaryManager         `json:"primary_manager,omitempty"`           // 主要负责人
	CustomFields            []*GetCoreHRCompanyRespCompanyCustomField            `json:"custom_fields,omitempty"`             // 自定义字段
	Currency                *GetCoreHRCompanyRespCompanyCurrency                 `json:"currency,omitempty"`                  // 默认币种
	Phone                   *GetCoreHRCompanyRespCompanyPhone                    `json:"phone,omitempty"`                     // 电话
	Fax                     *GetCoreHRCompanyRespCompanyFax                      `json:"fax,omitempty"`                       // 传真
	RegisteredOfficeAddress []*GetCoreHRCompanyRespCompanyRegisteredOfficeAddres `json:"registered_office_address,omitempty"` // 注册地址
	OfficeAddress           []*GetCoreHRCompanyRespCompanyOfficeAddres           `json:"office_address,omitempty"`            // 办公地址
}

// GetCoreHRCompanyRespCompanyCurrency ...
type GetCoreHRCompanyRespCompanyCurrency struct {
	ID                 string                                             `json:"id,omitempty"`                    // 货币id
	CountryRegionID    string                                             `json:"country_region_id,omitempty"`     // 货币所属国家/地区id, 详细信息可通过[【查询国家/地区信息】](https://open.feishu.cn/document/server-docs/corehr-v1/basic-infomation/location_data/list)接口查询获得
	CurrencyName       []*GetCoreHRCompanyRespCompanyCurrencyCurrencyName `json:"currency_name,omitempty"`         // 货币名称
	NumericCode        int64                                              `json:"numeric_code,omitempty"`          // 数字代码
	CurrencyAlpha3Code string                                             `json:"currency_alpha_3_code,omitempty"` // 三位字母代码
}

// GetCoreHRCompanyRespCompanyCurrencyCurrencyName ...
type GetCoreHRCompanyRespCompanyCurrencyCurrencyName struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// GetCoreHRCompanyRespCompanyCustomField ...
type GetCoreHRCompanyRespCompanyCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05")
}

// GetCoreHRCompanyRespCompanyFax ...
type GetCoreHRCompanyRespCompanyFax struct {
	AreaCode    *GetCoreHRCompanyRespCompanyFaxAreaCode `json:"area_code,omitempty"`    // 区号
	PhoneNumber string                                  `json:"phone_number,omitempty"` // 号码
}

// GetCoreHRCompanyRespCompanyFaxAreaCode ...
type GetCoreHRCompanyRespCompanyFaxAreaCode struct {
	EnumName string                                           `json:"enum_name,omitempty"` // 枚举值
	Display  []*GetCoreHRCompanyRespCompanyFaxAreaCodeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// GetCoreHRCompanyRespCompanyFaxAreaCodeDisplay ...
type GetCoreHRCompanyRespCompanyFaxAreaCodeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// GetCoreHRCompanyRespCompanyHiberarchyCommon ...
type GetCoreHRCompanyRespCompanyHiberarchyCommon struct {
	ParentID       string                                                    `json:"parent_id,omitempty"`       // 上级
	Name           []*GetCoreHRCompanyRespCompanyHiberarchyCommonName        `json:"name,omitempty"`            // 名称
	Type           *GetCoreHRCompanyRespCompanyHiberarchyCommonType          `json:"type,omitempty"`            // 组织类型, 枚举值可通过文档[【飞书人事枚举常量】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)组织类型（organization_type）枚举定义部分获得
	Active         bool                                                      `json:"active,omitempty"`          // 是否启用
	EffectiveTime  string                                                    `json:"effective_time,omitempty"`  // 生效时间
	ExpirationTime string                                                    `json:"expiration_time,omitempty"` // 失效时间
	Code           string                                                    `json:"code,omitempty"`            // 编码
	Description    []*GetCoreHRCompanyRespCompanyHiberarchyCommonDescription `json:"description,omitempty"`     // 描述
	TreeOrder      string                                                    `json:"tree_order,omitempty"`      // 树形排序
	ListOrder      string                                                    `json:"list_order,omitempty"`      // 列表排序
	CustomFields   []*GetCoreHRCompanyRespCompanyHiberarchyCommonCustomField `json:"custom_fields,omitempty"`   // 自定义字段
}

// GetCoreHRCompanyRespCompanyHiberarchyCommonCustomField ...
type GetCoreHRCompanyRespCompanyHiberarchyCommonCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05")
}

// GetCoreHRCompanyRespCompanyHiberarchyCommonDescription ...
type GetCoreHRCompanyRespCompanyHiberarchyCommonDescription struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// GetCoreHRCompanyRespCompanyHiberarchyCommonName ...
type GetCoreHRCompanyRespCompanyHiberarchyCommonName struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// GetCoreHRCompanyRespCompanyHiberarchyCommonType ...
type GetCoreHRCompanyRespCompanyHiberarchyCommonType struct {
	EnumName string                                                    `json:"enum_name,omitempty"` // 枚举值
	Display  []*GetCoreHRCompanyRespCompanyHiberarchyCommonTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// GetCoreHRCompanyRespCompanyHiberarchyCommonTypeDisplay ...
type GetCoreHRCompanyRespCompanyHiberarchyCommonTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// GetCoreHRCompanyRespCompanyIndustry ...
type GetCoreHRCompanyRespCompanyIndustry struct {
	EnumName string                                        `json:"enum_name,omitempty"` // 枚举值
	Display  []*GetCoreHRCompanyRespCompanyIndustryDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// GetCoreHRCompanyRespCompanyIndustryDisplay ...
type GetCoreHRCompanyRespCompanyIndustryDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// GetCoreHRCompanyRespCompanyLegalRepresentative ...
type GetCoreHRCompanyRespCompanyLegalRepresentative struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// GetCoreHRCompanyRespCompanyOfficeAddres ...
type GetCoreHRCompanyRespCompanyOfficeAddres struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// GetCoreHRCompanyRespCompanyPhone ...
type GetCoreHRCompanyRespCompanyPhone struct {
	AreaCode    *GetCoreHRCompanyRespCompanyPhoneAreaCode `json:"area_code,omitempty"`    // 区号
	PhoneNumber string                                    `json:"phone_number,omitempty"` // 号码
}

// GetCoreHRCompanyRespCompanyPhoneAreaCode ...
type GetCoreHRCompanyRespCompanyPhoneAreaCode struct {
	EnumName string                                             `json:"enum_name,omitempty"` // 枚举值
	Display  []*GetCoreHRCompanyRespCompanyPhoneAreaCodeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// GetCoreHRCompanyRespCompanyPhoneAreaCodeDisplay ...
type GetCoreHRCompanyRespCompanyPhoneAreaCodeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// GetCoreHRCompanyRespCompanyPrimaryManager ...
type GetCoreHRCompanyRespCompanyPrimaryManager struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// GetCoreHRCompanyRespCompanyRegisteredOfficeAddres ...
type GetCoreHRCompanyRespCompanyRegisteredOfficeAddres struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// GetCoreHRCompanyRespCompanySubType ...
type GetCoreHRCompanyRespCompanySubType struct {
	EnumName string                                       `json:"enum_name,omitempty"` // 枚举值
	Display  []*GetCoreHRCompanyRespCompanySubTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// GetCoreHRCompanyRespCompanySubTypeDisplay ...
type GetCoreHRCompanyRespCompanySubTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// GetCoreHRCompanyRespCompanyType ...
type GetCoreHRCompanyRespCompanyType struct {
	EnumName string                                    `json:"enum_name,omitempty"` // 枚举值
	Display  []*GetCoreHRCompanyRespCompanyTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// GetCoreHRCompanyRespCompanyTypeDisplay ...
type GetCoreHRCompanyRespCompanyTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// getCoreHRCompanyResp ...
type getCoreHRCompanyResp struct {
	Code int64                 `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                `json:"msg,omitempty"`  // 错误描述
	Data *GetCoreHRCompanyResp `json:"data,omitempty"`
}
