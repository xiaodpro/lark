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
	"io"
)

// UploadApprovalFile
//
// 当审批表单中有图片或附件控件时, 开发者需在创建审批实例前通过审批上传文件接口将文件上传到审批系统, 且附件上传大小限制为50M, 图片上传大小为10M。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDOyUjL1gjM14SN4ITN
func (r *ApprovalService) UploadApprovalFile(ctx context.Context, request *UploadApprovalFileReq, options ...MethodOptionFunc) (*UploadApprovalFileResp, *Response, error) {
	if r.cli.mock.mockApprovalUploadApprovalFile != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Approval#UploadApprovalFile mock enable")
		return r.cli.mock.mockApprovalUploadApprovalFile(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Approval",
		API:                   "UploadApprovalFile",
		Method:                "POST",
		URL:                   r.cli.wwwBaseURL + "/approval/openapi/v2/file/upload",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		IsFile:                true,
	}
	resp := new(uploadApprovalFileResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApprovalUploadApprovalFile mock ApprovalUploadApprovalFile method
func (r *Mock) MockApprovalUploadApprovalFile(f func(ctx context.Context, request *UploadApprovalFileReq, options ...MethodOptionFunc) (*UploadApprovalFileResp, *Response, error)) {
	r.mockApprovalUploadApprovalFile = f
}

// UnMockApprovalUploadApprovalFile un-mock ApprovalUploadApprovalFile method
func (r *Mock) UnMockApprovalUploadApprovalFile() {
	r.mockApprovalUploadApprovalFile = nil
}

// UploadApprovalFileReq ...
type UploadApprovalFileReq struct {
	Name    string    `json:"name,omitempty"`    // 文件名（需包含文件扩展名, 如“文件.doc”
	Type    string    `json:"type,omitempty"`    // 文件类型（image 或 attachment）
	Content io.Reader `json:"content,omitempty"` // 文件
}

// UploadApprovalFileResp ...
type UploadApprovalFileResp struct {
	Code string `json:"code,omitempty"` // 文件标识码（用于创建审批实例）
	URL  string `json:"url,omitempty"`  // 文件 url
}

// uploadApprovalFileResp ...
type uploadApprovalFileResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码, 非0表示失败
	Msg  string                  `json:"msg,omitempty"`  // 返回码的描述
	Data *UploadApprovalFileResp `json:"data,omitempty"` // 返回业务信息
}
