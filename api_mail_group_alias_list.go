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

// GetMailGroupAliasList 获取邮件组所有别名。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-alias/list
func (r *MailService) GetMailGroupAliasList(ctx context.Context, request *GetMailGroupAliasListReq, options ...MethodOptionFunc) (*GetMailGroupAliasListResp, *Response, error) {
	if r.cli.mock.mockMailGetMailGroupAliasList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Mail#GetMailGroupAliasList mock enable")
		return r.cli.mock.mockMailGetMailGroupAliasList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Mail",
		API:                   "GetMailGroupAliasList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/mail/v1/mailgroups/:mailgroup_id/aliases",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getMailGroupAliasListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMailGetMailGroupAliasList mock MailGetMailGroupAliasList method
func (r *Mock) MockMailGetMailGroupAliasList(f func(ctx context.Context, request *GetMailGroupAliasListReq, options ...MethodOptionFunc) (*GetMailGroupAliasListResp, *Response, error)) {
	r.mockMailGetMailGroupAliasList = f
}

// UnMockMailGetMailGroupAliasList un-mock MailGetMailGroupAliasList method
func (r *Mock) UnMockMailGetMailGroupAliasList() {
	r.mockMailGetMailGroupAliasList = nil
}

// GetMailGroupAliasListReq ...
type GetMailGroupAliasListReq struct {
	MailGroupID string `path:"mailgroup_id" json:"-"` // 邮件组id或邮件组邮箱地址, 示例值: "xxxxxxxxxxxxxxx 或 test_mail_group@xxx.xx"
}

// GetMailGroupAliasListResp ...
type GetMailGroupAliasListResp struct {
	Items []*GetMailGroupAliasListRespItem `json:"items,omitempty"` // 邮件组别名
}

// GetMailGroupAliasListRespItem ...
type GetMailGroupAliasListRespItem struct {
	PrimaryEmail string `json:"primary_email,omitempty"` // 主邮箱地址
	EmailAlias   string `json:"email_alias,omitempty"`   // 邮箱别名
}

// getMailGroupAliasListResp ...
type getMailGroupAliasListResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *GetMailGroupAliasListResp `json:"data,omitempty"`
}
