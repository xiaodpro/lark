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

// CreateSheet 使用该接口可以在指定的目录下创建在线表格。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet/create
func (r *DriveService) CreateSheet(ctx context.Context, request *CreateSheetReq, options ...MethodOptionFunc) (*CreateSheetResp, *Response, error) {
	if r.cli.mock.mockDriveCreateSheet != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#CreateSheet mock enable")
		return r.cli.mock.mockDriveCreateSheet(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "CreateSheet",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/sheets/v3/spreadsheets",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createSheetResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveCreateSheet mock DriveCreateSheet method
func (r *Mock) MockDriveCreateSheet(f func(ctx context.Context, request *CreateSheetReq, options ...MethodOptionFunc) (*CreateSheetResp, *Response, error)) {
	r.mockDriveCreateSheet = f
}

// UnMockDriveCreateSheet un-mock DriveCreateSheet method
func (r *Mock) UnMockDriveCreateSheet() {
	r.mockDriveCreateSheet = nil
}

// CreateSheetReq ...
type CreateSheetReq struct {
	Title       *string `json:"title,omitempty"`        // 表格标题, 示例值: "title"
	FolderToken *string `json:"folder_token,omitempty"` // 文件夹token, 获取方式见[概述](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/files/guide/introduction), 示例值: "fldcnMsNb*hIW9IjG1LVswg"
}

// CreateSheetResp ...
type CreateSheetResp struct {
	Spreadsheet *CreateSheetRespSpreadsheet `json:"spreadsheet,omitempty"` // 表格
}

// CreateSheetRespSpreadsheet ...
type CreateSheetRespSpreadsheet struct {
	Title            string `json:"title,omitempty"`             // 表格标题
	FolderToken      string `json:"folder_token,omitempty"`      // 文件夹token, 获取方式见[概述](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/files/guide/introduction)
	URL              string `json:"url,omitempty"`               // 文档url
	SpreadSheetToken string `json:"spreadsheet_token,omitempty"` // 表格token
}

// createSheetResp ...
type createSheetResp struct {
	Code int64            `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string           `json:"msg,omitempty"`  // 错误描述
	Data *CreateSheetResp `json:"data,omitempty"`
}
