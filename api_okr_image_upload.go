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

// UploadOKRImage 上传进展记录图片。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/okr-v1/image/upload
func (r *OKRService) UploadOKRImage(ctx context.Context, request *UploadOKRImageReq, options ...MethodOptionFunc) (*UploadOKRImageResp, *Response, error) {
	if r.cli.mock.mockOKRUploadOKRImage != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] OKR#UploadOKRImage mock enable")
		return r.cli.mock.mockOKRUploadOKRImage(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "OKR",
		API:                   "UploadOKRImage",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/okr/v1/images/upload",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(uploadOKRImageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockOKRUploadOKRImage mock OKRUploadOKRImage method
func (r *Mock) MockOKRUploadOKRImage(f func(ctx context.Context, request *UploadOKRImageReq, options ...MethodOptionFunc) (*UploadOKRImageResp, *Response, error)) {
	r.mockOKRUploadOKRImage = f
}

// UnMockOKRUploadOKRImage un-mock OKRUploadOKRImage method
func (r *Mock) UnMockOKRUploadOKRImage() {
	r.mockOKRUploadOKRImage = nil
}

// UploadOKRImageReq ...
type UploadOKRImageReq struct {
	Data       *UploadOKRImageReqData `json:"data,omitempty"`        // 图片, 示例值: file binary
	TargetID   string                 `json:"target_id,omitempty"`   // 图片的目标ID, 示例值: "6974586812998174252"
	TargetType int64                  `json:"target_type,omitempty"` // 图片使用的目标类型, 示例值: 1, 可选值有: 2: okr的O, 3: okr的KR
}

// UploadOKRImageReqData ...
type UploadOKRImageReqData struct {
}

// UploadOKRImageResp ...
type UploadOKRImageResp struct {
	FileToken string `json:"file_token,omitempty"` // 图片token
	URL       string `json:"url,omitempty"`        // 图片下载链接
}

// uploadOKRImageResp ...
type uploadOKRImageResp struct {
	Code int64               `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 错误描述
	Data *UploadOKRImageResp `json:"data,omitempty"`
}
