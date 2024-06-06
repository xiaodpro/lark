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

// UpdateSheetDataValidationDropdown 该接口根据 spreadsheetToken 、sheetId、dataValidationId 更新下拉列表的属性。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/datavalidation/update-datavalidation
// new doc: https://open.feishu.cn/document/server-docs/docs/sheets-v3/datavalidation/update-datavalidation
func (r *DriveService) UpdateSheetDataValidationDropdown(ctx context.Context, request *UpdateSheetDataValidationDropdownReq, options ...MethodOptionFunc) (*UpdateSheetDataValidationDropdownResp, *Response, error) {
	if r.cli.mock.mockDriveUpdateSheetDataValidationDropdown != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#UpdateSheetDataValidationDropdown mock enable")
		return r.cli.mock.mockDriveUpdateSheetDataValidationDropdown(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "UpdateSheetDataValidationDropdown",
		Method:                "PUT",
		URL:                   r.cli.openBaseURL + "/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/dataValidation/:sheetId/:dataValidationId",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateSheetDataValidationDropdownResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveUpdateSheetDataValidationDropdown mock DriveUpdateSheetDataValidationDropdown method
func (r *Mock) MockDriveUpdateSheetDataValidationDropdown(f func(ctx context.Context, request *UpdateSheetDataValidationDropdownReq, options ...MethodOptionFunc) (*UpdateSheetDataValidationDropdownResp, *Response, error)) {
	r.mockDriveUpdateSheetDataValidationDropdown = f
}

// UnMockDriveUpdateSheetDataValidationDropdown un-mock DriveUpdateSheetDataValidationDropdown method
func (r *Mock) UnMockDriveUpdateSheetDataValidationDropdown() {
	r.mockDriveUpdateSheetDataValidationDropdown = nil
}

// UpdateSheetDataValidationDropdownReq ...
type UpdateSheetDataValidationDropdownReq struct {
	SpreadSheetToken   string                                              `path:"spreadsheetToken" json:"-"`    // spreadsheet 的 token, 获取方式见[在线表格开发指南](https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/overview)
	SheetID            string                                              `path:"sheetId" json:"-"`             // 子sheet唯一识别参数
	DataValidationID   int64                                               `path:"dataValidationId" json:"-"`    // sheet中下拉列表的唯一标示id
	DataValidationType string                                              `json:"dataValidationType,omitempty"` // 下拉列表填"list"
	DataValidation     *UpdateSheetDataValidationDropdownReqDataValidation `json:"dataValidation,omitempty"`     // 下拉列表规则属性
}

// UpdateSheetDataValidationDropdownReqDataValidation ...
type UpdateSheetDataValidationDropdownReqDataValidation struct {
	ConditionValues []string                                                   `json:"conditionValues,omitempty"` // 下拉列表选项值, 需为字符串, 不能包含", ", 选项值最长100字符, 选项个数最多500个
	Options         *UpdateSheetDataValidationDropdownReqDataValidationOptions `json:"options,omitempty"`         // 可选属性
}

// UpdateSheetDataValidationDropdownReqDataValidationOptions ...
type UpdateSheetDataValidationDropdownReqDataValidationOptions struct {
	MultipleValues     *bool    `json:"multipleValues,omitempty"`     // 单选填false, 多选填true, 不填默认为false
	HighlightValidData *bool    `json:"highlightValidData,omitempty"` // 是否设置颜色和胶囊样式, 不填默认为false
	Colors             []string `json:"colors,omitempty"`             // 当highlightValidData为true时, color需填颜色, 与conditionValues中的值一一对应。需是RGB16进制格式, 如"#fffd00"
}

// UpdateSheetDataValidationDropdownResp ...
type UpdateSheetDataValidationDropdownResp struct {
	SpreadSheetToken string                                               `json:" spreadsheetToken,omitempty"` // spreadsheet的token
	SheetID          string                                               `json:" sheetId,omitempty"`          // 工作表 sheet 的 id
	DataValidation   *UpdateSheetDataValidationDropdownRespDataValidation `json:"dataValidation,omitempty"`
}

// UpdateSheetDataValidationDropdownRespDataValidation ...
type UpdateSheetDataValidationDropdownRespDataValidation struct {
	DataValidationID   int64                                                       `json:"dataValidationId,omitempty"`   // 唯一标示id
	DataValidationType string                                                      `json:"dataValidationType,omitempty"` // 下拉列表为"list"
	ConditionValues    []string                                                    `json:"conditionValues,omitempty"`    // 下拉列表选项值
	Options            *UpdateSheetDataValidationDropdownRespDataValidationOptions `json:"options,omitempty"`            // 可选属性
}

// UpdateSheetDataValidationDropdownRespDataValidationOptions ...
type UpdateSheetDataValidationDropdownRespDataValidationOptions struct {
	MultipleValues     *bool             `json:"multipleValues,omitempty"`     // 单选填false, 多选填true
	HighlightValidData *bool             `json:"highlightValidData,omitempty"` // 是否设置颜色和胶囊样式
	ColorValueMap      map[string]string `json:"colorValueMap,omitempty"`      // 当highlightValidData为true时, colorValueMap的key与conditionValues中的值一一对应, value为对应的颜色参数。
}

// updateSheetDataValidationDropdownResp ...
type updateSheetDataValidationDropdownResp struct {
	Code  int64                                  `json:"code,omitempty"` // 状态码, 0代表成功
	Msg   *string                                `json:"msg,omitempty"`  // 状态信息
	Data  *UpdateSheetDataValidationDropdownResp `json:"data,omitempty"`
	Error *ErrorDetail                           `json:"error,omitempty"`
}
