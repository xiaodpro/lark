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

// CreateMailGroupAlias 创建邮件组别名。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-alias/create
func (r *MailService) CreateMailGroupAlias(ctx context.Context, request *CreateMailGroupAliasReq, options ...MethodOptionFunc) (*CreateMailGroupAliasResp, *Response, error) {
	if r.cli.mock.mockMailCreateMailGroupAlias != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Mail#CreateMailGroupAlias mock enable")
		return r.cli.mock.mockMailCreateMailGroupAlias(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Mail",
		API:                   "CreateMailGroupAlias",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/mail/v1/mailgroups/:mailgroup_id/aliases",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createMailGroupAliasResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMailCreateMailGroupAlias mock MailCreateMailGroupAlias method
func (r *Mock) MockMailCreateMailGroupAlias(f func(ctx context.Context, request *CreateMailGroupAliasReq, options ...MethodOptionFunc) (*CreateMailGroupAliasResp, *Response, error)) {
	r.mockMailCreateMailGroupAlias = f
}

// UnMockMailCreateMailGroupAlias un-mock MailCreateMailGroupAlias method
func (r *Mock) UnMockMailCreateMailGroupAlias() {
	r.mockMailCreateMailGroupAlias = nil
}

// CreateMailGroupAliasReq ...
type CreateMailGroupAliasReq struct {
	MailGroupID string  `path:"mailgroup_id" json:"-"` // 邮件组id或邮件组邮箱地址, 示例值: "xxxxxx 或者 xxx@xx.xxx"
	EmailAlias  *string `json:"email_alias,omitempty"` // 邮箱别名, 示例值: "xxx@xx.xxx"
}

// CreateMailGroupAliasResp ...
type CreateMailGroupAliasResp struct {
	MailGroupAlias *CreateMailGroupAliasRespMailGroupAlias `json:"mailgroup_alias,omitempty"` // 邮件组别名
}

// CreateMailGroupAliasRespMailGroupAlias ...
type CreateMailGroupAliasRespMailGroupAlias struct {
	PrimaryEmail string `json:"primary_email,omitempty"` // 主邮箱地址
	EmailAlias   string `json:"email_alias,omitempty"`   // 邮箱别名
}

// createMailGroupAliasResp ...
type createMailGroupAliasResp struct {
	Code int64                     `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                    `json:"msg,omitempty"`  // 错误描述
	Data *CreateMailGroupAliasResp `json:"data,omitempty"`
}
