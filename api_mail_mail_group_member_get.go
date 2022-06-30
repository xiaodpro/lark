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

// GetMailGroupMember 获取邮件组单个成员信息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-member/get
func (r *MailService) GetMailGroupMember(ctx context.Context, request *GetMailGroupMemberReq, options ...MethodOptionFunc) (*GetMailGroupMemberResp, *Response, error) {
	if r.cli.mock.mockMailGetMailGroupMember != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Mail#GetMailGroupMember mock enable")
		return r.cli.mock.mockMailGetMailGroupMember(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Mail",
		API:                   "GetMailGroupMember",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/mail/v1/mailgroups/:mailgroup_id/members/:member_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getMailGroupMemberResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMailGetMailGroupMember mock MailGetMailGroupMember method
func (r *Mock) MockMailGetMailGroupMember(f func(ctx context.Context, request *GetMailGroupMemberReq, options ...MethodOptionFunc) (*GetMailGroupMemberResp, *Response, error)) {
	r.mockMailGetMailGroupMember = f
}

// UnMockMailGetMailGroupMember un-mock MailGetMailGroupMember method
func (r *Mock) UnMockMailGetMailGroupMember() {
	r.mockMailGetMailGroupMember = nil
}

// GetMailGroupMemberReq ...
type GetMailGroupMemberReq struct {
	MailGroupID      string            `path:"mailgroup_id" json:"-"`        // 邮件组ID或者邮件组地址, 示例值: "xxxxxxxxxxxxxxx 或 test_mail_group@xxx.xx"
	MemberID         string            `path:"member_id" json:"-"`           // 邮件组内成员唯一标识, 示例值: "xxxxxxxxxxxxxxx"
	UserIDType       *IDType           `query:"user_id_type" json:"-"`       // 用户 ID 类型, 示例值: "open_id", 可选值有: `open_id`: 用户的 open id, `union_id`: 用户的 union id, `user_id`: 用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	DepartmentIDType *DepartmentIDType `query:"department_id_type" json:"-"` // 此次调用中使用的部门ID的类型, 示例值: "open_department_id", 可选值有: `department_id`: 以自定义department_id来标识部门, `open_department_id`: 以open_department_id来标识部门
}

// GetMailGroupMemberResp ...
type GetMailGroupMemberResp struct {
	MemberID     string       `json:"member_id,omitempty"`     // 邮件组内成员唯一标识
	Email        string       `json:"email,omitempty"`         // 成员邮箱地址（当成员类型是EXTERNAL_USER/MAIL_GROUP/OTHER_MEMBER时有值）
	UserID       string       `json:"user_id,omitempty"`       // 租户内用户的唯一标识（当成员类型是USER时有值）
	DepartmentID string       `json:"department_id,omitempty"` // 租户内部门的唯一标识（当成员类型是DEPARTMENT时有值）
	Type         MailUserType `json:"type,omitempty"`          // 成员类型, 可选值有: `USER`: 内部用户, `DEPARTMENT`: 部门, `COMPANY`: 全员, `EXTERNAL_USER`: 外部用户, `MAIL_GROUP`: 邮件组, `OTHER_MEMBER`: 内部成员
}

// getMailGroupMemberResp ...
type getMailGroupMemberResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *GetMailGroupMemberResp `json:"data,omitempty"`
}
