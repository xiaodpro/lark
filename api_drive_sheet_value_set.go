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

// SetSheetValue 该接口用于根据 spreadsheetToken 和 range 向单个范围写入数据, 若范围内有数据, 将被更新覆盖；单次写入不超过5000行, 100列, 每个格子不超过5万字符。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uAjMzUjLwIzM14CMyMTN
// new doc: https://open.feishu.cn/document/server-docs/docs/sheets-v3/data-operation/write-data-to-a-single-range
func (r *DriveService) SetSheetValue(ctx context.Context, request *SetSheetValueReq, options ...MethodOptionFunc) (*SetSheetValueResp, *Response, error) {
	if r.cli.mock.mockDriveSetSheetValue != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#SetSheetValue mock enable")
		return r.cli.mock.mockDriveSetSheetValue(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "SetSheetValue",
		Method:                "PUT",
		URL:                   r.cli.openBaseURL + "/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(setSheetValueResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveSetSheetValue mock DriveSetSheetValue method
func (r *Mock) MockDriveSetSheetValue(f func(ctx context.Context, request *SetSheetValueReq, options ...MethodOptionFunc) (*SetSheetValueResp, *Response, error)) {
	r.mockDriveSetSheetValue = f
}

// UnMockDriveSetSheetValue un-mock DriveSetSheetValue method
func (r *Mock) UnMockDriveSetSheetValue() {
	r.mockDriveSetSheetValue = nil
}

// SetSheetValueReq ...
type SetSheetValueReq struct {
	SpreadSheetToken string                      `path:"spreadsheetToken" json:"-"` // spreadsheet的token, 获取方式见[在线表格开发指南](https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/overview)
	ValueRange       *SetSheetValueReqValueRange `json:"valueRange,omitempty"`      // 值与范围
}

// SetSheetValueReqValueRange ...
type SetSheetValueReqValueRange struct {
	Range  string           `json:"range,omitempty"`  // 更新范围, 包含 sheetId 与单元格范围两部分, 目前支持四种索引方式, 详见[在线表格开发指南](https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/overview), range所表示的范围需要大于等于values占用的范围
	Values [][]SheetContent `json:"values,omitempty"` // 需要写入的值, 如要写入公式、超链接、email、@人等, 可详看附录[sheet 支持写入数据类型](https://open.feishu.cn/document/ukTMukTMukTM/ugjN1UjL4YTN14CO2UTN)
}

// SetSheetValueResp ...
type SetSheetValueResp struct {
	SpreadSheetToken string `json:"spreadsheetToken,omitempty"` // spreadsheet 的 token
	UpdatedRange     string `json:"updatedRange,omitempty"`     // 写入的范围
	UpdatedRows      int64  `json:"updatedRows,omitempty"`      // 写入的行数
	UpdatedColumns   int64  `json:"updatedColumns,omitempty"`   // 写入的列数
	UpdatedCells     int64  `json:"updatedCells,omitempty"`     // 写入的单元格总数
	Revision         int64  `json:"revision,omitempty"`         // sheet 的版本号
}

// setSheetValueResp ...
type setSheetValueResp struct {
	Code int64              `json:"code,omitempty"`
	Msg  string             `json:"msg,omitempty"`
	Data *SetSheetValueResp `json:"data,omitempty"`
}
