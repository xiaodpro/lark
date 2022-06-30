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

// UpdateSheetFilterViewCondition 更新筛选视图范围的某列的筛选条件, condition id 即为列的字母号。
//
// 筛选条件参数可参考 [筛选视图的筛选条件指南](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/filter-view-condition-user-guide)
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/update
func (r *DriveService) UpdateSheetFilterViewCondition(ctx context.Context, request *UpdateSheetFilterViewConditionReq, options ...MethodOptionFunc) (*UpdateSheetFilterViewConditionResp, *Response, error) {
	if r.cli.mock.mockDriveUpdateSheetFilterViewCondition != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#UpdateSheetFilterViewCondition mock enable")
		return r.cli.mock.mockDriveUpdateSheetFilterViewCondition(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "UpdateSheetFilterViewCondition",
		Method:                "PUT",
		URL:                   r.cli.openBaseURL + "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id/conditions/:condition_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateSheetFilterViewConditionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveUpdateSheetFilterViewCondition mock DriveUpdateSheetFilterViewCondition method
func (r *Mock) MockDriveUpdateSheetFilterViewCondition(f func(ctx context.Context, request *UpdateSheetFilterViewConditionReq, options ...MethodOptionFunc) (*UpdateSheetFilterViewConditionResp, *Response, error)) {
	r.mockDriveUpdateSheetFilterViewCondition = f
}

// UnMockDriveUpdateSheetFilterViewCondition un-mock DriveUpdateSheetFilterViewCondition method
func (r *Mock) UnMockDriveUpdateSheetFilterViewCondition() {
	r.mockDriveUpdateSheetFilterViewCondition = nil
}

// UpdateSheetFilterViewConditionReq ...
type UpdateSheetFilterViewConditionReq struct {
	SpreadSheetToken string   `path:"spreadsheet_token" json:"-"` // 表格 token, 示例值: "shtcnmBA*yGehy8"
	SheetID          string   `path:"sheet_id" json:"-"`          // 子表 id, 示例值: "0b**12"
	FilterViewID     string   `path:"filter_view_id" json:"-"`    // 筛选视图 id, 示例值: "pH9hbVcCXA"
	ConditionID      string   `path:"condition_id" json:"-"`      // 列字母号, 示例值: "E"
	FilterType       *string  `json:"filter_type,omitempty"`      // 筛选类型, 示例值: "number"
	CompareType      *string  `json:"compare_type,omitempty"`     // 比较类型, 示例值: "less"
	Expected         []string `json:"expected,omitempty"`         // 筛选参数, 示例值: 6
}

// UpdateSheetFilterViewConditionResp ...
type UpdateSheetFilterViewConditionResp struct {
	Condition *UpdateSheetFilterViewConditionRespCondition `json:"condition,omitempty"` // 更新后的筛选条件
}

// UpdateSheetFilterViewConditionRespCondition ...
type UpdateSheetFilterViewConditionRespCondition struct {
	ConditionID string   `json:"condition_id,omitempty"` // 设置筛选条件的列, 使用字母号
	FilterType  string   `json:"filter_type,omitempty"`  // 筛选类型
	CompareType string   `json:"compare_type,omitempty"` // 比较类型
	Expected    []string `json:"expected,omitempty"`     // 筛选参数
}

// updateSheetFilterViewConditionResp ...
type updateSheetFilterViewConditionResp struct {
	Code int64                               `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                              `json:"msg,omitempty"`  // 错误描述
	Data *UpdateSheetFilterViewConditionResp `json:"data,omitempty"`
}
