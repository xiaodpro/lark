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

// GetSheetDataValidationDropdown
//
// 该接口根据 spreadsheetToken 、range 查询range内的下拉列表设置信息；单次查询范围不超过5000行, 100列。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/datavalidation/query-datavalidation
func (r *DriveService) GetSheetDataValidationDropdown(ctx context.Context, request *GetSheetDataValidationDropdownReq, options ...MethodOptionFunc) (*GetSheetDataValidationDropdownResp, *Response, error) {
	if r.cli.mock.mockDriveGetSheetDataValidationDropdown != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#GetSheetDataValidationDropdown mock enable")
		return r.cli.mock.mockDriveGetSheetDataValidationDropdown(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "GetSheetDataValidationDropdown",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/dataValidation",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getSheetDataValidationDropdownResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveGetSheetDataValidationDropdown mock DriveGetSheetDataValidationDropdown method
func (r *Mock) MockDriveGetSheetDataValidationDropdown(f func(ctx context.Context, request *GetSheetDataValidationDropdownReq, options ...MethodOptionFunc) (*GetSheetDataValidationDropdownResp, *Response, error)) {
	r.mockDriveGetSheetDataValidationDropdown = f
}

// UnMockDriveGetSheetDataValidationDropdown un-mock DriveGetSheetDataValidationDropdown method
func (r *Mock) UnMockDriveGetSheetDataValidationDropdown() {
	r.mockDriveGetSheetDataValidationDropdown = nil
}

// GetSheetDataValidationDropdownReq ...
type GetSheetDataValidationDropdownReq struct {
	SpreadSheetToken   string `path:"spreadsheetToken" json:"-"`    // spreadsheet 的 token, 获取方式见 [在线表格开发指南](https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/overview)
	Range              string `query:"range" json:"-"`              // 查询范围, 包含 sheetId 与单元格范围两部分, 目前支持四种索引方式, 详见[在线表格开发指南](https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/overview)
	DataValidationType string `query:"dataValidationType" json:"-"` // 下拉列表填"list"
}

// GetSheetDataValidationDropdownResp ...
type GetSheetDataValidationDropdownResp struct {
	SpreadSheetToken string                                              `json:"spreadsheetToken,omitempty"` // spreadsheet的token
	SheetID          string                                              `json:"sheetId,omitempty"`          // 工作表 sheet 的 id
	Revision         int64                                               `json:"revision,omitempty"`         // 版本号
	DataValidations  []*GetSheetDataValidationDropdownRespDataValidation `json:"dataValidations,omitempty"`  // 下拉列表数组, 不存在时为空
}

// GetSheetDataValidationDropdownRespDataValidation ...
type GetSheetDataValidationDropdownRespDataValidation struct {
	DataValidationID   int64                                                    `json:"dataValidationId,omitempty"`   // 唯一标示id
	DataValidationType string                                                   `json:"dataValidationType,omitempty"` // 下拉列表为"list"
	ConditionValues    []string                                                 `json:"conditionValues,omitempty"`    // 下拉列表选项值
	Options            *GetSheetDataValidationDropdownRespDataValidationOptions `json:"options,omitempty"`            // 可选属性
}

// GetSheetDataValidationDropdownRespDataValidationOptions ...
type GetSheetDataValidationDropdownRespDataValidationOptions struct {
	MultipleValues     *bool             `json:"multipleValues,omitempty"`     // 单选填false, 多选填true
	HighlightValidData *bool             `json:"highlightValidData,omitempty"` // 是否设置颜色和胶囊样式
	ColorValueMap      map[string]string `json:"colorValueMap,omitempty"`      // 当highlightValidData为true时, colorValueMap的key与conditionValues中的值一一对应, value为对应的颜色参数。
}

// getSheetDataValidationDropdownResp ...
type getSheetDataValidationDropdownResp struct {
	Code int64                               `json:"code,omitempty"` // 状态码, 0代表成功
	Msg  *string                             `json:"msg,omitempty"`  // 状态信息
	Data *GetSheetDataValidationDropdownResp `json:"data,omitempty"`
}
