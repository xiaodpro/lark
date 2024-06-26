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

// UpdateDriveDocContent 此接口只支持编辑旧版文档内容。要编辑新版文档（文档类型: `docx`）的内容, 调用以下接口:
//
// - [创建块](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-v1/document-block-children/create)
// - [更新块](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-v1/document-block/patch)
// - [批量更新块](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-v1/document-block/batch_update)
// - [删除块](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-v1/document-block-children/batch_delete)
// 该接口用于批量编辑更新文档内容, 包括更新标题、范围删除、插入内容。
// 前提条件:
// 在使用此接口前, 请仔细阅读[文档概述](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/docs-doc-overview)和[准备接入文档 API](https://open.feishu.cn/document/ukTMukTMukTM/ugzNzUjL4czM14CO3MTN/guide/getting-start)了解文档调用的规则和约束, 确保你的文档数据不会丢失或出错。
// 文档数据结构定义可参考: [文档数据结构概述](https://open.feishu.cn/document/ukTMukTMukTM/uAzM5YjLwMTO24CMzkjN)。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uYDM2YjL2AjN24iNwYjN
// new doc: https://open.feishu.cn/document/server-docs/docs/docs/docs/content/batch-update-document
//
// Deprecated
func (r *DriveService) UpdateDriveDocContent(ctx context.Context, request *UpdateDriveDocContentReq, options ...MethodOptionFunc) (*UpdateDriveDocContentResp, *Response, error) {
	if r.cli.mock.mockDriveUpdateDriveDocContent != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#UpdateDriveDocContent mock enable")
		return r.cli.mock.mockDriveUpdateDriveDocContent(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "UpdateDriveDocContent",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/doc/v2/:docToken/batch_update",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateDriveDocContentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveUpdateDriveDocContent mock DriveUpdateDriveDocContent method
func (r *Mock) MockDriveUpdateDriveDocContent(f func(ctx context.Context, request *UpdateDriveDocContentReq, options ...MethodOptionFunc) (*UpdateDriveDocContentResp, *Response, error)) {
	r.mockDriveUpdateDriveDocContent = f
}

// UnMockDriveUpdateDriveDocContent un-mock DriveUpdateDriveDocContent method
func (r *Mock) UnMockDriveUpdateDriveDocContent() {
	r.mockDriveUpdateDriveDocContent = nil
}

// UpdateDriveDocContentReq ...
type UpdateDriveDocContentReq struct {
	DocToken string   `path:"docToken" json:"-"`  // 文件的 token, 获取方式见[如何获取云文档资源相关 token](https://open.feishu.cn/document/ukTMukTMukTM/uczNzUjL3czM14yN3MTN#08bb5df6)
	Revision int64    `json:"Revision,omitempty"` // 文档的指定版本, 文档新创建后版本号是0, [获取方式](https://open.feishu.cn/document/ukTMukTMukTM/uUDM2YjL1AjN24SNwYjN), 要求>=0, post body json 字段
	Requests []string `json:"Requests,omitempty"` // post body json, OperationRequest 类型序列化 string 数组
}

// UpdateDriveDocContentResp ...
type UpdateDriveDocContentResp struct {
}

// updateDriveDocContentResp ...
type updateDriveDocContentResp struct {
	Code  int64                      `json:"code,omitempty"`
	Msg   string                     `json:"msg,omitempty"`
	Data  *UpdateDriveDocContentResp `json:"data,omitempty"`
	Error *ErrorDetail               `json:"error,omitempty"`
}
