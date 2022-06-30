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

// GetDriveDocContent
//
// 在使用此接口前, 请仔细阅读[文档概述](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/docs-doc-overview)和[准备接入文档 API](https://open.feishu.cn/document/ukTMukTMukTM/ugzNzUjL4czM14CO3MTN/guide/getting-start)了解文档调用的规则和约束, 确保你的文档数据不会丢失或出错。
// 文档数据结构定义可参考: [文档数据结构概述](https://open.feishu.cn/document/ukTMukTMukTM/uAzM5YjLwMTO24CMzkjN)
// 此接口只支持获取旧版文档富文本内容, 如果需要获取新版文档的富文本内容, 请调用新版文档相关接口:
// - [获取文档所有块](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-v1/document-block/list)
// - [获取指定块](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-v1/document-block/get)
// - [获取指定块下所有子块](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-v1/document-block-children/get)
// 该接口用于获取结构化的文档内容。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDM2YjL1AjN24SNwYjN
func (r *DriveService) GetDriveDocContent(ctx context.Context, request *GetDriveDocContentReq, options ...MethodOptionFunc) (*GetDriveDocContentResp, *Response, error) {
	if r.cli.mock.mockDriveGetDriveDocContent != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#GetDriveDocContent mock enable")
		return r.cli.mock.mockDriveGetDriveDocContent(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "GetDriveDocContent",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/doc/v2/:docToken/content",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getDriveDocContentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveGetDriveDocContent mock DriveGetDriveDocContent method
func (r *Mock) MockDriveGetDriveDocContent(f func(ctx context.Context, request *GetDriveDocContentReq, options ...MethodOptionFunc) (*GetDriveDocContentResp, *Response, error)) {
	r.mockDriveGetDriveDocContent = f
}

// UnMockDriveGetDriveDocContent un-mock DriveGetDriveDocContent method
func (r *Mock) UnMockDriveGetDriveDocContent() {
	r.mockDriveGetDriveDocContent = nil
}

// GetDriveDocContentReq ...
type GetDriveDocContentReq struct {
	DocToken string `path:"docToken" json:"-"` // 获取方式详见[如何获取云文档资源相关 token](https://open.feishu.cn/document/ukTMukTMukTM/uczNzUjL3czM14yN3MTN#08bb5df6)
}

// GetDriveDocContentResp ...
type GetDriveDocContentResp struct {
	Content  string `json:"content,omitempty"`  // 详情参考[文档数据结构](https://open.feishu.cn/document/ukTMukTMukTM/ukDM2YjL5AjN24SOwYjN)
	Revision int64  `json:"revision,omitempty"` // 文档当前版本号
}

// getDriveDocContentResp ...
type getDriveDocContentResp struct {
	Code int64                   `json:"code,omitempty"`
	Msg  string                  `json:"msg,omitempty"`
	Data *GetDriveDocContentResp `json:"data,omitempty"`
}
