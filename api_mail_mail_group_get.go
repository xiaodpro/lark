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

// GetMailGroup 获取特定邮件组信息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup/get
func (r *MailService) GetMailGroup(ctx context.Context, request *GetMailGroupReq, options ...MethodOptionFunc) (*GetMailGroupResp, *Response, error) {
	if r.cli.mock.mockMailGetMailGroup != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Mail#GetMailGroup mock enable")
		return r.cli.mock.mockMailGetMailGroup(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Mail",
		API:                   "GetMailGroup",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/mail/v1/mailgroups/:mailgroup_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getMailGroupResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMailGetMailGroup mock MailGetMailGroup method
func (r *Mock) MockMailGetMailGroup(f func(ctx context.Context, request *GetMailGroupReq, options ...MethodOptionFunc) (*GetMailGroupResp, *Response, error)) {
	r.mockMailGetMailGroup = f
}

// UnMockMailGetMailGroup un-mock MailGetMailGroup method
func (r *Mock) UnMockMailGetMailGroup() {
	r.mockMailGetMailGroup = nil
}

// GetMailGroupReq ...
type GetMailGroupReq struct {
	MailGroupID string `path:"mailgroup_id" json:"-"` // 邮件组ID或者邮件组地址, 示例值: "xxxxxxxxxxxxxxx 或 test_mail_group@xxx.xx"
}

// GetMailGroupResp ...
type GetMailGroupResp struct {
	MailGroupID             string `json:"mailgroup_id,omitempty"`               // 邮件组ID
	Email                   string `json:"email,omitempty"`                      // 邮件组地址
	Name                    string `json:"name,omitempty"`                       // 邮件组名称
	Description             string `json:"description,omitempty"`                // 邮件组描述
	DirectMembersCount      string `json:"direct_members_count,omitempty"`       // 邮件组成员数量
	IncludeExternalMember   bool   `json:"include_external_member,omitempty"`    // 是否包含外部成员
	IncludeAllCompanyMember bool   `json:"include_all_company_member,omitempty"` // 是否是全员邮件组
	WhoCanSendMail          string `json:"who_can_send_mail,omitempty"`          // 谁可发送邮件到此邮件组, 可选值有: `ANYONE`: 任何人, `ALL_INTERNAL_USERS`: 仅组织内部成员, `ALL_GROUP_MEMBERS`: 仅邮件组成员, `CUSTOM_MEMBERS`: 自定义成员
}

// getMailGroupResp ...
type getMailGroupResp struct {
	Code int64             `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string            `json:"msg,omitempty"`  // 错误描述
	Data *GetMailGroupResp `json:"data,omitempty"`
}
