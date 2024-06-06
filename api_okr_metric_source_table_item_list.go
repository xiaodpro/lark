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

// GetOKRMetricSourceTableItemList 获取指定指标表下的所有指标项（仅限 OKR 企业版使用）。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/okr-v1/metric_source-table-item/list
//
// Deprecated
func (r *OKRService) GetOKRMetricSourceTableItemList(ctx context.Context, request *GetOKRMetricSourceTableItemListReq, options ...MethodOptionFunc) (*GetOKRMetricSourceTableItemListResp, *Response, error) {
	if r.cli.mock.mockOKRGetOKRMetricSourceTableItemList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] OKR#GetOKRMetricSourceTableItemList mock enable")
		return r.cli.mock.mockOKRGetOKRMetricSourceTableItemList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "OKR",
		API:                   "GetOKRMetricSourceTableItemList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/okr/v1/metric_sources/:metric_source_id/tables/:metric_table_id/items",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getOKRMetricSourceTableItemListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockOKRGetOKRMetricSourceTableItemList mock OKRGetOKRMetricSourceTableItemList method
func (r *Mock) MockOKRGetOKRMetricSourceTableItemList(f func(ctx context.Context, request *GetOKRMetricSourceTableItemListReq, options ...MethodOptionFunc) (*GetOKRMetricSourceTableItemListResp, *Response, error)) {
	r.mockOKRGetOKRMetricSourceTableItemList = f
}

// UnMockOKRGetOKRMetricSourceTableItemList un-mock OKRGetOKRMetricSourceTableItemList method
func (r *Mock) UnMockOKRGetOKRMetricSourceTableItemList() {
	r.mockOKRGetOKRMetricSourceTableItemList = nil
}

// GetOKRMetricSourceTableItemListReq ...
type GetOKRMetricSourceTableItemListReq struct {
	MetricSourceID string  `path:"metric_source_id" json:"-"` // okr指标库id, 示例值: "7041857032248410131"
	MetricTableID  string  `path:"metric_table_id" json:"-"`  // okr指标表id, 示例值: "7041857032248410131"
	UserIDType     *IDType `query:"user_id_type" json:"-"`    // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	PageToken      *string `query:"page_token" json:"-"`      // 页码标识, 获取第一页传空, 每次查询会返回下一页的page_token, 示例值: "6969864184272078374"
	PageSize       *int64  `query:"page_size" json:"-"`       // 每页获取记录数, 示例值: "10"
}

// GetOKRMetricSourceTableItemListResp ...
type GetOKRMetricSourceTableItemListResp struct {
	Total     int64                                      `json:"total,omitempty"`      // 符合条件的记录总数
	HasMore   bool                                       `json:"has_more,omitempty"`   // 是否有下一页
	PageToken string                                     `json:"page_token,omitempty"` // 下一页页码
	Items     []*GetOKRMetricSourceTableItemListRespItem `json:"items,omitempty"`      // 指标项列表
}

// GetOKRMetricSourceTableItemListRespItem ...
type GetOKRMetricSourceTableItemListRespItem struct {
	MetricItemID       string                                             `json:"metric_item_id,omitempty"`       // 指标项id
	UserID             string                                             `json:"user_id,omitempty"`              // 指标承接人员id
	PeriodID           string                                             `json:"period_id,omitempty"`            // 指标的okr周期
	MetricUnit         *GetOKRMetricSourceTableItemListRespItemMetricUnit `json:"metric_unit,omitempty"`          // 指标单位
	MetricInitialValue float64                                            `json:"metric_initial_value,omitempty"` // 指标起始值
	MetricTargetValue  float64                                            `json:"metric_target_value,omitempty"`  // 指标目标值
	MetricCurrentValue float64                                            `json:"metric_current_value,omitempty"` // 指标进度值
	SupportedUserID    string                                             `json:"supported_user_id,omitempty"`    // 指标支撑的上级人员id
	KrID               string                                             `json:"kr_id,omitempty"`                // 指标关联的kr
	UpdatedAt          string                                             `json:"updated_at,omitempty"`           // 更新时间
	UpdatedBy          string                                             `json:"updated_by,omitempty"`           // 更新人
}

// GetOKRMetricSourceTableItemListRespItemMetricUnit ...
type GetOKRMetricSourceTableItemListRespItemMetricUnit struct {
	ZhCn string `json:"zh_cn,omitempty"` // 指标单位中文
	EnUs string `json:"en_us,omitempty"` // 指标单位英文
	JaJp string `json:"ja_jp,omitempty"` // 指标单位日文
}

// getOKRMetricSourceTableItemListResp ...
type getOKRMetricSourceTableItemListResp struct {
	Code  int64                                `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                               `json:"msg,omitempty"`  // 错误描述
	Data  *GetOKRMetricSourceTableItemListResp `json:"data,omitempty"`
	Error *ErrorDetail                         `json:"error,omitempty"`
}
