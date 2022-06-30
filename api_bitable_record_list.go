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

// GetBitableRecordList 该接口用于列出数据表中的现有记录, 单次最多列出 100 行记录, 支持分页获取。
//
// 该接口支持调用频率上限为 1000 次/分钟
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/list
func (r *BitableService) GetBitableRecordList(ctx context.Context, request *GetBitableRecordListReq, options ...MethodOptionFunc) (*GetBitableRecordListResp, *Response, error) {
	if r.cli.mock.mockBitableGetBitableRecordList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#GetBitableRecordList mock enable")
		return r.cli.mock.mockBitableGetBitableRecordList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "GetBitableRecordList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getBitableRecordListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBitableGetBitableRecordList mock BitableGetBitableRecordList method
func (r *Mock) MockBitableGetBitableRecordList(f func(ctx context.Context, request *GetBitableRecordListReq, options ...MethodOptionFunc) (*GetBitableRecordListResp, *Response, error)) {
	r.mockBitableGetBitableRecordList = f
}

// UnMockBitableGetBitableRecordList un-mock BitableGetBitableRecordList method
func (r *Mock) UnMockBitableGetBitableRecordList() {
	r.mockBitableGetBitableRecordList = nil
}

// GetBitableRecordListReq ...
type GetBitableRecordListReq struct {
	AppToken          string  `path:"app_token" json:"-"`            // bitable app token, 示例值: "bascnCMII2ORej2RItqpZZUNMIe"
	TableID           string  `path:"table_id" json:"-"`             // table id, 示例值: "tblxI2tWaxP5dG7p"
	ViewID            *string `query:"view_id" json:"-"`             // 视图 id, 注意: 如 filter 或 sort 有值, view_id 会被忽略, 示例值: "vewqhz51lk"
	Filter            *string `query:"filter" json:"-"`              // 筛选参数, 注意: 1.筛选记录的表达式不超过2000个字符, 2.不支持对“人员”以及“关联字段”的属性进行过滤筛选, 如人员的 OpenID, 3.仅支持字段在页面展示字符值进行筛选, 详细请参考[记录筛选开发指南](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/filter), 示例值: "示例表达式: AND(CurrentValue.[身高]>180, CurrentValue.[体重]>150)"
	Sort              *string `query:"sort" json:"-"`                // 排序参数, 注意: 1.表达式需要不超过1000字符, 2.不支持对带“公式”和“关联字段”的表的使用, 示例值: "["字段1 DESC", "字段2 ASC"], 注意: 使用引号将字段名称和顺序逆序连接起来。"
	FieldNames        *string `query:"field_names" json:"-"`         // 字段名称, 示例值: "["字段1"]"
	TextFieldAsArray  *bool   `query:"text_field_as_array" json:"-"` // 控制多行文本字段数据的返回格式, true 表示以数组形式返回, 注意: 1.多行文本中如果有超链接部分, 则会返回链接的 URL, 2.目前可以返回多行文本中 URL 类型为多维表格链接、飞书 doc、飞书 sheet的URL类型以及@人员的数据结构, 示例值: true
	UserIDType        *IDType `query:"user_id_type" json:"-"`        // 用户 ID 类型, 示例值: "open_id", 可选值有: `open_id`: 用户的 open id, `union_id`: 用户的 union id, `user_id`: 用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	DisplayFormulaRef *bool   `query:"display_formula_ref" json:"-"` // 控制公式、查找引用是否显示完整的原样返回结果, 示例值: true
	AutomaticFields   *bool   `query:"automatic_fields" json:"-"`    // 控制是否返回自动计算的字段, 例如 `created_by`/`created_time`/`last_modified_by`/`last_modified_time`, true 表示返回, 示例值: true
	PageToken         *string `query:"page_token" json:"-"`          // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "recn0hoyXL"
	PageSize          *int64  `query:"page_size" json:"-"`           // 分页大小, 示例值: 10, 最大值: `100`
}

// GetBitableRecordListResp ...
type GetBitableRecordListResp struct {
	HasMore   bool                            `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                          `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	Total     int64                           `json:"total,omitempty"`      // 总数
	Items     []*GetBitableRecordListRespItem `json:"items,omitempty"`      // 记录信息
}

// GetBitableRecordListRespItem ...
type GetBitableRecordListRespItem struct {
	RecordID         string                                      `json:"record_id,omitempty"`          // 记录 id
	CreatedBy        *GetBitableRecordListRespItemCreatedBy      `json:"created_by,omitempty"`         // 创建人
	CreatedTime      int64                                       `json:"created_time,omitempty"`       // 创建时间
	LastModifiedBy   *GetBitableRecordListRespItemLastModifiedBy `json:"last_modified_by,omitempty"`   // 修改人
	LastModifiedTime int64                                       `json:"last_modified_time,omitempty"` // 最近更新时间
	Fields           map[string]interface{}                      `json:"fields,omitempty"`             // 记录字段, 关于支持新增的字段类型, 请参考[接入指南](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/bitable/notification)
}

// GetBitableRecordListRespItemCreatedBy ...
type GetBitableRecordListRespItemCreatedBy struct {
	ID     string `json:"id,omitempty"`      // 人员Id
	Name   string `json:"name,omitempty"`    // 中文姓名
	EnName string `json:"en_name,omitempty"` // 英文姓名
	Email  string `json:"email,omitempty"`   // 邮箱
}

// GetBitableRecordListRespItemLastModifiedBy ...
type GetBitableRecordListRespItemLastModifiedBy struct {
	ID     string `json:"id,omitempty"`      // 人员Id
	Name   string `json:"name,omitempty"`    // 中文姓名
	EnName string `json:"en_name,omitempty"` // 英文姓名
	Email  string `json:"email,omitempty"`   // 邮箱
}

// getBitableRecordListResp ...
type getBitableRecordListResp struct {
	Code int64                     `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                    `json:"msg,omitempty"`  // 错误描述
	Data *GetBitableRecordListResp `json:"data,omitempty"`
}
