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

// CreateIMTagRelation 绑定标签到业务实体。目前支持给会话打标签。
//
// doc: https://open.larkoffice.com/document/uAjLw4CM/ukTMukTMukTM/group/im-v2/biz_entity_tag_relation/create
func (r *MessageService) CreateIMTagRelation(ctx context.Context, request *CreateIMTagRelationReq, options ...MethodOptionFunc) (*CreateIMTagRelationResp, *Response, error) {
	if r.cli.mock.mockMessageCreateIMTagRelation != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Message#CreateIMTagRelation mock enable")
		return r.cli.mock.mockMessageCreateIMTagRelation(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Message",
		API:                   "CreateIMTagRelation",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v2/biz_entity_tag_relation",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createIMTagRelationResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMessageCreateIMTagRelation mock MessageCreateIMTagRelation method
func (r *Mock) MockMessageCreateIMTagRelation(f func(ctx context.Context, request *CreateIMTagRelationReq, options ...MethodOptionFunc) (*CreateIMTagRelationResp, *Response, error)) {
	r.mockMessageCreateIMTagRelation = f
}

// UnMockMessageCreateIMTagRelation un-mock MessageCreateIMTagRelation method
func (r *Mock) UnMockMessageCreateIMTagRelation() {
	r.mockMessageCreateIMTagRelation = nil
}

// CreateIMTagRelationReq ...
type CreateIMTagRelationReq struct {
	TagBizType  string   `json:"tag_biz_type,omitempty"`  // 业务类型, 示例值: "chat", 可选值有: chat: chat 会话类型
	BizEntityID string   `json:"biz_entity_id,omitempty"` // 业务实体 ID, 示例值: "oc_xxxxx"
	TagIDs      []string `json:"tag_ids,omitempty"`       // 标签 ID, 示例值: ["71616xxxx"], 长度范围: `0` ～ `40`
}

// CreateIMTagRelationResp ...
type CreateIMTagRelationResp struct {
}

// createIMTagRelationResp ...
type createIMTagRelationResp struct {
	Code  int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                   `json:"msg,omitempty"`  // 错误描述
	Data  *CreateIMTagRelationResp `json:"data,omitempty"`
	Error *ErrorDetail             `json:"error,omitempty"`
}
