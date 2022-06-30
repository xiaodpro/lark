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

// MatchBaikeEntity 将关键词与词条名、别名精准匹配, 并返回对应的 词条 ID
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/baike-v1/entity/match
func (r *BaikeService) MatchBaikeEntity(ctx context.Context, request *MatchBaikeEntityReq, options ...MethodOptionFunc) (*MatchBaikeEntityResp, *Response, error) {
	if r.cli.mock.mockBaikeMatchBaikeEntity != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Baike#MatchBaikeEntity mock enable")
		return r.cli.mock.mockBaikeMatchBaikeEntity(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Baike",
		API:                   "MatchBaikeEntity",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/baike/v1/entities/match",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(matchBaikeEntityResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBaikeMatchBaikeEntity mock BaikeMatchBaikeEntity method
func (r *Mock) MockBaikeMatchBaikeEntity(f func(ctx context.Context, request *MatchBaikeEntityReq, options ...MethodOptionFunc) (*MatchBaikeEntityResp, *Response, error)) {
	r.mockBaikeMatchBaikeEntity = f
}

// UnMockBaikeMatchBaikeEntity un-mock BaikeMatchBaikeEntity method
func (r *Mock) UnMockBaikeMatchBaikeEntity() {
	r.mockBaikeMatchBaikeEntity = nil
}

// MatchBaikeEntityReq ...
type MatchBaikeEntityReq struct {
	Word string `json:"word,omitempty"` // 搜索关键词, 将与词条名、别名进行精准匹配, 示例值: "企业百科", 长度范围: `1` ～ `100` 字符
}

// MatchBaikeEntityResp ...
type MatchBaikeEntityResp struct {
	Results []*MatchBaikeEntityRespResult `json:"results,omitempty"` // 搜索结果
}

// MatchBaikeEntityRespResult ...
type MatchBaikeEntityRespResult struct {
	EntityID string `json:"entity_id,omitempty"` // 词条 ID
	Type     int64  `json:"type,omitempty"`      // 命中的字段, 可选值有: <md-enum>, <md-enum-item key="0" >词条名</md-enum-item>, <md-enum-item key="1" >全称</md-enum-item>, <md-enum-item key="2" >别名</md-enum-item>, </md-enum>
}

// matchBaikeEntityResp ...
type matchBaikeEntityResp struct {
	Code int64                 `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                `json:"msg,omitempty"`  // 错误描述
	Data *MatchBaikeEntityResp `json:"data,omitempty"`
}
