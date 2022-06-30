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

// DeleteSheetFilterView 删除指定 id 对应的筛选视图。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view/delete
func (r *DriveService) DeleteSheetFilterView(ctx context.Context, request *DeleteSheetFilterViewReq, options ...MethodOptionFunc) (*DeleteSheetFilterViewResp, *Response, error) {
	if r.cli.mock.mockDriveDeleteSheetFilterView != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#DeleteSheetFilterView mock enable")
		return r.cli.mock.mockDriveDeleteSheetFilterView(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "DeleteSheetFilterView",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(deleteSheetFilterViewResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveDeleteSheetFilterView mock DriveDeleteSheetFilterView method
func (r *Mock) MockDriveDeleteSheetFilterView(f func(ctx context.Context, request *DeleteSheetFilterViewReq, options ...MethodOptionFunc) (*DeleteSheetFilterViewResp, *Response, error)) {
	r.mockDriveDeleteSheetFilterView = f
}

// UnMockDriveDeleteSheetFilterView un-mock DriveDeleteSheetFilterView method
func (r *Mock) UnMockDriveDeleteSheetFilterView() {
	r.mockDriveDeleteSheetFilterView = nil
}

// DeleteSheetFilterViewReq ...
type DeleteSheetFilterViewReq struct {
	SpreadSheetToken string `path:"spreadsheet_token" json:"-"` // 表格 token, 示例值: "shtcnmBA*yGehy8"
	SheetID          string `path:"sheet_id" json:"-"`          // 子表 id, 示例值: "0b**12"
	FilterViewID     string `path:"filter_view_id" json:"-"`    // 筛选视图 id, 示例值: "pH9hbVcCXA"
}

// DeleteSheetFilterViewResp ...
type DeleteSheetFilterViewResp struct {
}

// deleteSheetFilterViewResp ...
type deleteSheetFilterViewResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *DeleteSheetFilterViewResp `json:"data,omitempty"`
}
