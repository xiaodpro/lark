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

// GetMailUserMailboxAliasList 获取用户邮箱所有别名。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/user_mailbox-alias/list
func (r *MailService) GetMailUserMailboxAliasList(ctx context.Context, request *GetMailUserMailboxAliasListReq, options ...MethodOptionFunc) (*GetMailUserMailboxAliasListResp, *Response, error) {
	if r.cli.mock.mockMailGetMailUserMailboxAliasList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Mail#GetMailUserMailboxAliasList mock enable")
		return r.cli.mock.mockMailGetMailUserMailboxAliasList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Mail",
		API:                   "GetMailUserMailboxAliasList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/aliases",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getMailUserMailboxAliasListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMailGetMailUserMailboxAliasList mock MailGetMailUserMailboxAliasList method
func (r *Mock) MockMailGetMailUserMailboxAliasList(f func(ctx context.Context, request *GetMailUserMailboxAliasListReq, options ...MethodOptionFunc) (*GetMailUserMailboxAliasListResp, *Response, error)) {
	r.mockMailGetMailUserMailboxAliasList = f
}

// UnMockMailGetMailUserMailboxAliasList un-mock MailGetMailUserMailboxAliasList method
func (r *Mock) UnMockMailGetMailUserMailboxAliasList() {
	r.mockMailGetMailUserMailboxAliasList = nil
}

// GetMailUserMailboxAliasListReq ...
type GetMailUserMailboxAliasListReq struct {
	UserMailboxID string  `path:"user_mailbox_id" json:"-"` // 用户邮箱地址, 示例值: "user@xxx.xx"
	PageToken     *string `query:"page_token" json:"-"`     // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "xxx"
	PageSize      *int64  `query:"page_size" json:"-"`      // 分页大小, 示例值: 10, 最大值: `20`
}

// GetMailUserMailboxAliasListResp ...
type GetMailUserMailboxAliasListResp struct {
	Items []*GetMailUserMailboxAliasListRespItem `json:"items,omitempty"` // 用户邮箱别名
}

// GetMailUserMailboxAliasListRespItem ...
type GetMailUserMailboxAliasListRespItem struct {
	PrimaryEmail string `json:"primary_email,omitempty"` // 主邮箱地址
	EmailAlias   string `json:"email_alias,omitempty"`   // 邮箱别名
}

// getMailUserMailboxAliasListResp ...
type getMailUserMailboxAliasListResp struct {
	Code int64                            `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                           `json:"msg,omitempty"`  // 错误描述
	Data *GetMailUserMailboxAliasListResp `json:"data,omitempty"`
}
