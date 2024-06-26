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

// GetCoreHRCompanyList 批量查询公司。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/company/list
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/organization-management/company/list
func (r *CoreHRService) GetCoreHRCompanyList(ctx context.Context, request *GetCoreHRCompanyListReq, options ...MethodOptionFunc) (*GetCoreHRCompanyListResp, *Response, error) {
	if r.cli.mock.mockCoreHRGetCoreHRCompanyList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#GetCoreHRCompanyList mock enable")
		return r.cli.mock.mockCoreHRGetCoreHRCompanyList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "GetCoreHRCompanyList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/companies",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getCoreHRCompanyListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRGetCoreHRCompanyList mock CoreHRGetCoreHRCompanyList method
func (r *Mock) MockCoreHRGetCoreHRCompanyList(f func(ctx context.Context, request *GetCoreHRCompanyListReq, options ...MethodOptionFunc) (*GetCoreHRCompanyListResp, *Response, error)) {
	r.mockCoreHRGetCoreHRCompanyList = f
}

// UnMockCoreHRGetCoreHRCompanyList un-mock CoreHRGetCoreHRCompanyList method
func (r *Mock) UnMockCoreHRGetCoreHRCompanyList() {
	r.mockCoreHRGetCoreHRCompanyList = nil
}

// GetCoreHRCompanyListReq ...
type GetCoreHRCompanyListReq struct {
	PageToken *string `query:"page_token" json:"-"` // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: 1231231987
	PageSize  int64   `query:"page_size" json:"-"`  // 分页大小, 示例值: 100
}

// GetCoreHRCompanyListResp ...
type GetCoreHRCompanyListResp struct {
	Items     []*GetCoreHRCompanyListRespItem `json:"items,omitempty"`      // 查询的公司信息
	HasMore   bool                            `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                          `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
}

// GetCoreHRCompanyListRespItem ...
type GetCoreHRCompanyListRespItem struct {
	ID                          string                                                   `json:"id,omitempty"`                             // 公司 ID
	HiberarchyCommon            *GetCoreHRCompanyListRespItemHiberarchyCommon            `json:"hiberarchy_common,omitempty"`              // 层级关系, 内层字段见实体
	Type                        *GetCoreHRCompanyListRespItemType                        `json:"type,omitempty"`                           // 性质, 枚举值可通过文档[飞书人事枚举常量](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)公司类型（company_type）枚举定义部分获得
	IndustryList                []*GetCoreHRCompanyListRespItemIndustry                  `json:"industry_list,omitempty"`                  // 行业, 枚举值可通过文档[飞书人事枚举常量](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)行业（industry）枚举定义部分获得
	LegalRepresentative         []*GetCoreHRCompanyListRespItemLegalRepresentative       `json:"legal_representative,omitempty"`           // 法定代表人
	PostCode                    string                                                   `json:"post_code,omitempty"`                      // 邮编
	TaxPayerID                  string                                                   `json:"tax_payer_id,omitempty"`                   // 纳税人识别号
	Confidential                bool                                                     `json:"confidential,omitempty"`                   // 是否保密
	SubTypeList                 []*GetCoreHRCompanyListRespItemSubType                   `json:"sub_type_list,omitempty"`                  // 主体类型, 枚举值可通过文档[飞书人事枚举常量](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)主体类型（company_sub_type）枚举定义部分获得
	BranchCompany               bool                                                     `json:"branch_company,omitempty"`                 // 是否为分公司
	PrimaryManager              []*GetCoreHRCompanyListRespItemPrimaryManager            `json:"primary_manager,omitempty"`                // 主要负责人
	CustomFields                []*GetCoreHRCompanyListRespItemCustomField               `json:"custom_fields,omitempty"`                  // 自定义字段
	Currency                    *GetCoreHRCompanyListRespItemCurrency                    `json:"currency,omitempty"`                       // 默认币种
	Phone                       *GetCoreHRCompanyListRespItemPhone                       `json:"phone,omitempty"`                          // 电话
	Fax                         *GetCoreHRCompanyListRespItemFax                         `json:"fax,omitempty"`                            // 传真
	RegisteredOfficeAddress     []*GetCoreHRCompanyListRespItemRegisteredOfficeAddres    `json:"registered_office_address,omitempty"`      // 完整注册地址
	OfficeAddress               []*GetCoreHRCompanyListRespItemOfficeAddres              `json:"office_address,omitempty"`                 // 完整办公地址
	RegisteredOfficeAddressInfo *GetCoreHRCompanyListRespItemRegisteredOfficeAddressInfo `json:"registered_office_address_info,omitempty"` // 注册地址详细信息
	OfficeAddressInfo           *GetCoreHRCompanyListRespItemOfficeAddressInfo           `json:"office_address_info,omitempty"`            // 办公地址详细信息
}

// GetCoreHRCompanyListRespItemCurrency ...
type GetCoreHRCompanyListRespItemCurrency struct {
	ID                 string                                              `json:"id,omitempty"`                    // 货币id
	CountryRegionID    string                                              `json:"country_region_id,omitempty"`     // 货币所属国家/地区id, 详细信息可通过[查询国家/地区信息](https://open.feishu.cn/document/server-docs/corehr-v1/basic-infomation/location_data/list)接口查询获得
	CurrencyName       []*GetCoreHRCompanyListRespItemCurrencyCurrencyName `json:"currency_name,omitempty"`         // 货币名称
	NumericCode        int64                                               `json:"numeric_code,omitempty"`          // 数字代码
	CurrencyAlpha3Code string                                              `json:"currency_alpha_3_code,omitempty"` // 三位字母代码
}

// GetCoreHRCompanyListRespItemCurrencyCurrencyName ...
type GetCoreHRCompanyListRespItemCurrencyCurrencyName struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// GetCoreHRCompanyListRespItemCustomField ...
type GetCoreHRCompanyListRespItemCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05")
}

// GetCoreHRCompanyListRespItemFax ...
type GetCoreHRCompanyListRespItemFax struct {
	AreaCode    *GetCoreHRCompanyListRespItemFaxAreaCode `json:"area_code,omitempty"`    // 区号
	PhoneNumber string                                   `json:"phone_number,omitempty"` // 号码
}

// GetCoreHRCompanyListRespItemFaxAreaCode ...
type GetCoreHRCompanyListRespItemFaxAreaCode struct {
	EnumName string                                            `json:"enum_name,omitempty"` // 枚举值
	Display  []*GetCoreHRCompanyListRespItemFaxAreaCodeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// GetCoreHRCompanyListRespItemFaxAreaCodeDisplay ...
type GetCoreHRCompanyListRespItemFaxAreaCodeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// GetCoreHRCompanyListRespItemHiberarchyCommon ...
type GetCoreHRCompanyListRespItemHiberarchyCommon struct {
	ParentID       string                                                     `json:"parent_id,omitempty"`       // 上级 ID
	Name           []*GetCoreHRCompanyListRespItemHiberarchyCommonName        `json:"name,omitempty"`            // 名称
	Type           *GetCoreHRCompanyListRespItemHiberarchyCommonType          `json:"type,omitempty"`            // 组织类型, 枚举值可通过文档[飞书人事枚举常量](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)组织类型（organization_type）枚举定义部分获得
	Active         bool                                                       `json:"active,omitempty"`          // 是否启用
	EffectiveTime  string                                                     `json:"effective_time,omitempty"`  // 生效时间
	ExpirationTime string                                                     `json:"expiration_time,omitempty"` // 失效时间
	Code           string                                                     `json:"code,omitempty"`            // 编码
	Description    []*GetCoreHRCompanyListRespItemHiberarchyCommonDescription `json:"description,omitempty"`     // 描述
	TreeOrder      string                                                     `json:"tree_order,omitempty"`      // 树形排序
	ListOrder      string                                                     `json:"list_order,omitempty"`      // 列表排序
	CustomFields   []*GetCoreHRCompanyListRespItemHiberarchyCommonCustomField `json:"custom_fields,omitempty"`   // 自定义字段
}

// GetCoreHRCompanyListRespItemHiberarchyCommonCustomField ...
type GetCoreHRCompanyListRespItemHiberarchyCommonCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05")
}

// GetCoreHRCompanyListRespItemHiberarchyCommonDescription ...
type GetCoreHRCompanyListRespItemHiberarchyCommonDescription struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// GetCoreHRCompanyListRespItemHiberarchyCommonName ...
type GetCoreHRCompanyListRespItemHiberarchyCommonName struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// GetCoreHRCompanyListRespItemHiberarchyCommonType ...
type GetCoreHRCompanyListRespItemHiberarchyCommonType struct {
	EnumName string                                                     `json:"enum_name,omitempty"` // 枚举值
	Display  []*GetCoreHRCompanyListRespItemHiberarchyCommonTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// GetCoreHRCompanyListRespItemHiberarchyCommonTypeDisplay ...
type GetCoreHRCompanyListRespItemHiberarchyCommonTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// GetCoreHRCompanyListRespItemIndustry ...
type GetCoreHRCompanyListRespItemIndustry struct {
	EnumName string                                         `json:"enum_name,omitempty"` // 枚举值
	Display  []*GetCoreHRCompanyListRespItemIndustryDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// GetCoreHRCompanyListRespItemIndustryDisplay ...
type GetCoreHRCompanyListRespItemIndustryDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// GetCoreHRCompanyListRespItemLegalRepresentative ...
type GetCoreHRCompanyListRespItemLegalRepresentative struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// GetCoreHRCompanyListRespItemOfficeAddres ...
type GetCoreHRCompanyListRespItemOfficeAddres struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// GetCoreHRCompanyListRespItemOfficeAddressInfo ...
type GetCoreHRCompanyListRespItemOfficeAddressInfo struct {
	PostalCode string `json:"postal_code,omitempty"` // 邮政编码
}

// GetCoreHRCompanyListRespItemPhone ...
type GetCoreHRCompanyListRespItemPhone struct {
	AreaCode    *GetCoreHRCompanyListRespItemPhoneAreaCode `json:"area_code,omitempty"`    // 区号
	PhoneNumber string                                     `json:"phone_number,omitempty"` // 号码
}

// GetCoreHRCompanyListRespItemPhoneAreaCode ...
type GetCoreHRCompanyListRespItemPhoneAreaCode struct {
	EnumName string                                              `json:"enum_name,omitempty"` // 枚举值
	Display  []*GetCoreHRCompanyListRespItemPhoneAreaCodeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// GetCoreHRCompanyListRespItemPhoneAreaCodeDisplay ...
type GetCoreHRCompanyListRespItemPhoneAreaCodeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// GetCoreHRCompanyListRespItemPrimaryManager ...
type GetCoreHRCompanyListRespItemPrimaryManager struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// GetCoreHRCompanyListRespItemRegisteredOfficeAddres ...
type GetCoreHRCompanyListRespItemRegisteredOfficeAddres struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// GetCoreHRCompanyListRespItemRegisteredOfficeAddressInfo ...
type GetCoreHRCompanyListRespItemRegisteredOfficeAddressInfo struct {
	PostalCode string `json:"postal_code,omitempty"` // 邮政编码
}

// GetCoreHRCompanyListRespItemSubType ...
type GetCoreHRCompanyListRespItemSubType struct {
	EnumName string                                        `json:"enum_name,omitempty"` // 枚举值
	Display  []*GetCoreHRCompanyListRespItemSubTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// GetCoreHRCompanyListRespItemSubTypeDisplay ...
type GetCoreHRCompanyListRespItemSubTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// GetCoreHRCompanyListRespItemType ...
type GetCoreHRCompanyListRespItemType struct {
	EnumName string                                     `json:"enum_name,omitempty"` // 枚举值
	Display  []*GetCoreHRCompanyListRespItemTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// GetCoreHRCompanyListRespItemTypeDisplay ...
type GetCoreHRCompanyListRespItemTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// getCoreHRCompanyListResp ...
type getCoreHRCompanyListResp struct {
	Code  int64                     `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                    `json:"msg,omitempty"`  // 错误描述
	Data  *GetCoreHRCompanyListResp `json:"data,omitempty"`
	Error *ErrorDetail              `json:"error,omitempty"`
}
