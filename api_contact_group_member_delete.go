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

// DeleteContactGroupMember 从用户组中移除成员 (目前成员仅支持用户, 未来会支持部门), 如果应用的通讯录权限范围是“全部员工”, 则可将任何成员移出任何用户组。如果应用的通讯录权限范围不是“全部员工”, 则仅可将通讯录权限范围中的成员从通讯录权限范围的用户组中移除, [点击了解通讯录权限范围](https://open.feishu.cn/document/ukTMukTMukTM/uETNz4SM1MjLxUzM/v3/guides/scope_authority)。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group-member/remove
func (r *ContactService) DeleteContactGroupMember(ctx context.Context, request *DeleteContactGroupMemberReq, options ...MethodOptionFunc) (*DeleteContactGroupMemberResp, *Response, error) {
	if r.cli.mock.mockContactDeleteContactGroupMember != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Contact#DeleteContactGroupMember mock enable")
		return r.cli.mock.mockContactDeleteContactGroupMember(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "DeleteContactGroupMember",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/group/:group_id/member/remove",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deleteContactGroupMemberResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactDeleteContactGroupMember mock ContactDeleteContactGroupMember method
func (r *Mock) MockContactDeleteContactGroupMember(f func(ctx context.Context, request *DeleteContactGroupMemberReq, options ...MethodOptionFunc) (*DeleteContactGroupMemberResp, *Response, error)) {
	r.mockContactDeleteContactGroupMember = f
}

// UnMockContactDeleteContactGroupMember un-mock ContactDeleteContactGroupMember method
func (r *Mock) UnMockContactDeleteContactGroupMember() {
	r.mockContactDeleteContactGroupMember = nil
}

// DeleteContactGroupMemberReq ...
type DeleteContactGroupMemberReq struct {
	GroupID      string `path:"group_id" json:"-"`        // 用户组ID, 示例值: "g198123"
	MemberType   string `json:"member_type,omitempty"`    // 用户组成员的类型, 取值为 user, 示例值: "user", 可选值有: `user`: 用户。返回归属于该用户组的用户列表, 默认值: `user`
	MemberID     string `json:"member_id,omitempty"`      // 操作移除的用户组成员ID, 示例值: "xj82871k"
	MemberIDType IDType `json:"member_id_type,omitempty"` // 当member_type =user时候, member_id_type表示user_id_type, 枚举值为open_id, union_id, user_id, 示例值: "open_id", 可选值有: `open_id`: member_type =user时候, 表示用户的open_id, `union_id`: member_type =user时候, 表示用户的union_id, `user_id`: member_type =user时候, 表示用户的user_id, 默认值: `open_id`
}

// DeleteContactGroupMemberResp ...
type DeleteContactGroupMemberResp struct {
}

// deleteContactGroupMemberResp ...
type deleteContactGroupMemberResp struct {
	Code int64                         `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                        `json:"msg,omitempty"`  // 错误描述
	Data *DeleteContactGroupMemberResp `json:"data,omitempty"`
}
