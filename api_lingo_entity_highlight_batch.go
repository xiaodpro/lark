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

// BatchHighlightLingoEntity 通过这个接口, 可以传入一段文本, 获取这段文本中所有词条的 ID
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/lingo-v1/entity/batch_highlight
func (r *LingoService) BatchHighlightLingoEntity(ctx context.Context, request *BatchHighlightLingoEntityReq, options ...MethodOptionFunc) (*BatchHighlightLingoEntityResp, *Response, error) {
	if r.cli.mock.mockLingoBatchHighlightLingoEntity != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Lingo#BatchHighlightLingoEntity mock enable")
		return r.cli.mock.mockLingoBatchHighlightLingoEntity(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Lingo",
		API:                   "BatchHighlightLingoEntity",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/lingo/v1/entities/batch_highlight",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(batchHighlightLingoEntityResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockLingoBatchHighlightLingoEntity mock LingoBatchHighlightLingoEntity method
func (r *Mock) MockLingoBatchHighlightLingoEntity(f func(ctx context.Context, request *BatchHighlightLingoEntityReq, options ...MethodOptionFunc) (*BatchHighlightLingoEntityResp, *Response, error)) {
	r.mockLingoBatchHighlightLingoEntity = f
}

// UnMockLingoBatchHighlightLingoEntity un-mock LingoBatchHighlightLingoEntity method
func (r *Mock) UnMockLingoBatchHighlightLingoEntity() {
	r.mockLingoBatchHighlightLingoEntity = nil
}

// BatchHighlightLingoEntityReq ...
type BatchHighlightLingoEntityReq struct {
	Texts []string `json:"texts,omitempty"` // 一批需要被识别词条的文本（一批不要超过20段文本, 每段文本不要超过1000字）, 示例值: ["飞书词典是飞书提供的一款知识管理工具"], 长度范围: `1` ～ `20`
}

// BatchHighlightLingoEntityResp ...
type BatchHighlightLingoEntityResp struct {
	Phrases [][]*BatchHighlightLingoEntityRespPhrases `json:"phrases,omitempty"` // 和输入texts对应长度的段落实体词信息
}

// batchHighlightLingoEntityResp ...
type batchHighlightLingoEntityResp struct {
	Code int64                          `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                         `json:"msg,omitempty"`  // 错误描述
	Data *BatchHighlightLingoEntityResp `json:"data,omitempty"`
}