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

// SubscribeDriveFile 该接口仅支持文档拥有者和文档管理者订阅文档的通知事件（但目前文档管理者仅能接收到文件编辑事件）。可订阅的文档类型为旧版文档、新版文档、电子表格和多维表格。在调用该接口之前请确保正确[配置事件回调网址和订阅事件类型](https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM#2eb3504a), 目前已支持的事件类型请参考[事件列表](https://open.feishu.cn/document/ukTMukTMukTM/uYDNxYjL2QTM24iN0EjN/event-list)。
//
// 目前只支持订阅事件列表中所有文档事件, 暂不支持只订阅某个或某些事件。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/subscribe
// new doc: https://open.feishu.cn/document/server-docs/docs/drive-v1/event/subscribe
func (r *DriveService) SubscribeDriveFile(ctx context.Context, request *SubscribeDriveFileReq, options ...MethodOptionFunc) (*SubscribeDriveFileResp, *Response, error) {
	if r.cli.mock.mockDriveSubscribeDriveFile != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#SubscribeDriveFile mock enable")
		return r.cli.mock.mockDriveSubscribeDriveFile(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "SubscribeDriveFile",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/files/:file_token/subscribe",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(subscribeDriveFileResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveSubscribeDriveFile mock DriveSubscribeDriveFile method
func (r *Mock) MockDriveSubscribeDriveFile(f func(ctx context.Context, request *SubscribeDriveFileReq, options ...MethodOptionFunc) (*SubscribeDriveFileResp, *Response, error)) {
	r.mockDriveSubscribeDriveFile = f
}

// UnMockDriveSubscribeDriveFile un-mock DriveSubscribeDriveFile method
func (r *Mock) UnMockDriveSubscribeDriveFile() {
	r.mockDriveSubscribeDriveFile = nil
}

// SubscribeDriveFileReq ...
type SubscribeDriveFileReq struct {
	FileToken string   `path:"file_token" json:"-"`  // 文档token, 示例值: "doccnxxxxxxxxxxxxxxxxxxxxxx"
	FileType  FileType `query:"file_type" json:"-"`  // 文档类型, 示例值: doc, 可选值有: doc: 文档, docx: 新版文档, sheet: 表格, bitable: 多维表格, folder: 文件夹
	EventType *string  `query:"event_type" json:"-"` // 事件类型, 订阅为folder类型时必填, 示例值: file.created_in_folder_v1
}

// SubscribeDriveFileResp ...
type SubscribeDriveFileResp struct {
}

// subscribeDriveFileResp ...
type subscribeDriveFileResp struct {
	Code  int64                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                  `json:"msg,omitempty"`  // 错误描述
	Data  *SubscribeDriveFileResp `json:"data,omitempty"`
	Error *ErrorDetail            `json:"error,omitempty"`
}
