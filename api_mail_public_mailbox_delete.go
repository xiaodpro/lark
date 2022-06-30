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

// DeletePublicMailbox 该接口会永久删除公共邮箱地址。可用于释放邮箱回收站的公共邮箱地址, 一旦删除, 该邮箱地址将无法恢复。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox/delete
func (r *MailService) DeletePublicMailbox(ctx context.Context, request *DeletePublicMailboxReq, options ...MethodOptionFunc) (*DeletePublicMailboxResp, *Response, error) {
	if r.cli.mock.mockMailDeletePublicMailbox != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Mail#DeletePublicMailbox mock enable")
		return r.cli.mock.mockMailDeletePublicMailbox(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Mail",
		API:                   "DeletePublicMailbox",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/mail/v1/public_mailboxes/:public_mailbox_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deletePublicMailboxResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMailDeletePublicMailbox mock MailDeletePublicMailbox method
func (r *Mock) MockMailDeletePublicMailbox(f func(ctx context.Context, request *DeletePublicMailboxReq, options ...MethodOptionFunc) (*DeletePublicMailboxResp, *Response, error)) {
	r.mockMailDeletePublicMailbox = f
}

// UnMockMailDeletePublicMailbox un-mock MailDeletePublicMailbox method
func (r *Mock) UnMockMailDeletePublicMailbox() {
	r.mockMailDeletePublicMailbox = nil
}

// DeletePublicMailboxReq ...
type DeletePublicMailboxReq struct {
	PublicMailboxID string `path:"public_mailbox_id" json:"-"` // 要释放的公共邮箱地址, 示例值: "xxxxxx@abc.com"
}

// DeletePublicMailboxResp ...
type DeletePublicMailboxResp struct {
}

// deletePublicMailboxResp ...
type deletePublicMailboxResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *DeletePublicMailboxResp `json:"data,omitempty"`
}
