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

// DeleteSheetConditionFormat
//
// 该接口用于删除已有的条件格式, 单次最多支持删除10个条件格式, 每个条件格式的删除会返回成功或者失败, 失败的情况包括各种参数的校验。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/conditionformat/condition-format-delete
func (r *DriveService) DeleteSheetConditionFormat(ctx context.Context, request *DeleteSheetConditionFormatReq, options ...MethodOptionFunc) (*DeleteSheetConditionFormatResp, *Response, error) {
	if r.cli.mock.mockDriveDeleteSheetConditionFormat != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#DeleteSheetConditionFormat mock enable")
		return r.cli.mock.mockDriveDeleteSheetConditionFormat(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "DeleteSheetConditionFormat",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/condition_formats/batch_delete",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(deleteSheetConditionFormatResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveDeleteSheetConditionFormat mock DriveDeleteSheetConditionFormat method
func (r *Mock) MockDriveDeleteSheetConditionFormat(f func(ctx context.Context, request *DeleteSheetConditionFormatReq, options ...MethodOptionFunc) (*DeleteSheetConditionFormatResp, *Response, error)) {
	r.mockDriveDeleteSheetConditionFormat = f
}

// UnMockDriveDeleteSheetConditionFormat un-mock DriveDeleteSheetConditionFormat method
func (r *Mock) UnMockDriveDeleteSheetConditionFormat() {
	r.mockDriveDeleteSheetConditionFormat = nil
}

// DeleteSheetConditionFormatReq ...
type DeleteSheetConditionFormatReq struct {
	SpreadSheetToken string                                   `path:"spreadsheetToken" json:"-"` // sheet 的 token, 获取方式见 [在线表格开发指南](https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/overview)
	SheetCfIDs       *DeleteSheetConditionFormatReqSheetCfIDs `json:"sheet_cf_ids,omitempty"`    // 表格条件格式id
}

// DeleteSheetConditionFormatReqSheetCfIDs ...
type DeleteSheetConditionFormatReqSheetCfIDs struct {
	SheetID string `json:"sheet_id,omitempty"` // sheet的id
	CfID    string `json:"cf_id,omitempty"`    // 条件格式id
}

// DeleteSheetConditionFormatResp ...
type DeleteSheetConditionFormatResp struct {
	Responses []*DeleteSheetConditionFormatRespResponse `json:"responses,omitempty"` // 响应
}

// DeleteSheetConditionFormatRespResponse ...
type DeleteSheetConditionFormatRespResponse struct {
	SheetID string `json:"sheet_id,omitempty"` // sheet的Id
	CfID    string `json:"cf_id,omitempty"`    // 条件格式id
	ResCode int64  `json:"res_code,omitempty"` // 条件格式删除状态码, 0表示成功, 非0表示失败
	ResMsg  string `json:"res_msg,omitempty"`  // 条件格式删除返回的状态信息, 空表示成功, 非空表示失败原因
}

// deleteSheetConditionFormatResp ...
type deleteSheetConditionFormatResp struct {
	Code int64                           `json:"code,omitempty"`
	Msg  string                          `json:"msg,omitempty"`
	Data *DeleteSheetConditionFormatResp `json:"data,omitempty"`
}
