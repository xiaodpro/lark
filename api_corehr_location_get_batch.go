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

// BatchGetCoreHRLocation 通过地点 ID 批量获取地点信息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/location/batch_get
func (r *CoreHRService) BatchGetCoreHRLocation(ctx context.Context, request *BatchGetCoreHRLocationReq, options ...MethodOptionFunc) (*BatchGetCoreHRLocationResp, *Response, error) {
	if r.cli.mock.mockCoreHRBatchGetCoreHRLocation != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#BatchGetCoreHRLocation mock enable")
		return r.cli.mock.mockCoreHRBatchGetCoreHRLocation(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "BatchGetCoreHRLocation",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v2/locations/batch_get",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(batchGetCoreHRLocationResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRBatchGetCoreHRLocation mock CoreHRBatchGetCoreHRLocation method
func (r *Mock) MockCoreHRBatchGetCoreHRLocation(f func(ctx context.Context, request *BatchGetCoreHRLocationReq, options ...MethodOptionFunc) (*BatchGetCoreHRLocationResp, *Response, error)) {
	r.mockCoreHRBatchGetCoreHRLocation = f
}

// UnMockCoreHRBatchGetCoreHRLocation un-mock CoreHRBatchGetCoreHRLocation method
func (r *Mock) UnMockCoreHRBatchGetCoreHRLocation() {
	r.mockCoreHRBatchGetCoreHRLocation = nil
}

// BatchGetCoreHRLocationReq ...
type BatchGetCoreHRLocationReq struct {
	LocationIDs []string `json:"location_ids,omitempty"` // 地点 ID 列表, 示例值: ["1215"], 长度范围: `1` ～ `100`
}

// BatchGetCoreHRLocationResp ...
type BatchGetCoreHRLocationResp struct {
	Items []*BatchGetCoreHRLocationRespItem `json:"items,omitempty"` // 查询的地点信息
}

// BatchGetCoreHRLocationRespItem ...
type BatchGetCoreHRLocationRespItem struct {
	LocationID         string                                          `json:"location_id,omitempty"`           // 地点 ID
	HiberarchyCommon   *BatchGetCoreHRLocationRespItemHiberarchyCommon `json:"hiberarchy_common,omitempty"`     // 地点基本信息
	LocationUsageList  []*BatchGetCoreHRLocationRespItemLocationUsage  `json:"location_usage_list,omitempty"`   // 地点用途
	Address            []*BatchGetCoreHRLocationRespItemAddres         `json:"address,omitempty"`               // 地址
	WorkingHoursTypeID string                                          `json:"working_hours_type_id,omitempty"` // 工时制度
	EffectiveTime      string                                          `json:"effective_time,omitempty"`        // 生效时间
	ExpirationTime     string                                          `json:"expiration_time,omitempty"`       // 失效时间
	CustomFields       []*BatchGetCoreHRLocationRespItemCustomField    `json:"custom_fields,omitempty"`         // 自定义字段
	Locale             *BatchGetCoreHRLocationRespItemLocale           `json:"locale,omitempty"`                // 区域设置
	TimeZoneID         string                                          `json:"time_zone_id,omitempty"`          // 时区
	DisplayLanguageID  string                                          `json:"display_language_id,omitempty"`   // 默认显示语言
}

// BatchGetCoreHRLocationRespItemAddres ...
type BatchGetCoreHRLocationRespItemAddres struct {
	FullAddressLocalScript   string                                             `json:"full_address_local_script,omitempty"`   // 完整地址（本地文字）
	FullAddressWesternScript string                                             `json:"full_address_western_script,omitempty"` // 完整地址（西方文字）
	AddressID                string                                             `json:"address_id,omitempty"`                  // 地址 ID
	CountryRegionID          string                                             `json:"country_region_id,omitempty"`           // 国家 / 地区
	RegionID                 string                                             `json:"region_id,omitempty"`                   // 主要行政区
	LocalAddressLine1        string                                             `json:"local_address_line1,omitempty"`         // 地址行 1（非拉丁语系的本地文字）
	LocalAddressLine2        string                                             `json:"local_address_line2,omitempty"`         // 地址行 2（非拉丁语系的本地文字）
	LocalAddressLine3        string                                             `json:"local_address_line3,omitempty"`         // 地址行 3（非拉丁语系的本地文字）
	LocalAddressLine4        string                                             `json:"local_address_line4,omitempty"`         // 地址行 4（非拉丁语系的本地文字）
	LocalAddressLine5        string                                             `json:"local_address_line5,omitempty"`         // 地址行 5（非拉丁语系的本地文字）
	LocalAddressLine6        string                                             `json:"local_address_line6,omitempty"`         // 地址行 6（非拉丁语系的本地文字）
	LocalAddressLine7        string                                             `json:"local_address_line7,omitempty"`         // 地址行 7（非拉丁语系的本地文字）
	LocalAddressLine8        string                                             `json:"local_address_line8,omitempty"`         // 地址行 8（非拉丁语系的本地文字）
	LocalAddressLine9        string                                             `json:"local_address_line9,omitempty"`         // 地址行 9（非拉丁语系的本地文字）
	PostalCode               string                                             `json:"postal_code,omitempty"`                 // 邮政编码
	AddressTypeList          []*BatchGetCoreHRLocationRespItemAddresAddressType `json:"address_type_list,omitempty"`           // 地址类型
	IsPrimary                bool                                               `json:"is_primary,omitempty"`                  // 主要地址
	IsPublic                 bool                                               `json:"is_public,omitempty"`                   // 公开地址
	CustomFields             []*BatchGetCoreHRLocationRespItemAddresCustomField `json:"custom_fields,omitempty"`               // 自定义字段
}

// BatchGetCoreHRLocationRespItemAddresAddressType ...
type BatchGetCoreHRLocationRespItemAddresAddressType struct {
	EnumName string                                                    `json:"enum_name,omitempty"` // 枚举值
	Display  []*BatchGetCoreHRLocationRespItemAddresAddressTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// BatchGetCoreHRLocationRespItemAddresAddressTypeDisplay ...
type BatchGetCoreHRLocationRespItemAddresAddressTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// BatchGetCoreHRLocationRespItemAddresCustomField ...
type BatchGetCoreHRLocationRespItemAddresCustomField struct {
	CustomApiName string                                               `json:"custom_api_name,omitempty"` // 自定义字段 apiname, 即自定义字段的唯一标识
	Name          *BatchGetCoreHRLocationRespItemAddresCustomFieldName `json:"name,omitempty"`            // 自定义字段名称
	Type          int64                                                `json:"type,omitempty"`            // 自定义字段类型
	Value         string                                               `json:"value,omitempty"`           // 字段值, 是 json 转义后的字符串, 根据元数据定义不同, 字段格式不同（如 123, 123.23, "true", ["id1", "id2"], "2006-01-02 15:04:05"）
}

// BatchGetCoreHRLocationRespItemAddresCustomFieldName ...
type BatchGetCoreHRLocationRespItemAddresCustomFieldName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// BatchGetCoreHRLocationRespItemCustomField ...
type BatchGetCoreHRLocationRespItemCustomField struct {
	CustomApiName string                                         `json:"custom_api_name,omitempty"` // 自定义字段 apiname, 即自定义字段的唯一标识
	Name          *BatchGetCoreHRLocationRespItemCustomFieldName `json:"name,omitempty"`            // 自定义字段名称
	Type          int64                                          `json:"type,omitempty"`            // 自定义字段类型
	Value         string                                         `json:"value,omitempty"`           // 字段值, 是 json 转义后的字符串, 根据元数据定义不同, 字段格式不同（如 123, 123.23, "true", ["id1", "id2"], "2006-01-02 15:04:05"）
}

// BatchGetCoreHRLocationRespItemCustomFieldName ...
type BatchGetCoreHRLocationRespItemCustomFieldName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// BatchGetCoreHRLocationRespItemHiberarchyCommon ...
type BatchGetCoreHRLocationRespItemHiberarchyCommon struct {
	ParentID       string                                                       `json:"parent_id,omitempty"`       // 上级组织
	Name           []*BatchGetCoreHRLocationRespItemHiberarchyCommonName        `json:"name,omitempty"`            // 名称
	Type           *BatchGetCoreHRLocationRespItemHiberarchyCommonType          `json:"type,omitempty"`            // 组织类型
	Active         bool                                                         `json:"active,omitempty"`          // 启用
	EffectiveTime  string                                                       `json:"effective_time,omitempty"`  // 生效时间
	ExpirationTime string                                                       `json:"expiration_time,omitempty"` // 失效时间
	Code           string                                                       `json:"code,omitempty"`            // 编码
	Description    []*BatchGetCoreHRLocationRespItemHiberarchyCommonDescription `json:"description,omitempty"`     // 描述
	TreeOrder      string                                                       `json:"tree_order,omitempty"`      // 树形排序, 代表同层级的部门排序序号
	ListOrder      string                                                       `json:"list_order,omitempty"`      // 列表排序, 代表所有部门的混排序号
	CustomFields   []*BatchGetCoreHRLocationRespItemHiberarchyCommonCustomField `json:"custom_fields,omitempty"`   // 自定义字段
}

// BatchGetCoreHRLocationRespItemHiberarchyCommonCustomField ...
type BatchGetCoreHRLocationRespItemHiberarchyCommonCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(123, 123.23, true, [\"id1\", \"id2\], 2006-01-02 15:04:05])
}

// BatchGetCoreHRLocationRespItemHiberarchyCommonDescription ...
type BatchGetCoreHRLocationRespItemHiberarchyCommonDescription struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// BatchGetCoreHRLocationRespItemHiberarchyCommonName ...
type BatchGetCoreHRLocationRespItemHiberarchyCommonName struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// BatchGetCoreHRLocationRespItemHiberarchyCommonType ...
type BatchGetCoreHRLocationRespItemHiberarchyCommonType struct {
	EnumName string                                                       `json:"enum_name,omitempty"` // 枚举值
	Display  []*BatchGetCoreHRLocationRespItemHiberarchyCommonTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// BatchGetCoreHRLocationRespItemHiberarchyCommonTypeDisplay ...
type BatchGetCoreHRLocationRespItemHiberarchyCommonTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// BatchGetCoreHRLocationRespItemLocale ...
type BatchGetCoreHRLocationRespItemLocale struct {
	EnumName string                                         `json:"enum_name,omitempty"` // 枚举值
	Display  []*BatchGetCoreHRLocationRespItemLocaleDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// BatchGetCoreHRLocationRespItemLocaleDisplay ...
type BatchGetCoreHRLocationRespItemLocaleDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// BatchGetCoreHRLocationRespItemLocationUsage ...
type BatchGetCoreHRLocationRespItemLocationUsage struct {
	EnumName string                                                `json:"enum_name,omitempty"` // 枚举值
	Display  []*BatchGetCoreHRLocationRespItemLocationUsageDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// BatchGetCoreHRLocationRespItemLocationUsageDisplay ...
type BatchGetCoreHRLocationRespItemLocationUsageDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// batchGetCoreHRLocationResp ...
type batchGetCoreHRLocationResp struct {
	Code  int64                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                      `json:"msg,omitempty"`  // 错误描述
	Data  *BatchGetCoreHRLocationResp `json:"data,omitempty"`
	Error *ErrorDetail                `json:"error,omitempty"`
}