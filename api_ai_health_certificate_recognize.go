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

// RecognizeAIHealthCertificate 健康证识别接口, 支持JPG/JPEG/PNG/BMP四种文件类型的一次性的识别。
//
// 单租户限流: 10QPS, 同租户下的应用没有限流, 共享本租户的 10QPS 限流
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/document_ai-v1/health_certificate/recognize
func (r *AIService) RecognizeAIHealthCertificate(ctx context.Context, request *RecognizeAIHealthCertificateReq, options ...MethodOptionFunc) (*RecognizeAIHealthCertificateResp, *Response, error) {
	if r.cli.mock.mockAIRecognizeAIHealthCertificate != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] AI#RecognizeAIHealthCertificate mock enable")
		return r.cli.mock.mockAIRecognizeAIHealthCertificate(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "AI",
		API:                   "RecognizeAIHealthCertificate",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/document_ai/v1/health_certificate/recognize",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
		IsFile:                true,
	}
	resp := new(recognizeAIHealthCertificateResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAIRecognizeAIHealthCertificate mock AIRecognizeAIHealthCertificate method
func (r *Mock) MockAIRecognizeAIHealthCertificate(f func(ctx context.Context, request *RecognizeAIHealthCertificateReq, options ...MethodOptionFunc) (*RecognizeAIHealthCertificateResp, *Response, error)) {
	r.mockAIRecognizeAIHealthCertificate = f
}

// UnMockAIRecognizeAIHealthCertificate un-mock AIRecognizeAIHealthCertificate method
func (r *Mock) UnMockAIRecognizeAIHealthCertificate() {
	r.mockAIRecognizeAIHealthCertificate = nil
}

// RecognizeAIHealthCertificateReq ...
type RecognizeAIHealthCertificateReq struct {
	File io.Reader `json:"file,omitempty"` // 识别的健康证源文件, 示例值: file binary
}

// RecognizeAIHealthCertificateResp ...
type RecognizeAIHealthCertificateResp struct {
	HealthCertificate *RecognizeAIHealthCertificateRespHealthCertificate `json:"health_certificate,omitempty"` // 健康证信息
}

// RecognizeAIHealthCertificateRespHealthCertificate ...
type RecognizeAIHealthCertificateRespHealthCertificate struct {
	Entities []*RecognizeAIHealthCertificateRespHealthCertificateEntitie `json:"entities,omitempty"` // 识别出的实体类型
}

// RecognizeAIHealthCertificateRespHealthCertificateEntitie ...
type RecognizeAIHealthCertificateRespHealthCertificateEntitie struct {
	Type  string `json:"type,omitempty"`  // 识别的字段种类, 可选值有: name: 姓名, issued_by: 发证机关, date_of_handling: 办证日期, date_of_issue: 发证日期, date_of_medical_examination: 体检日期, valid_date: 有效日期, other_date: 其他日期
	Value string `json:"value,omitempty"` // 识别出字段的文本信息
}

// recognizeAIHealthCertificateResp ...
type recognizeAIHealthCertificateResp struct {
	Code  int64                             `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                            `json:"msg,omitempty"`  // 错误描述
	Data  *RecognizeAIHealthCertificateResp `json:"data,omitempty"`
	Error *ErrorDetail                      `json:"error,omitempty"`
}
