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

// GetContactGroupMember 通过该接口可查询某个用户组的成员列表（支持查询成员中的用户和部门）, 本接口支持普通用户组和动态用户组。如果应用的通讯录权限范围是“全部员工”, 则可查询企业内任何用户组的成员列表。如果应用的通讯录权限范围不是“全部员工”, 则仅可查询通讯录权限范围中的用户组的成员列表, [点击了解通讯录权限范围](https://open.feishu.cn/document/ukTMukTMukTM/uETNz4SM1MjLxUzM/v3/guides/scope_authority)。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group-member/simplelist
// new doc: https://open.feishu.cn/document/server-docs/contact-v3/group-member/simplelist
func (r *ContactService) GetContactGroupMember(ctx context.Context, request *GetContactGroupMemberReq, options ...MethodOptionFunc) (*GetContactGroupMemberResp, *Response, error) {
	if r.cli.mock.mockContactGetContactGroupMember != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Contact#GetContactGroupMember mock enable")
		return r.cli.mock.mockContactGetContactGroupMember(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "GetContactGroupMember",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/group/:group_id/member/simplelist",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getContactGroupMemberResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactGetContactGroupMember mock ContactGetContactGroupMember method
func (r *Mock) MockContactGetContactGroupMember(f func(ctx context.Context, request *GetContactGroupMemberReq, options ...MethodOptionFunc) (*GetContactGroupMemberResp, *Response, error)) {
	r.mockContactGetContactGroupMember = f
}

// UnMockContactGetContactGroupMember un-mock ContactGetContactGroupMember method
func (r *Mock) UnMockContactGetContactGroupMember() {
	r.mockContactGetContactGroupMember = nil
}

// GetContactGroupMemberReq ...
type GetContactGroupMemberReq struct {
	GroupID      string  `path:"group_id" json:"-"`        // 用户组ID, 示例值: "g128187"
	PageSize     *int64  `query:"page_size" json:"-"`      // 分页大小, 示例值: 50, 默认值: `50`, 最大值: `100`
	PageToken    *string `query:"page_token" json:"-"`     // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: AQD9/Rn9eij9Pm39ED40/dk53s4Ebp882DYfFaPFbz00L4CMZJrqGdzNyc8BcZtDbwVUvRmQTvyMYicnGWrde9X56TgdBuS+JKiSIkdexPw=
	MemberIDType *IDType `query:"member_id_type" json:"-"` // 欲获取成员ID类型, 当member_type=user时候, member_id_type表示user_id_type, 枚举值open_id, union_id和user_id, 当member_type=department时候, member_id_type表示department_id_type, 枚举值open_id和department_id, 示例值: open_id, 可选值有: open_id: member_type =user时候, 表示用户的open_id, union_id: member_type =user时候, 表示用户的union_id, user_id: member_type =user时候, 表示用户的user_id, department_id: member_type=department时候, 表示部门的department_id, 默认值: `open_id`
	MemberType   *string `query:"member_type" json:"-"`    // 欲获取的用户组成员类型, 示例值: user, 可选值有: user: 用户。返回归属于该用户组的用户列表, department: 部门。返回归属于该用户组的部门列表, 默认值: `user`
}

// GetContactGroupMemberResp ...
type GetContactGroupMemberResp struct {
	Memberlist []*GetContactGroupMemberRespMember `json:"memberlist,omitempty"` // 成员列表
	PageToken  string                             `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore    bool                               `json:"has_more,omitempty"`   // 是否还有更多项
}

// GetContactGroupMemberRespMember ...
type GetContactGroupMemberRespMember struct {
	MemberID     string `json:"member_id,omitempty"`      // 成员ID
	MemberType   string `json:"member_type,omitempty"`    // 用户组成员的类型, 取值为 user或department。
	MemberIDType IDType `json:"member_id_type,omitempty"` // 当member_type为user时, member_id_type表示user_id_type, 可选值为open_id, union_id, user_id。仅在请求参数中有效, 响应体中不会返回此参数。
}

// getContactGroupMemberResp ...
type getContactGroupMemberResp struct {
	Code  int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                     `json:"msg,omitempty"`  // 错误描述
	Data  *GetContactGroupMemberResp `json:"data,omitempty"`
	Error *ErrorDetail               `json:"error,omitempty"`
}
