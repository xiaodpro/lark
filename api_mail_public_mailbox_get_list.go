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

// GetPublicMailboxList 分页批量获取公共邮箱列表。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox/list
func (r *MailService) GetPublicMailboxList(ctx context.Context, request *GetPublicMailboxListReq, options ...MethodOptionFunc) (*GetPublicMailboxListResp, *Response, error) {
	if r.cli.mock.mockMailGetPublicMailboxList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Mail#GetPublicMailboxList mock enable")
		return r.cli.mock.mockMailGetPublicMailboxList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Mail",
		API:                   "GetPublicMailboxList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/mail/v1/public_mailboxes",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getPublicMailboxListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMailGetPublicMailboxList mock MailGetPublicMailboxList method
func (r *Mock) MockMailGetPublicMailboxList(f func(ctx context.Context, request *GetPublicMailboxListReq, options ...MethodOptionFunc) (*GetPublicMailboxListResp, *Response, error)) {
	r.mockMailGetPublicMailboxList = f
}

// UnMockMailGetPublicMailboxList un-mock MailGetPublicMailboxList method
func (r *Mock) UnMockMailGetPublicMailboxList() {
	r.mockMailGetPublicMailboxList = nil
}

// GetPublicMailboxListReq ...
type GetPublicMailboxListReq struct {
	PageToken *string `query:"page_token" json:"-"` // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "xxx"
	PageSize  *int64  `query:"page_size" json:"-"`  // 分页大小, 示例值: 10, 最大值: `200`
}

// GetPublicMailboxListResp ...
type GetPublicMailboxListResp struct {
	HasMore   bool                            `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                          `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	Items     []*GetPublicMailboxListRespItem `json:"items,omitempty"`      // 公共邮箱列表
}

// GetPublicMailboxListRespItem ...
type GetPublicMailboxListRespItem struct {
	PublicMailboxID string `json:"public_mailbox_id,omitempty"` // 公共邮箱唯一标识
	Email           string `json:"email,omitempty"`             // 公共邮箱地址
	Name            string `json:"name,omitempty"`              // 公共邮箱名称
}

// getPublicMailboxListResp ...
type getPublicMailboxListResp struct {
	Code int64                     `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                    `json:"msg,omitempty"`  // 错误描述
	Data *GetPublicMailboxListResp `json:"data,omitempty"`
}
